// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryContentListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryContentListResponseBody
	GetCode() *int32
	SetData(v []map[string]interface{}) *QueryContentListResponseBody
	GetData() []map[string]interface{}
	SetErrorMsg(v string) *QueryContentListResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *QueryContentListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryContentListResponseBody
	GetSuccess() *bool
}

type QueryContentListResponseBody struct {
	Code      *int32                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg  *string                  `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryContentListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryContentListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryContentListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryContentListResponseBody) GetData() []map[string]interface{} {
	return s.Data
}

func (s *QueryContentListResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryContentListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryContentListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryContentListResponseBody) SetCode(v int32) *QueryContentListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryContentListResponseBody) SetData(v []map[string]interface{}) *QueryContentListResponseBody {
	s.Data = v
	return s
}

func (s *QueryContentListResponseBody) SetErrorMsg(v string) *QueryContentListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryContentListResponseBody) SetRequestId(v string) *QueryContentListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryContentListResponseBody) SetSuccess(v bool) *QueryContentListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryContentListResponseBody) Validate() error {
	return dara.Validate(s)
}
