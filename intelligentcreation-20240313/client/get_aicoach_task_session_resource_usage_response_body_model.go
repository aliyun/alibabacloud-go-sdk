// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachTaskSessionResourceUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAudioUsage(v int32) *GetAICoachTaskSessionResourceUsageResponseBody
	GetAudioUsage() *int32
	SetDeductionStatus(v int32) *GetAICoachTaskSessionResourceUsageResponseBody
	GetDeductionStatus() *int32
	SetErrorCode(v string) *GetAICoachTaskSessionResourceUsageResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetAICoachTaskSessionResourceUsageResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetAICoachTaskSessionResourceUsageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAICoachTaskSessionResourceUsageResponseBody
	GetSuccess() *bool
	SetTokenUsage(v int32) *GetAICoachTaskSessionResourceUsageResponseBody
	GetTokenUsage() *int32
}

type GetAICoachTaskSessionResourceUsageResponseBody struct {
	// example:
	//
	// 60
	AudioUsage *int32 `json:"audioUsage,omitempty" xml:"audioUsage,omitempty"`
	// example:
	//
	// 0：待扣费；1：完成扣费
	DeductionStatus *int32 `json:"deductionStatus,omitempty" xml:"deductionStatus,omitempty"`
	// example:
	//
	// InternalServerError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 系统异常，请联系管理员
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4D902811-B75C-5D1B-8882-D515F8E2F977
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 1000
	TokenUsage *int32 `json:"tokenUsage,omitempty" xml:"tokenUsage,omitempty"`
}

func (s GetAICoachTaskSessionResourceUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachTaskSessionResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetAICoachTaskSessionResourceUsageResponseBody) GetAudioUsage() *int32 {
	return s.AudioUsage
}

func (s *GetAICoachTaskSessionResourceUsageResponseBody) GetDeductionStatus() *int32 {
	return s.DeductionStatus
}

func (s *GetAICoachTaskSessionResourceUsageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAICoachTaskSessionResourceUsageResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAICoachTaskSessionResourceUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAICoachTaskSessionResourceUsageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAICoachTaskSessionResourceUsageResponseBody) GetTokenUsage() *int32 {
	return s.TokenUsage
}

func (s *GetAICoachTaskSessionResourceUsageResponseBody) SetAudioUsage(v int32) *GetAICoachTaskSessionResourceUsageResponseBody {
	s.AudioUsage = &v
	return s
}

func (s *GetAICoachTaskSessionResourceUsageResponseBody) SetDeductionStatus(v int32) *GetAICoachTaskSessionResourceUsageResponseBody {
	s.DeductionStatus = &v
	return s
}

func (s *GetAICoachTaskSessionResourceUsageResponseBody) SetErrorCode(v string) *GetAICoachTaskSessionResourceUsageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAICoachTaskSessionResourceUsageResponseBody) SetErrorMessage(v string) *GetAICoachTaskSessionResourceUsageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAICoachTaskSessionResourceUsageResponseBody) SetRequestId(v string) *GetAICoachTaskSessionResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAICoachTaskSessionResourceUsageResponseBody) SetSuccess(v bool) *GetAICoachTaskSessionResourceUsageResponseBody {
	s.Success = &v
	return s
}

func (s *GetAICoachTaskSessionResourceUsageResponseBody) SetTokenUsage(v int32) *GetAICoachTaskSessionResourceUsageResponseBody {
	s.TokenUsage = &v
	return s
}

func (s *GetAICoachTaskSessionResourceUsageResponseBody) Validate() error {
	return dara.Validate(s)
}
