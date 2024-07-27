package route

import (
    "github.com/matryer/is"
    "src/app/config"
    "src/app/entity"
    "testing"
)

func Test_append_route_by_path_returns_original_pattern(t *testing.T) {
    // Given
    routes := []entity.Route{{Pattern: "GET /images/"}}

    // When
    routes = appendApiByPath(routes)

    // Then
    i := is.New(t)
    i.Equal(routes[0].Pattern, "GET /images/")
}

func Test_append_route_by_path_returns_new_pattern(t *testing.T) {
    // Given
    routes := []entity.Route{{Pattern: "GET /images/"}}
    config.AppInfo.Service = "service"
    config.AppInfo.ApiByPathPrefix = "conf_api"

    // When
    routes = appendApiByPath(routes)

    // Then
    i := is.New(t)
    i.Equal(routes[1].Pattern, "GET /conf_api/service/images/")
}

func Test_append_route_by_path_without_method(t *testing.T) {
    // Given
    routes := []entity.Route{{Pattern: "/images/"}}
    config.AppInfo.Service = "service"
    config.AppInfo.ApiByPathPrefix = "conf_api"

    // When
    routes = appendApiByPath(routes)

    // Then
    i := is.New(t)
    i.Equal(routes[1].Pattern, "/conf_api/service/images/")
}

func Test_append_route_by_path_without_method_with_space_on_the_end(t *testing.T) {
    // Given
    routes := []entity.Route{{Pattern: "/images/1 2"}}
    config.AppInfo.Service = "service"
    config.AppInfo.ApiByPathPrefix = "conf_api"

    // When
    routes = appendApiByPath(routes)

    // Then
    i := is.New(t)
    i.Equal(routes[1].Pattern, "/conf_api/service/images/1 2")
}
