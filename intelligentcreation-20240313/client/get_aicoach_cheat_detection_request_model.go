// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachCheatDetectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionId(v string) *GetAICoachCheatDetectionRequest
	GetSessionId() *string
}

type GetAICoachCheatDetectionRequest struct {
	// example:
	//
	// 79e954faffe2415ebd18188ba787d78e
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s GetAICoachCheatDetectionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachCheatDetectionRequest) GoString() string {
	return s.String()
}

func (s *GetAICoachCheatDetectionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetAICoachCheatDetectionRequest) SetSessionId(v string) *GetAICoachCheatDetectionRequest {
	s.SessionId = &v
	return s
}

func (s *GetAICoachCheatDetectionRequest) Validate() error {
	return dara.Validate(s)
}
