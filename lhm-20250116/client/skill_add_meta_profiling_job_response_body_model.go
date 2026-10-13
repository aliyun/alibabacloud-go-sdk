// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkillAddMetaProfilingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *SkillAddMetaProfilingJobResponseBody
	GetData() *int64
	SetErrCode(v string) *SkillAddMetaProfilingJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *SkillAddMetaProfilingJobResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *SkillAddMetaProfilingJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SkillAddMetaProfilingJobResponseBody
	GetSuccess() *bool
}

type SkillAddMetaProfilingJobResponseBody struct {
	// The ID of the created or updated metadata profiling task.
	//
	// example:
	//
	// 100256
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// The error code. This parameter is an empty string if the call is successful.
	//
	// example:
	//
	// InvalidParameter
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. This parameter is an empty string if the call is successful.
	//
	// example:
	//
	// The parameter is invalid or missing.
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues for this call.
	//
	// example:
	//
	// 3F45B8C1-2A6E-4C9B-9E1F-7D2A1B0C5E66
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success, and a value of false indicates failure. If the call fails, troubleshoot the issue based on the errCode and errMessage parameters.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SkillAddMetaProfilingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SkillAddMetaProfilingJobResponseBody) GoString() string {
	return s.String()
}

func (s *SkillAddMetaProfilingJobResponseBody) GetData() *int64 {
	return s.Data
}

func (s *SkillAddMetaProfilingJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *SkillAddMetaProfilingJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *SkillAddMetaProfilingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SkillAddMetaProfilingJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SkillAddMetaProfilingJobResponseBody) SetData(v int64) *SkillAddMetaProfilingJobResponseBody {
	s.Data = &v
	return s
}

func (s *SkillAddMetaProfilingJobResponseBody) SetErrCode(v string) *SkillAddMetaProfilingJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SkillAddMetaProfilingJobResponseBody) SetErrMessage(v string) *SkillAddMetaProfilingJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SkillAddMetaProfilingJobResponseBody) SetRequestId(v string) *SkillAddMetaProfilingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SkillAddMetaProfilingJobResponseBody) SetSuccess(v bool) *SkillAddMetaProfilingJobResponseBody {
	s.Success = &v
	return s
}

func (s *SkillAddMetaProfilingJobResponseBody) Validate() error {
	return dara.Validate(s)
}
