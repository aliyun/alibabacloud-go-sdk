// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyAgOneKeyOnlyCheckerTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApplyAgOneKeyOnlyCheckerTaskResponseBody
	GetCode() *string
	SetData(v string) *ApplyAgOneKeyOnlyCheckerTaskResponseBody
	GetData() *string
	SetMessage(v string) *ApplyAgOneKeyOnlyCheckerTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *ApplyAgOneKeyOnlyCheckerTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyAgOneKeyOnlyCheckerTaskResponseBody
	GetSuccess() *bool
}

type ApplyAgOneKeyOnlyCheckerTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApplyAgOneKeyOnlyCheckerTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyAgOneKeyOnlyCheckerTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyAgOneKeyOnlyCheckerTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApplyAgOneKeyOnlyCheckerTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *ApplyAgOneKeyOnlyCheckerTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyAgOneKeyOnlyCheckerTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyAgOneKeyOnlyCheckerTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyAgOneKeyOnlyCheckerTaskResponseBody) SetCode(v string) *ApplyAgOneKeyOnlyCheckerTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyAgOneKeyOnlyCheckerTaskResponseBody) SetData(v string) *ApplyAgOneKeyOnlyCheckerTaskResponseBody {
	s.Data = &v
	return s
}

func (s *ApplyAgOneKeyOnlyCheckerTaskResponseBody) SetMessage(v string) *ApplyAgOneKeyOnlyCheckerTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyAgOneKeyOnlyCheckerTaskResponseBody) SetRequestId(v string) *ApplyAgOneKeyOnlyCheckerTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyAgOneKeyOnlyCheckerTaskResponseBody) SetSuccess(v bool) *ApplyAgOneKeyOnlyCheckerTaskResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyAgOneKeyOnlyCheckerTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
