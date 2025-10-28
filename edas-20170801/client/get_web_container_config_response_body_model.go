// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWebContainerConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetWebContainerConfigResponseBody
	GetCode() *int32
	SetMessage(v string) *GetWebContainerConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWebContainerConfigResponseBody
	GetRequestId() *string
	SetWebContainerConfig(v *GetWebContainerConfigResponseBodyWebContainerConfig) *GetWebContainerConfigResponseBody
	GetWebContainerConfig() *GetWebContainerConfigResponseBodyWebContainerConfig
}

type GetWebContainerConfigResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// 4823-bhjf-23u4-eiufh
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Tomcat configurations of the application.
	WebContainerConfig *GetWebContainerConfigResponseBodyWebContainerConfig `json:"WebContainerConfig,omitempty" xml:"WebContainerConfig,omitempty" type:"Struct"`
}

func (s GetWebContainerConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWebContainerConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebContainerConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetWebContainerConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWebContainerConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWebContainerConfigResponseBody) GetWebContainerConfig() *GetWebContainerConfigResponseBodyWebContainerConfig {
	return s.WebContainerConfig
}

func (s *GetWebContainerConfigResponseBody) SetCode(v int32) *GetWebContainerConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetWebContainerConfigResponseBody) SetMessage(v string) *GetWebContainerConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetWebContainerConfigResponseBody) SetRequestId(v string) *GetWebContainerConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWebContainerConfigResponseBody) SetWebContainerConfig(v *GetWebContainerConfigResponseBodyWebContainerConfig) *GetWebContainerConfigResponseBody {
	s.WebContainerConfig = v
	return s
}

func (s *GetWebContainerConfigResponseBody) Validate() error {
	if s.WebContainerConfig != nil {
		if err := s.WebContainerConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWebContainerConfigResponseBodyWebContainerConfig struct {
	// The type of the context path.
	//
	// example:
	//
	// root
	ContextInputType *string `json:"ContextInputType,omitempty" xml:"ContextInputType,omitempty"`
	// The context path.
	//
	// example:
	//
	// ROOT
	ContextPath *string `json:"ContextPath,omitempty" xml:"ContextPath,omitempty"`
	// The HTTP service port.
	//
	// example:
	//
	// 8080
	HttpPort *int32 `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	// The maximum number of threads.
	//
	// example:
	//
	// 500
	MaxThreads *int32 `json:"MaxThreads,omitempty" xml:"MaxThreads,omitempty"`
	// The content of the server.xml file customized by using advanced configurations.
	//
	// example:
	//
	// <Server port=*****
	ServerXml *string `json:"ServerXml,omitempty" xml:"ServerXml,omitempty"`
	// The URI encoding scheme.
	//
	// example:
	//
	// ISO-8859-1
	UriEncoding *string `json:"UriEncoding,omitempty" xml:"UriEncoding,omitempty"`
	// Indicates whether advanced configurations are used to customize the server.xml file.
	//
	// example:
	//
	// true
	UseAdvancedServerXml *bool `json:"UseAdvancedServerXml,omitempty" xml:"UseAdvancedServerXml,omitempty"`
	// Indicates whether the encoding scheme specified in the request body is used for uniform resource identifier (URI) query parameters.
	//
	// example:
	//
	// true
	UseBodyEncoding *bool `json:"UseBodyEncoding,omitempty" xml:"UseBodyEncoding,omitempty"`
	// Indicates whether the default configurations are used.
	//
	// example:
	//
	// true
	UseDefaultConfig *bool `json:"UseDefaultConfig,omitempty" xml:"UseDefaultConfig,omitempty"`
}

func (s GetWebContainerConfigResponseBodyWebContainerConfig) String() string {
	return dara.Prettify(s)
}

func (s GetWebContainerConfigResponseBodyWebContainerConfig) GoString() string {
	return s.String()
}

func (s *GetWebContainerConfigResponseBodyWebContainerConfig) GetContextInputType() *string {
	return s.ContextInputType
}

func (s *GetWebContainerConfigResponseBodyWebContainerConfig) GetContextPath() *string {
	return s.ContextPath
}

func (s *GetWebContainerConfigResponseBodyWebContainerConfig) GetHttpPort() *int32 {
	return s.HttpPort
}

func (s *GetWebContainerConfigResponseBodyWebContainerConfig) GetMaxThreads() *int32 {
	return s.MaxThreads
}

func (s *GetWebContainerConfigResponseBodyWebContainerConfig) GetServerXml() *string {
	return s.ServerXml
}

func (s *GetWebContainerConfigResponseBodyWebContainerConfig) GetUriEncoding() *string {
	return s.UriEncoding
}

func (s *GetWebContainerConfigResponseBodyWebContainerConfig) GetUseAdvancedServerXml() *bool {
	return s.UseAdvancedServerXml
}

func (s *GetWebContainerConfigResponseBodyWebContainerConfig) GetUseBodyEncoding() *bool {
	return s.UseBodyEncoding
}

func (s *GetWebContainerConfigResponseBodyWebContainerConfig) GetUseDefaultConfig() *bool {
	return s.UseDefaultConfig
}

func (s *GetWebContainerConfigResponseBodyWebContainerConfig) SetContextInputType(v string) *GetWebContainerConfigResponseBodyWebContainerConfig {
	s.ContextInputType = &v
	return s
}

func (s *GetWebContainerConfigResponseBodyWebContainerConfig) SetContextPath(v string) *GetWebContainerConfigResponseBodyWebContainerConfig {
	s.ContextPath = &v
	return s
}

func (s *GetWebContainerConfigResponseBodyWebContainerConfig) SetHttpPort(v int32) *GetWebContainerConfigResponseBodyWebContainerConfig {
	s.HttpPort = &v
	return s
}

func (s *GetWebContainerConfigResponseBodyWebContainerConfig) SetMaxThreads(v int32) *GetWebContainerConfigResponseBodyWebContainerConfig {
	s.MaxThreads = &v
	return s
}

func (s *GetWebContainerConfigResponseBodyWebContainerConfig) SetServerXml(v string) *GetWebContainerConfigResponseBodyWebContainerConfig {
	s.ServerXml = &v
	return s
}

func (s *GetWebContainerConfigResponseBodyWebContainerConfig) SetUriEncoding(v string) *GetWebContainerConfigResponseBodyWebContainerConfig {
	s.UriEncoding = &v
	return s
}

func (s *GetWebContainerConfigResponseBodyWebContainerConfig) SetUseAdvancedServerXml(v bool) *GetWebContainerConfigResponseBodyWebContainerConfig {
	s.UseAdvancedServerXml = &v
	return s
}

func (s *GetWebContainerConfigResponseBodyWebContainerConfig) SetUseBodyEncoding(v bool) *GetWebContainerConfigResponseBodyWebContainerConfig {
	s.UseBodyEncoding = &v
	return s
}

func (s *GetWebContainerConfigResponseBodyWebContainerConfig) SetUseDefaultConfig(v bool) *GetWebContainerConfigResponseBodyWebContainerConfig {
	s.UseDefaultConfig = &v
	return s
}

func (s *GetWebContainerConfigResponseBodyWebContainerConfig) Validate() error {
	return dara.Validate(s)
}
