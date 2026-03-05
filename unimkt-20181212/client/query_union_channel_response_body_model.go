// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUnionChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryUnionChannelResponseBody
	GetCode() *int32
	SetData(v map[string]interface{}) *QueryUnionChannelResponseBody
	GetData() map[string]interface{}
	SetErrorMsg(v string) *QueryUnionChannelResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *QueryUnionChannelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryUnionChannelResponseBody
	GetSuccess() *bool
}

type QueryUnionChannelResponseBody struct {
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string                `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUnionChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUnionChannelResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUnionChannelResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryUnionChannelResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *QueryUnionChannelResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryUnionChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUnionChannelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUnionChannelResponseBody) SetCode(v int32) *QueryUnionChannelResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUnionChannelResponseBody) SetData(v map[string]interface{}) *QueryUnionChannelResponseBody {
	s.Data = v
	return s
}

func (s *QueryUnionChannelResponseBody) SetErrorMsg(v string) *QueryUnionChannelResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryUnionChannelResponseBody) SetRequestId(v string) *QueryUnionChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUnionChannelResponseBody) SetSuccess(v bool) *QueryUnionChannelResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUnionChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
