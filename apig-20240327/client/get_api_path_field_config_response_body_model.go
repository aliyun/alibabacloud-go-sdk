// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiPathFieldConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetApiPathFieldConfigResponseBody
	GetCode() *string
	SetData(v *GetApiPathFieldConfigResponseBodyData) *GetApiPathFieldConfigResponseBody
	GetData() *GetApiPathFieldConfigResponseBodyData
	SetMessage(v string) *GetApiPathFieldConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetApiPathFieldConfigResponseBody
	GetRequestId() *string
}

type GetApiPathFieldConfigResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetApiPathFieldConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetApiPathFieldConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApiPathFieldConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetApiPathFieldConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetApiPathFieldConfigResponseBody) GetData() *GetApiPathFieldConfigResponseBodyData {
	return s.Data
}

func (s *GetApiPathFieldConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetApiPathFieldConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApiPathFieldConfigResponseBody) SetCode(v string) *GetApiPathFieldConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetApiPathFieldConfigResponseBody) SetData(v *GetApiPathFieldConfigResponseBodyData) *GetApiPathFieldConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetApiPathFieldConfigResponseBody) SetMessage(v string) *GetApiPathFieldConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetApiPathFieldConfigResponseBody) SetRequestId(v string) *GetApiPathFieldConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApiPathFieldConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApiPathFieldConfigResponseBodyData struct {
	// The ID of the HTTP API.
	//
	// example:
	//
	// api-123456
	ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// The ID of the cloud-native API gateway.
	//
	// example:
	//
	// gw-123456
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The list of field configurations grouped by API request path.
	PathConfigs []*GetApiPathFieldConfigResponseBodyDataPathConfigs `json:"pathConfigs,omitempty" xml:"pathConfigs,omitempty" type:"Repeated"`
}

func (s GetApiPathFieldConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetApiPathFieldConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetApiPathFieldConfigResponseBodyData) GetApiId() *string {
	return s.ApiId
}

func (s *GetApiPathFieldConfigResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetApiPathFieldConfigResponseBodyData) GetPathConfigs() []*GetApiPathFieldConfigResponseBodyDataPathConfigs {
	return s.PathConfigs
}

func (s *GetApiPathFieldConfigResponseBodyData) SetApiId(v string) *GetApiPathFieldConfigResponseBodyData {
	s.ApiId = &v
	return s
}

func (s *GetApiPathFieldConfigResponseBodyData) SetGatewayId(v string) *GetApiPathFieldConfigResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetApiPathFieldConfigResponseBodyData) SetPathConfigs(v []*GetApiPathFieldConfigResponseBodyDataPathConfigs) *GetApiPathFieldConfigResponseBodyData {
	s.PathConfigs = v
	return s
}

func (s *GetApiPathFieldConfigResponseBodyData) Validate() error {
	if s.PathConfigs != nil {
		for _, item := range s.PathConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApiPathFieldConfigResponseBodyDataPathConfigs struct {
	// The collection of field configurations under the path, grouped by mapping key, where each group represents a set of field configuration items.
	FieldPaths map[string][]*DataPathConfigsFieldPathsValue `json:"fieldPaths,omitempty" xml:"fieldPaths,omitempty"`
	// The request path of the API.
	//
	// example:
	//
	// /v1/chat/completions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s GetApiPathFieldConfigResponseBodyDataPathConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetApiPathFieldConfigResponseBodyDataPathConfigs) GoString() string {
	return s.String()
}

func (s *GetApiPathFieldConfigResponseBodyDataPathConfigs) GetFieldPaths() map[string][]*DataPathConfigsFieldPathsValue {
	return s.FieldPaths
}

func (s *GetApiPathFieldConfigResponseBodyDataPathConfigs) GetPath() *string {
	return s.Path
}

func (s *GetApiPathFieldConfigResponseBodyDataPathConfigs) SetFieldPaths(v map[string][]*DataPathConfigsFieldPathsValue) *GetApiPathFieldConfigResponseBodyDataPathConfigs {
	s.FieldPaths = v
	return s
}

func (s *GetApiPathFieldConfigResponseBodyDataPathConfigs) SetPath(v string) *GetApiPathFieldConfigResponseBodyDataPathConfigs {
	s.Path = &v
	return s
}

func (s *GetApiPathFieldConfigResponseBodyDataPathConfigs) Validate() error {
	return dara.Validate(s)
}
