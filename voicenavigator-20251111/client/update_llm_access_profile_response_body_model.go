// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLlmAccessProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateLlmAccessProfileResponseBody
	GetCode() *string
	SetData(v string) *UpdateLlmAccessProfileResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateLlmAccessProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateLlmAccessProfileResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateLlmAccessProfileResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateLlmAccessProfileResponseBody
	GetRequestId() *string
}

type UpdateLlmAccessProfileResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// e48e45dd-e47a-4744-a063-f08cbebb1c5b
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
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLlmAccessProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLlmAccessProfileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLlmAccessProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateLlmAccessProfileResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateLlmAccessProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateLlmAccessProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateLlmAccessProfileResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateLlmAccessProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLlmAccessProfileResponseBody) SetCode(v string) *UpdateLlmAccessProfileResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLlmAccessProfileResponseBody) SetData(v string) *UpdateLlmAccessProfileResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateLlmAccessProfileResponseBody) SetHttpStatusCode(v int32) *UpdateLlmAccessProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateLlmAccessProfileResponseBody) SetMessage(v string) *UpdateLlmAccessProfileResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLlmAccessProfileResponseBody) SetParams(v []*string) *UpdateLlmAccessProfileResponseBody {
	s.Params = v
	return s
}

func (s *UpdateLlmAccessProfileResponseBody) SetRequestId(v string) *UpdateLlmAccessProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLlmAccessProfileResponseBody) Validate() error {
	return dara.Validate(s)
}
