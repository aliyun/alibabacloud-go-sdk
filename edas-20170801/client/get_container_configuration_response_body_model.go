// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContainerConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetContainerConfigurationResponseBody
	GetCode() *int32
	SetContainerConfiguration(v *GetContainerConfigurationResponseBodyContainerConfiguration) *GetContainerConfigurationResponseBody
	GetContainerConfiguration() *GetContainerConfigurationResponseBodyContainerConfiguration
	SetMessage(v string) *GetContainerConfigurationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetContainerConfigurationResponseBody
	GetRequestId() *string
}

type GetContainerConfigurationResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The Tomcat configuration.
	ContainerConfiguration *GetContainerConfigurationResponseBodyContainerConfiguration `json:"ContainerConfiguration,omitempty" xml:"ContainerConfiguration,omitempty" type:"Struct"`
	// The message returned for the request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 34F8-FDG9-*****************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetContainerConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContainerConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetContainerConfigurationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetContainerConfigurationResponseBody) GetContainerConfiguration() *GetContainerConfigurationResponseBodyContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *GetContainerConfigurationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetContainerConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContainerConfigurationResponseBody) SetCode(v int32) *GetContainerConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *GetContainerConfigurationResponseBody) SetContainerConfiguration(v *GetContainerConfigurationResponseBodyContainerConfiguration) *GetContainerConfigurationResponseBody {
	s.ContainerConfiguration = v
	return s
}

func (s *GetContainerConfigurationResponseBody) SetMessage(v string) *GetContainerConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *GetContainerConfigurationResponseBody) SetRequestId(v string) *GetContainerConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContainerConfigurationResponseBody) Validate() error {
	if s.ContainerConfiguration != nil {
		if err := s.ContainerConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContainerConfigurationResponseBodyContainerConfiguration struct {
	// The context path of the Tomcat container.
	//
	// example:
	//
	// /
	ContextPath *string `json:"ContextPath,omitempty" xml:"ContextPath,omitempty"`
	// The application port number for the Tomcat container. The value specified in the application configuration is returned.
	//
	// example:
	//
	// 8080
	HttpPort *int32 `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	// The maximum number of threads in the Tomcat container.
	//
	// 	- If no instance group is specified, the configuration of the application is returned.
	//
	// 	- If no application is specified, the default configuration is returned.
	//
	// example:
	//
	// 400
	MaxThreads *int32 `json:"MaxThreads,omitempty" xml:"MaxThreads,omitempty"`
	// The Uniform Resource Identifier (URI) encoding scheme. Valid values: ISO-8859-1, GBK, GB2312, and UTF-8.
	//
	// 	- If no instance group is specified, the configuration of the application is returned.
	//
	// 	- If no application is specified, the default configuration is returned.
	//
	// example:
	//
	// ISO-8859-1
	URIEncoding *string `json:"URIEncoding,omitempty" xml:"URIEncoding,omitempty"`
	// Indicates whether useBodyEncodingForURI is enabled in the Tomcat container.
	//
	// 	- If no instance group is specified, the configuration of the application is returned.
	//
	// 	- If no application is specified, the default configuration is returned.
	//
	// example:
	//
	// true
	UseBodyEncoding *bool `json:"UseBodyEncoding,omitempty" xml:"UseBodyEncoding,omitempty"`
}

func (s GetContainerConfigurationResponseBodyContainerConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetContainerConfigurationResponseBodyContainerConfiguration) GoString() string {
	return s.String()
}

func (s *GetContainerConfigurationResponseBodyContainerConfiguration) GetContextPath() *string {
	return s.ContextPath
}

func (s *GetContainerConfigurationResponseBodyContainerConfiguration) GetHttpPort() *int32 {
	return s.HttpPort
}

func (s *GetContainerConfigurationResponseBodyContainerConfiguration) GetMaxThreads() *int32 {
	return s.MaxThreads
}

func (s *GetContainerConfigurationResponseBodyContainerConfiguration) GetURIEncoding() *string {
	return s.URIEncoding
}

func (s *GetContainerConfigurationResponseBodyContainerConfiguration) GetUseBodyEncoding() *bool {
	return s.UseBodyEncoding
}

func (s *GetContainerConfigurationResponseBodyContainerConfiguration) SetContextPath(v string) *GetContainerConfigurationResponseBodyContainerConfiguration {
	s.ContextPath = &v
	return s
}

func (s *GetContainerConfigurationResponseBodyContainerConfiguration) SetHttpPort(v int32) *GetContainerConfigurationResponseBodyContainerConfiguration {
	s.HttpPort = &v
	return s
}

func (s *GetContainerConfigurationResponseBodyContainerConfiguration) SetMaxThreads(v int32) *GetContainerConfigurationResponseBodyContainerConfiguration {
	s.MaxThreads = &v
	return s
}

func (s *GetContainerConfigurationResponseBodyContainerConfiguration) SetURIEncoding(v string) *GetContainerConfigurationResponseBodyContainerConfiguration {
	s.URIEncoding = &v
	return s
}

func (s *GetContainerConfigurationResponseBodyContainerConfiguration) SetUseBodyEncoding(v bool) *GetContainerConfigurationResponseBodyContainerConfiguration {
	s.UseBodyEncoding = &v
	return s
}

func (s *GetContainerConfigurationResponseBodyContainerConfiguration) Validate() error {
	return dara.Validate(s)
}
