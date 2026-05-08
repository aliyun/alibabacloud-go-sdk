// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishAICoachTaskSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionId(v string) *FinishAICoachTaskSessionRequest
	GetSessionId() *string
	SetUid(v string) *FinishAICoachTaskSessionRequest
	GetUid() *string
}

type FinishAICoachTaskSessionRequest struct {
	// example:
	//
	// 111
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 222
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s FinishAICoachTaskSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s FinishAICoachTaskSessionRequest) GoString() string {
	return s.String()
}

func (s *FinishAICoachTaskSessionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *FinishAICoachTaskSessionRequest) GetUid() *string {
	return s.Uid
}

func (s *FinishAICoachTaskSessionRequest) SetSessionId(v string) *FinishAICoachTaskSessionRequest {
	s.SessionId = &v
	return s
}

func (s *FinishAICoachTaskSessionRequest) SetUid(v string) *FinishAICoachTaskSessionRequest {
	s.Uid = &v
	return s
}

func (s *FinishAICoachTaskSessionRequest) Validate() error {
	return dara.Validate(s)
}
