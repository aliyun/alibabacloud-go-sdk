// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLlmAccessProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteLlmAccessProfileResponseBody
	GetCode() *string
	SetData(v string) *DeleteLlmAccessProfileResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteLlmAccessProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteLlmAccessProfileResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteLlmAccessProfileResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteLlmAccessProfileResponseBody
	GetRequestId() *string
}

type DeleteLlmAccessProfileResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66b
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 9DB8BA95-4513-54B9-9C67-A28909CFB4AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLlmAccessProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLlmAccessProfileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLlmAccessProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteLlmAccessProfileResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteLlmAccessProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteLlmAccessProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteLlmAccessProfileResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteLlmAccessProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLlmAccessProfileResponseBody) SetCode(v string) *DeleteLlmAccessProfileResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteLlmAccessProfileResponseBody) SetData(v string) *DeleteLlmAccessProfileResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteLlmAccessProfileResponseBody) SetHttpStatusCode(v int32) *DeleteLlmAccessProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteLlmAccessProfileResponseBody) SetMessage(v string) *DeleteLlmAccessProfileResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLlmAccessProfileResponseBody) SetParams(v []*string) *DeleteLlmAccessProfileResponseBody {
	s.Params = v
	return s
}

func (s *DeleteLlmAccessProfileResponseBody) SetRequestId(v string) *DeleteLlmAccessProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLlmAccessProfileResponseBody) Validate() error {
	return dara.Validate(s)
}
