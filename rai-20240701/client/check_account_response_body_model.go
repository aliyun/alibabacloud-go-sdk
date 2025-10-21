// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckAccountResponseBody
	GetCode() *string
	SetData(v *CheckAccountResponseBodyData) *CheckAccountResponseBody
	GetData() *CheckAccountResponseBodyData
	SetHttpStatusCode(v int32) *CheckAccountResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CheckAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckAccountResponseBody
	GetSuccess() *bool
}

type CheckAccountResponseBody struct {
	// example:
	//
	// 00000
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CheckAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckAccountResponseBody) GetData() *CheckAccountResponseBodyData {
	return s.Data
}

func (s *CheckAccountResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CheckAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckAccountResponseBody) SetCode(v string) *CheckAccountResponseBody {
	s.Code = &v
	return s
}

func (s *CheckAccountResponseBody) SetData(v *CheckAccountResponseBodyData) *CheckAccountResponseBody {
	s.Data = v
	return s
}

func (s *CheckAccountResponseBody) SetHttpStatusCode(v int32) *CheckAccountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckAccountResponseBody) SetMessage(v string) *CheckAccountResponseBody {
	s.Message = &v
	return s
}

func (s *CheckAccountResponseBody) SetRequestId(v string) *CheckAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAccountResponseBody) SetSuccess(v bool) *CheckAccountResponseBody {
	s.Success = &v
	return s
}

func (s *CheckAccountResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckAccountResponseBodyData struct {
	// example:
	//
	// 1
	CheckResult *int32 `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
}

func (s CheckAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckAccountResponseBodyData) GetCheckResult() *int32 {
	return s.CheckResult
}

func (s *CheckAccountResponseBodyData) SetCheckResult(v int32) *CheckAccountResponseBodyData {
	s.CheckResult = &v
	return s
}

func (s *CheckAccountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
