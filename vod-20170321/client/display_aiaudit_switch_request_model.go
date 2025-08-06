// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisplayAIAuditSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DisplayAIAuditSwitchRequest
	GetOwnerId() *int64
}

type DisplayAIAuditSwitchRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DisplayAIAuditSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s DisplayAIAuditSwitchRequest) GoString() string {
	return s.String()
}

func (s *DisplayAIAuditSwitchRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisplayAIAuditSwitchRequest) SetOwnerId(v int64) *DisplayAIAuditSwitchRequest {
	s.OwnerId = &v
	return s
}

func (s *DisplayAIAuditSwitchRequest) Validate() error {
	return dara.Validate(s)
}
