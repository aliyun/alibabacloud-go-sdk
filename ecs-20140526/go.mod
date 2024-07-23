module github.com/alibabacloud-go/ecs-20140526/v4

go 1.16

require (
	github.com/alibabacloud-go/darabonba-openapi/v2 v2.0.7
	github.com/alibabacloud-go/endpoint-util v1.1.0
	github.com/alibabacloud-go/openapi-util v0.1.0
	github.com/alibabacloud-go/tea v1.2.1
	github.com/alibabacloud-go/tea-utils/v2 v2.0.5
)

retract v4.24.17 // This version has critical bugs.
