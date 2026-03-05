// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUnionContentInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryUnionContentInfoResponseBody
	GetCode() *int32
	SetData(v map[string]interface{}) *QueryUnionContentInfoResponseBody
	GetData() map[string]interface{}
	SetErrorMsg(v string) *QueryUnionContentInfoResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *QueryUnionContentInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryUnionContentInfoResponseBody
	GetSuccess() *bool
}

type QueryUnionContentInfoResponseBody struct {
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string                `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUnionContentInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUnionContentInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUnionContentInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryUnionContentInfoResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *QueryUnionContentInfoResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryUnionContentInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUnionContentInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUnionContentInfoResponseBody) SetCode(v int32) *QueryUnionContentInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUnionContentInfoResponseBody) SetData(v map[string]interface{}) *QueryUnionContentInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryUnionContentInfoResponseBody) SetErrorMsg(v string) *QueryUnionContentInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryUnionContentInfoResponseBody) SetRequestId(v string) *QueryUnionContentInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUnionContentInfoResponseBody) SetSuccess(v bool) *QueryUnionContentInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUnionContentInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
