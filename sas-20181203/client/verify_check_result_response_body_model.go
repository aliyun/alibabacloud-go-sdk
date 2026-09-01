// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCheckResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *VerifyCheckResultResponseBodyData) *VerifyCheckResultResponseBody
	GetData() *VerifyCheckResultResponseBodyData
	SetRequestId(v string) *VerifyCheckResultResponseBody
	GetRequestId() *string
}

type VerifyCheckResultResponseBody struct {
	// The returned data.
	Data *VerifyCheckResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. Alibaba Cloud generates a unique ID for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 7C0A3FA0-AA32-5660-8989-85A5582F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyCheckResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyCheckResultResponseBody) GetData() *VerifyCheckResultResponseBodyData {
	return s.Data
}

func (s *VerifyCheckResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyCheckResultResponseBody) SetData(v *VerifyCheckResultResponseBodyData) *VerifyCheckResultResponseBody {
	s.Data = v
	return s
}

func (s *VerifyCheckResultResponseBody) SetRequestId(v string) *VerifyCheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyCheckResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyCheckResultResponseBodyData struct {
	// The operation code of the cloud service configuration check task. Valid values:
	//
	// - **Throttling**: Rate limited.
	//
	// - **ActionTrialUnauthorized**: Unauthorized error.
	//
	// example:
	//
	// Throttling
	OperateCode *string `json:"OperateCode,omitempty" xml:"OperateCode,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 8eded533-5348-468c-aa1d-0aa2934a7***
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The throttling duration. Unit: seconds.
	//
	// example:
	//
	// 1800
	ThrottlingTimeSecond *int32 `json:"ThrottlingTimeSecond,omitempty" xml:"ThrottlingTimeSecond,omitempty"`
}

func (s VerifyCheckResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VerifyCheckResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifyCheckResultResponseBodyData) GetOperateCode() *string {
	return s.OperateCode
}

func (s *VerifyCheckResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *VerifyCheckResultResponseBodyData) GetThrottlingTimeSecond() *int32 {
	return s.ThrottlingTimeSecond
}

func (s *VerifyCheckResultResponseBodyData) SetOperateCode(v string) *VerifyCheckResultResponseBodyData {
	s.OperateCode = &v
	return s
}

func (s *VerifyCheckResultResponseBodyData) SetTaskId(v string) *VerifyCheckResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *VerifyCheckResultResponseBodyData) SetThrottlingTimeSecond(v int32) *VerifyCheckResultResponseBodyData {
	s.ThrottlingTimeSecond = &v
	return s
}

func (s *VerifyCheckResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
