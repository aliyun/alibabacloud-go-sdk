// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContainerConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateContainerConfigurationRequest
	GetAppId() *string
	SetContextPath(v string) *UpdateContainerConfigurationRequest
	GetContextPath() *string
	SetGroupId(v string) *UpdateContainerConfigurationRequest
	GetGroupId() *string
	SetHttpPort(v int32) *UpdateContainerConfigurationRequest
	GetHttpPort() *int32
	SetMaxThreads(v int32) *UpdateContainerConfigurationRequest
	GetMaxThreads() *int32
	SetURIEncoding(v string) *UpdateContainerConfigurationRequest
	GetURIEncoding() *string
	SetUseBodyEncoding(v bool) *UpdateContainerConfigurationRequest
	GetUseBodyEncoding() *bool
}

type UpdateContainerConfigurationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// c627c157-560d-43ff-************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The context path of the Tomcat container. The context path can be an empty string, a null WAR package name, a root directory, or other custom non-empty strings. It can contain letters, digits, hyphens (-), and underscores (_). Take note of the following items:
	//
	// 	- If this parameter is not specified when you configure the application instance group, the configuration of the application is applied.
	//
	// 	- If this parameter is not specified when you configure the Tomcat container for an application, the root directory `/` is used.
	//
	// example:
	//
	// /
	ContextPath *string `json:"ContextPath,omitempty" xml:"ContextPath,omitempty"`
	// The ID of the application instance group.
	//
	// 	- If an ID is specified, this operation configures the Tomcat container for the specified application instance group.
	//
	// 	- If you set this parameter to "", this operation configures the Tomcat container for the application.
	//
	// example:
	//
	// 8123db90-880f-**************
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The application port number for the Tomcat container. Take note of the following items:
	//
	// 	- If this parameter is not specified when you configure the application instance group, the configuration of the application is applied.
	//
	// 	- If this parameter is not specified when you configure the application, the default port 8080 is applied.
	//
	// example:
	//
	// 8080
	HttpPort *int32 `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	// The maximum number of threads. Take note of the following items:
	//
	// 	- If this parameter is not specified when you configure the application instance group, the configuration of the application is applied.
	//
	// 	- If this parameter is not specified when you configure the application, the default value 250 is applied.
	//
	// example:
	//
	// 20
	MaxThreads *int32 `json:"MaxThreads,omitempty" xml:"MaxThreads,omitempty"`
	// The uniform resource identifier (URI) encoding scheme. Valid values: ISO-8859-1, GBK, GB2312, and UTF-8. Take note of the following items:
	//
	// 	- If this parameter is not specified when you configure the application instance group, the configuration of the application is applied.
	//
	// 	- If this parameter is not specified when you configure the application, the default URI encoding scheme in the Tomcat container is applied.
	//
	// example:
	//
	// ISO-8859-1
	URIEncoding *string `json:"URIEncoding,omitempty" xml:"URIEncoding,omitempty"`
	// Specifies whether to use the encoding scheme specified in the request body for URI query parameters. Take note of the following items:
	//
	// 	- If this parameter is not specified when you configure the application instance group, the configuration of the application is applied.
	//
	// 	- If this parameter is not specified when you configure the application, the default value false is applied.
	//
	// example:
	//
	// true
	UseBodyEncoding *bool `json:"UseBodyEncoding,omitempty" xml:"UseBodyEncoding,omitempty"`
}

func (s UpdateContainerConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContainerConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateContainerConfigurationRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateContainerConfigurationRequest) GetContextPath() *string {
	return s.ContextPath
}

func (s *UpdateContainerConfigurationRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateContainerConfigurationRequest) GetHttpPort() *int32 {
	return s.HttpPort
}

func (s *UpdateContainerConfigurationRequest) GetMaxThreads() *int32 {
	return s.MaxThreads
}

func (s *UpdateContainerConfigurationRequest) GetURIEncoding() *string {
	return s.URIEncoding
}

func (s *UpdateContainerConfigurationRequest) GetUseBodyEncoding() *bool {
	return s.UseBodyEncoding
}

func (s *UpdateContainerConfigurationRequest) SetAppId(v string) *UpdateContainerConfigurationRequest {
	s.AppId = &v
	return s
}

func (s *UpdateContainerConfigurationRequest) SetContextPath(v string) *UpdateContainerConfigurationRequest {
	s.ContextPath = &v
	return s
}

func (s *UpdateContainerConfigurationRequest) SetGroupId(v string) *UpdateContainerConfigurationRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateContainerConfigurationRequest) SetHttpPort(v int32) *UpdateContainerConfigurationRequest {
	s.HttpPort = &v
	return s
}

func (s *UpdateContainerConfigurationRequest) SetMaxThreads(v int32) *UpdateContainerConfigurationRequest {
	s.MaxThreads = &v
	return s
}

func (s *UpdateContainerConfigurationRequest) SetURIEncoding(v string) *UpdateContainerConfigurationRequest {
	s.URIEncoding = &v
	return s
}

func (s *UpdateContainerConfigurationRequest) SetUseBodyEncoding(v bool) *UpdateContainerConfigurationRequest {
	s.UseBodyEncoding = &v
	return s
}

func (s *UpdateContainerConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
