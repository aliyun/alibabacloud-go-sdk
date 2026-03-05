// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUnionTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryUnionTaskListResponseBody
	GetCode() *int32
	SetData(v map[string]interface{}) *QueryUnionTaskListResponseBody
	GetData() map[string]interface{}
	SetErrorMsg(v string) *QueryUnionTaskListResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *QueryUnionTaskListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryUnionTaskListResponseBody
	GetSuccess() *bool
}

type QueryUnionTaskListResponseBody struct {
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string                `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUnionTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUnionTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUnionTaskListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryUnionTaskListResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *QueryUnionTaskListResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryUnionTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUnionTaskListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUnionTaskListResponseBody) SetCode(v int32) *QueryUnionTaskListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUnionTaskListResponseBody) SetData(v map[string]interface{}) *QueryUnionTaskListResponseBody {
	s.Data = v
	return s
}

func (s *QueryUnionTaskListResponseBody) SetErrorMsg(v string) *QueryUnionTaskListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryUnionTaskListResponseBody) SetRequestId(v string) *QueryUnionTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUnionTaskListResponseBody) SetSuccess(v bool) *QueryUnionTaskListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUnionTaskListResponseBody) Validate() error {
	return dara.Validate(s)
}
