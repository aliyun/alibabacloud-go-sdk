// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachTaskSessionReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionId(v string) *GetAICoachTaskSessionReportRequest
	GetSessionId() *string
	SetUid(v string) *GetAICoachTaskSessionReportRequest
	GetUid() *string
}

type GetAICoachTaskSessionReportRequest struct {
	// example:
	//
	// 1111
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 1707732338016307
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetAICoachTaskSessionReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachTaskSessionReportRequest) GoString() string {
	return s.String()
}

func (s *GetAICoachTaskSessionReportRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetAICoachTaskSessionReportRequest) GetUid() *string {
	return s.Uid
}

func (s *GetAICoachTaskSessionReportRequest) SetSessionId(v string) *GetAICoachTaskSessionReportRequest {
	s.SessionId = &v
	return s
}

func (s *GetAICoachTaskSessionReportRequest) SetUid(v string) *GetAICoachTaskSessionReportRequest {
	s.Uid = &v
	return s
}

func (s *GetAICoachTaskSessionReportRequest) Validate() error {
	return dara.Validate(s)
}
