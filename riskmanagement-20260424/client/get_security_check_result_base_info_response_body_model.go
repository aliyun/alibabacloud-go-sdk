// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityCheckResultBaseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSecurityCheckResultBaseInfoResponseBody
	GetCode() *string
	SetData(v *GetSecurityCheckResultBaseInfoResponseBodyData) *GetSecurityCheckResultBaseInfoResponseBody
	GetData() *GetSecurityCheckResultBaseInfoResponseBodyData
	SetMessage(v string) *GetSecurityCheckResultBaseInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSecurityCheckResultBaseInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSecurityCheckResultBaseInfoResponseBody
	GetSuccess() *bool
}

type GetSecurityCheckResultBaseInfoResponseBody struct {
	// example:
	//
	// 200
	Code *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSecurityCheckResultBaseInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful‌
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 739705BB-B0EF-554B-B3A8-383F4F93E067
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSecurityCheckResultBaseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityCheckResultBaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecurityCheckResultBaseInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSecurityCheckResultBaseInfoResponseBody) GetData() *GetSecurityCheckResultBaseInfoResponseBodyData {
	return s.Data
}

func (s *GetSecurityCheckResultBaseInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSecurityCheckResultBaseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecurityCheckResultBaseInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSecurityCheckResultBaseInfoResponseBody) SetCode(v string) *GetSecurityCheckResultBaseInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetSecurityCheckResultBaseInfoResponseBody) SetData(v *GetSecurityCheckResultBaseInfoResponseBodyData) *GetSecurityCheckResultBaseInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetSecurityCheckResultBaseInfoResponseBody) SetMessage(v string) *GetSecurityCheckResultBaseInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetSecurityCheckResultBaseInfoResponseBody) SetRequestId(v string) *GetSecurityCheckResultBaseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecurityCheckResultBaseInfoResponseBody) SetSuccess(v bool) *GetSecurityCheckResultBaseInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetSecurityCheckResultBaseInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecurityCheckResultBaseInfoResponseBodyData struct {
	// example:
	//
	// true
	ConfigCompleted *string `json:"ConfigCompleted,omitempty" xml:"ConfigCompleted,omitempty"`
	// example:
	//
	// 1
	PendingSecurityAlertCount *int32 `json:"PendingSecurityAlertCount,omitempty" xml:"PendingSecurityAlertCount,omitempty"`
	// example:
	//
	// 5
	PendingVulnerabilityCount *int32 `json:"PendingVulnerabilityCount,omitempty" xml:"PendingVulnerabilityCount,omitempty"`
	// example:
	//
	// 30%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s GetSecurityCheckResultBaseInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityCheckResultBaseInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSecurityCheckResultBaseInfoResponseBodyData) GetConfigCompleted() *string {
	return s.ConfigCompleted
}

func (s *GetSecurityCheckResultBaseInfoResponseBodyData) GetPendingSecurityAlertCount() *int32 {
	return s.PendingSecurityAlertCount
}

func (s *GetSecurityCheckResultBaseInfoResponseBodyData) GetPendingVulnerabilityCount() *int32 {
	return s.PendingVulnerabilityCount
}

func (s *GetSecurityCheckResultBaseInfoResponseBodyData) GetProgress() *string {
	return s.Progress
}

func (s *GetSecurityCheckResultBaseInfoResponseBodyData) SetConfigCompleted(v string) *GetSecurityCheckResultBaseInfoResponseBodyData {
	s.ConfigCompleted = &v
	return s
}

func (s *GetSecurityCheckResultBaseInfoResponseBodyData) SetPendingSecurityAlertCount(v int32) *GetSecurityCheckResultBaseInfoResponseBodyData {
	s.PendingSecurityAlertCount = &v
	return s
}

func (s *GetSecurityCheckResultBaseInfoResponseBodyData) SetPendingVulnerabilityCount(v int32) *GetSecurityCheckResultBaseInfoResponseBodyData {
	s.PendingVulnerabilityCount = &v
	return s
}

func (s *GetSecurityCheckResultBaseInfoResponseBodyData) SetProgress(v string) *GetSecurityCheckResultBaseInfoResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetSecurityCheckResultBaseInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
