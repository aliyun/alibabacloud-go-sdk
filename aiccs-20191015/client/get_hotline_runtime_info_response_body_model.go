// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineRuntimeInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotlineRuntimeInfoResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *GetHotlineRuntimeInfoResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *GetHotlineRuntimeInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotlineRuntimeInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHotlineRuntimeInfoResponseBody
	GetSuccess() *bool
}

type GetHotlineRuntimeInfoResponseBody struct {
	// example:
	//
	// 200
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotlineRuntimeInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineRuntimeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineRuntimeInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotlineRuntimeInfoResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetHotlineRuntimeInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotlineRuntimeInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotlineRuntimeInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHotlineRuntimeInfoResponseBody) SetCode(v string) *GetHotlineRuntimeInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineRuntimeInfoResponseBody) SetData(v map[string]interface{}) *GetHotlineRuntimeInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetHotlineRuntimeInfoResponseBody) SetMessage(v string) *GetHotlineRuntimeInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineRuntimeInfoResponseBody) SetRequestId(v string) *GetHotlineRuntimeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineRuntimeInfoResponseBody) SetSuccess(v bool) *GetHotlineRuntimeInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineRuntimeInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
