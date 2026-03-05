// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUnionTaskInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryUnionTaskInfoResponseBody
	GetCode() *int32
	SetData(v map[string]interface{}) *QueryUnionTaskInfoResponseBody
	GetData() map[string]interface{}
	SetErrorMsg(v string) *QueryUnionTaskInfoResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *QueryUnionTaskInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryUnionTaskInfoResponseBody
	GetSuccess() *bool
}

type QueryUnionTaskInfoResponseBody struct {
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string                `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUnionTaskInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUnionTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUnionTaskInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryUnionTaskInfoResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *QueryUnionTaskInfoResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryUnionTaskInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUnionTaskInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUnionTaskInfoResponseBody) SetCode(v int32) *QueryUnionTaskInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUnionTaskInfoResponseBody) SetData(v map[string]interface{}) *QueryUnionTaskInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryUnionTaskInfoResponseBody) SetErrorMsg(v string) *QueryUnionTaskInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryUnionTaskInfoResponseBody) SetRequestId(v string) *QueryUnionTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUnionTaskInfoResponseBody) SetSuccess(v bool) *QueryUnionTaskInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUnionTaskInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
