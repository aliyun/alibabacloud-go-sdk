// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJudgeAgExistQuietPeriodResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *JudgeAgExistQuietPeriodResponseBody
	GetCode() *string
	SetData(v bool) *JudgeAgExistQuietPeriodResponseBody
	GetData() *bool
	SetMessage(v string) *JudgeAgExistQuietPeriodResponseBody
	GetMessage() *string
	SetRequestId(v string) *JudgeAgExistQuietPeriodResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *JudgeAgExistQuietPeriodResponseBody
	GetSuccess() *bool
}

type JudgeAgExistQuietPeriodResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s JudgeAgExistQuietPeriodResponseBody) String() string {
	return dara.Prettify(s)
}

func (s JudgeAgExistQuietPeriodResponseBody) GoString() string {
	return s.String()
}

func (s *JudgeAgExistQuietPeriodResponseBody) GetCode() *string {
	return s.Code
}

func (s *JudgeAgExistQuietPeriodResponseBody) GetData() *bool {
	return s.Data
}

func (s *JudgeAgExistQuietPeriodResponseBody) GetMessage() *string {
	return s.Message
}

func (s *JudgeAgExistQuietPeriodResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *JudgeAgExistQuietPeriodResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *JudgeAgExistQuietPeriodResponseBody) SetCode(v string) *JudgeAgExistQuietPeriodResponseBody {
	s.Code = &v
	return s
}

func (s *JudgeAgExistQuietPeriodResponseBody) SetData(v bool) *JudgeAgExistQuietPeriodResponseBody {
	s.Data = &v
	return s
}

func (s *JudgeAgExistQuietPeriodResponseBody) SetMessage(v string) *JudgeAgExistQuietPeriodResponseBody {
	s.Message = &v
	return s
}

func (s *JudgeAgExistQuietPeriodResponseBody) SetRequestId(v string) *JudgeAgExistQuietPeriodResponseBody {
	s.RequestId = &v
	return s
}

func (s *JudgeAgExistQuietPeriodResponseBody) SetSuccess(v bool) *JudgeAgExistQuietPeriodResponseBody {
	s.Success = &v
	return s
}

func (s *JudgeAgExistQuietPeriodResponseBody) Validate() error {
	return dara.Validate(s)
}
