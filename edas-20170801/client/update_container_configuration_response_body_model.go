// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContainerConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateContainerConfigurationResponseBody
	GetCode() *int32
	SetContainerConfiguration(v *UpdateContainerConfigurationResponseBodyContainerConfiguration) *UpdateContainerConfigurationResponseBody
	GetContainerConfiguration() *UpdateContainerConfigurationResponseBodyContainerConfiguration
	SetMessage(v string) *UpdateContainerConfigurationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateContainerConfigurationResponseBody
	GetRequestId() *string
}

type UpdateContainerConfigurationResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configuration of the Tomcat container.
	ContainerConfiguration *UpdateContainerConfigurationResponseBodyContainerConfiguration `json:"ContainerConfiguration,omitempty" xml:"ContainerConfiguration,omitempty" type:"Struct"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D16979DC-4D42-***************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateContainerConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContainerConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContainerConfigurationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateContainerConfigurationResponseBody) GetContainerConfiguration() *UpdateContainerConfigurationResponseBodyContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *UpdateContainerConfigurationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateContainerConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContainerConfigurationResponseBody) SetCode(v int32) *UpdateContainerConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateContainerConfigurationResponseBody) SetContainerConfiguration(v *UpdateContainerConfigurationResponseBodyContainerConfiguration) *UpdateContainerConfigurationResponseBody {
	s.ContainerConfiguration = v
	return s
}

func (s *UpdateContainerConfigurationResponseBody) SetMessage(v string) *UpdateContainerConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateContainerConfigurationResponseBody) SetRequestId(v string) *UpdateContainerConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContainerConfigurationResponseBody) Validate() error {
	if s.ContainerConfiguration != nil {
		if err := s.ContainerConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateContainerConfigurationResponseBodyContainerConfiguration struct {
	// The context path of the Tomcat container.
	//
	// example:
	//
	// /
	ContextPath *string `json:"ContextPath,omitempty" xml:"ContextPath,omitempty"`
	// The application port number for the Tomcat container.
	//
	// example:
	//
	// 8080
	HttpPort *int32 `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	// The maximum number of threads.
	//
	// example:
	//
	// 20
	MaxThreads *int32 `json:"MaxThreads,omitempty" xml:"MaxThreads,omitempty"`
	// The URI encoding scheme. Valid values: ISO-8859-1, GBK, GB2312, and UTF-8.
	//
	// example:
	//
	// ISO-8859-1
	URIEncoding *string `json:"URIEncoding,omitempty" xml:"URIEncoding,omitempty"`
	// Indicates whether useBodyEncodingForURI is enabled.
	//
	// example:
	//
	// true
	UseBodyEncoding *bool `json:"UseBodyEncoding,omitempty" xml:"UseBodyEncoding,omitempty"`
}

func (s UpdateContainerConfigurationResponseBodyContainerConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateContainerConfigurationResponseBodyContainerConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateContainerConfigurationResponseBodyContainerConfiguration) GetContextPath() *string {
	return s.ContextPath
}

func (s *UpdateContainerConfigurationResponseBodyContainerConfiguration) GetHttpPort() *int32 {
	return s.HttpPort
}

func (s *UpdateContainerConfigurationResponseBodyContainerConfiguration) GetMaxThreads() *int32 {
	return s.MaxThreads
}

func (s *UpdateContainerConfigurationResponseBodyContainerConfiguration) GetURIEncoding() *string {
	return s.URIEncoding
}

func (s *UpdateContainerConfigurationResponseBodyContainerConfiguration) GetUseBodyEncoding() *bool {
	return s.UseBodyEncoding
}

func (s *UpdateContainerConfigurationResponseBodyContainerConfiguration) SetContextPath(v string) *UpdateContainerConfigurationResponseBodyContainerConfiguration {
	s.ContextPath = &v
	return s
}

func (s *UpdateContainerConfigurationResponseBodyContainerConfiguration) SetHttpPort(v int32) *UpdateContainerConfigurationResponseBodyContainerConfiguration {
	s.HttpPort = &v
	return s
}

func (s *UpdateContainerConfigurationResponseBodyContainerConfiguration) SetMaxThreads(v int32) *UpdateContainerConfigurationResponseBodyContainerConfiguration {
	s.MaxThreads = &v
	return s
}

func (s *UpdateContainerConfigurationResponseBodyContainerConfiguration) SetURIEncoding(v string) *UpdateContainerConfigurationResponseBodyContainerConfiguration {
	s.URIEncoding = &v
	return s
}

func (s *UpdateContainerConfigurationResponseBodyContainerConfiguration) SetUseBodyEncoding(v bool) *UpdateContainerConfigurationResponseBodyContainerConfiguration {
	s.UseBodyEncoding = &v
	return s
}

func (s *UpdateContainerConfigurationResponseBodyContainerConfiguration) Validate() error {
	return dara.Validate(s)
}
