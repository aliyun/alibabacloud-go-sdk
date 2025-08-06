// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisplayAIAuditSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsDisplay(v bool) *DisplayAIAuditSwitchResponseBody
	GetIsDisplay() *bool
	SetRequestId(v string) *DisplayAIAuditSwitchResponseBody
	GetRequestId() *string
}

type DisplayAIAuditSwitchResponseBody struct {
	IsDisplay *bool   `json:"IsDisplay,omitempty" xml:"IsDisplay,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisplayAIAuditSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisplayAIAuditSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *DisplayAIAuditSwitchResponseBody) GetIsDisplay() *bool {
	return s.IsDisplay
}

func (s *DisplayAIAuditSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisplayAIAuditSwitchResponseBody) SetIsDisplay(v bool) *DisplayAIAuditSwitchResponseBody {
	s.IsDisplay = &v
	return s
}

func (s *DisplayAIAuditSwitchResponseBody) SetRequestId(v string) *DisplayAIAuditSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisplayAIAuditSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
