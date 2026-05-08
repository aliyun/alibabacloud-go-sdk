// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseAICoachTaskSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionId(v string) *CloseAICoachTaskSessionRequest
	GetSessionId() *string
	SetUid(v string) *CloseAICoachTaskSessionRequest
	GetUid() *string
}

type CloseAICoachTaskSessionRequest struct {
	// example:
	//
	// 11
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 273610276967782972
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CloseAICoachTaskSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseAICoachTaskSessionRequest) GoString() string {
	return s.String()
}

func (s *CloseAICoachTaskSessionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *CloseAICoachTaskSessionRequest) GetUid() *string {
	return s.Uid
}

func (s *CloseAICoachTaskSessionRequest) SetSessionId(v string) *CloseAICoachTaskSessionRequest {
	s.SessionId = &v
	return s
}

func (s *CloseAICoachTaskSessionRequest) SetUid(v string) *CloseAICoachTaskSessionRequest {
	s.Uid = &v
	return s
}

func (s *CloseAICoachTaskSessionRequest) Validate() error {
	return dara.Validate(s)
}
