// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAICoachTaskReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAICoachTaskReportResponseBody
	GetRequestId() *string
	SetSessionId(v string) *CreateAICoachTaskReportResponseBody
	GetSessionId() *string
}

type CreateAICoachTaskReportResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 0E8B1746-AE35-5C4B-A3A8-345B274AE32C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 79e954faffe2415ebd18188ba787d78e
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s CreateAICoachTaskReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAICoachTaskReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAICoachTaskReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAICoachTaskReportResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateAICoachTaskReportResponseBody) SetRequestId(v string) *CreateAICoachTaskReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAICoachTaskReportResponseBody) SetSessionId(v string) *CreateAICoachTaskReportResponseBody {
	s.SessionId = &v
	return s
}

func (s *CreateAICoachTaskReportResponseBody) Validate() error {
	return dara.Validate(s)
}
