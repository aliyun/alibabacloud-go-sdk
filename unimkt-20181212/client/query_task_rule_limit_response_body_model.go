// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskRuleLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryTaskRuleLimitResponseBody
	GetCode() *int32
	SetData(v []map[string]interface{}) *QueryTaskRuleLimitResponseBody
	GetData() []map[string]interface{}
	SetErrorMsg(v string) *QueryTaskRuleLimitResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *QueryTaskRuleLimitResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryTaskRuleLimitResponseBody
	GetSuccess() *bool
}

type QueryTaskRuleLimitResponseBody struct {
	Code      *int32                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg  *string                  `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTaskRuleLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskRuleLimitResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskRuleLimitResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryTaskRuleLimitResponseBody) GetData() []map[string]interface{} {
	return s.Data
}

func (s *QueryTaskRuleLimitResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryTaskRuleLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTaskRuleLimitResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryTaskRuleLimitResponseBody) SetCode(v int32) *QueryTaskRuleLimitResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTaskRuleLimitResponseBody) SetData(v []map[string]interface{}) *QueryTaskRuleLimitResponseBody {
	s.Data = v
	return s
}

func (s *QueryTaskRuleLimitResponseBody) SetErrorMsg(v string) *QueryTaskRuleLimitResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryTaskRuleLimitResponseBody) SetRequestId(v string) *QueryTaskRuleLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTaskRuleLimitResponseBody) SetSuccess(v bool) *QueryTaskRuleLimitResponseBody {
	s.Success = &v
	return s
}

func (s *QueryTaskRuleLimitResponseBody) Validate() error {
	return dara.Validate(s)
}
