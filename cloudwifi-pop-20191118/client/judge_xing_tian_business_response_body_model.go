// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJudgeXingTianBusinessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *JudgeXingTianBusinessResponseBody
	GetData() *bool
	SetErrorCode(v int32) *JudgeXingTianBusinessResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *JudgeXingTianBusinessResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *JudgeXingTianBusinessResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *JudgeXingTianBusinessResponseBody
	GetRequestId() *string
}

type JudgeXingTianBusinessResponseBody struct {
	Data         *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s JudgeXingTianBusinessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s JudgeXingTianBusinessResponseBody) GoString() string {
	return s.String()
}

func (s *JudgeXingTianBusinessResponseBody) GetData() *bool {
	return s.Data
}

func (s *JudgeXingTianBusinessResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *JudgeXingTianBusinessResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *JudgeXingTianBusinessResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *JudgeXingTianBusinessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *JudgeXingTianBusinessResponseBody) SetData(v bool) *JudgeXingTianBusinessResponseBody {
	s.Data = &v
	return s
}

func (s *JudgeXingTianBusinessResponseBody) SetErrorCode(v int32) *JudgeXingTianBusinessResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *JudgeXingTianBusinessResponseBody) SetErrorMessage(v string) *JudgeXingTianBusinessResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *JudgeXingTianBusinessResponseBody) SetIsSuccess(v bool) *JudgeXingTianBusinessResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *JudgeXingTianBusinessResponseBody) SetRequestId(v string) *JudgeXingTianBusinessResponseBody {
	s.RequestId = &v
	return s
}

func (s *JudgeXingTianBusinessResponseBody) Validate() error {
	return dara.Validate(s)
}
