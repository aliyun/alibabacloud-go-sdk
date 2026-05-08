// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachTaskSessionResourceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionId(v string) *GetAICoachTaskSessionResourceUsageRequest
	GetSessionId() *string
}

type GetAICoachTaskSessionResourceUsageRequest struct {
	// example:
	//
	// ccce827bcd2c4b13a0fd9ea8657a48cc
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s GetAICoachTaskSessionResourceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachTaskSessionResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *GetAICoachTaskSessionResourceUsageRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetAICoachTaskSessionResourceUsageRequest) SetSessionId(v string) *GetAICoachTaskSessionResourceUsageRequest {
	s.SessionId = &v
	return s
}

func (s *GetAICoachTaskSessionResourceUsageRequest) Validate() error {
	return dara.Validate(s)
}
