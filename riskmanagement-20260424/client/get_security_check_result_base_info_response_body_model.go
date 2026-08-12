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
	// The status code.
	//
	// >  200: The request was successful. Other codes (such as 500 or 400): An error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data list.
	Data *GetSecurityCheckResultBaseInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The prompt message.
	//
	// example:
	//
	// successful‌
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 739705BB-B0EF-554B-B3A8-383F4F93E067
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the current API call itself was successful. This does not indicate the success of subsequent business operations.
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
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
	// The configuration item check result.
	//
	// example:
	//
	// true
	ConfigCompleted *string `json:"ConfigCompleted,omitempty" xml:"ConfigCompleted,omitempty"`
	// The number of pending security alerts.
	//
	// example:
	//
	// 1
	PendingSecurityAlertCount *int32 `json:"PendingSecurityAlertCount,omitempty" xml:"PendingSecurityAlertCount,omitempty"`
	// The number of pending vulnerabilities.
	//
	// example:
	//
	// 5
	PendingVulnerabilityCount *int32 `json:"PendingVulnerabilityCount,omitempty" xml:"PendingVulnerabilityCount,omitempty"`
	// The percentage of the health check task progress.
	//
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
