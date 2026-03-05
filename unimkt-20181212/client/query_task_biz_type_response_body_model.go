// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskBizTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryTaskBizTypeResponseBody
	GetCode() *int32
	SetData(v []map[string]interface{}) *QueryTaskBizTypeResponseBody
	GetData() []map[string]interface{}
	SetErrorMsg(v string) *QueryTaskBizTypeResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *QueryTaskBizTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryTaskBizTypeResponseBody
	GetSuccess() *bool
}

type QueryTaskBizTypeResponseBody struct {
	Code      *int32                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg  *string                  `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTaskBizTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskBizTypeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskBizTypeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryTaskBizTypeResponseBody) GetData() []map[string]interface{} {
	return s.Data
}

func (s *QueryTaskBizTypeResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryTaskBizTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTaskBizTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryTaskBizTypeResponseBody) SetCode(v int32) *QueryTaskBizTypeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTaskBizTypeResponseBody) SetData(v []map[string]interface{}) *QueryTaskBizTypeResponseBody {
	s.Data = v
	return s
}

func (s *QueryTaskBizTypeResponseBody) SetErrorMsg(v string) *QueryTaskBizTypeResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryTaskBizTypeResponseBody) SetRequestId(v string) *QueryTaskBizTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTaskBizTypeResponseBody) SetSuccess(v bool) *QueryTaskBizTypeResponseBody {
	s.Success = &v
	return s
}

func (s *QueryTaskBizTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
