// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyFirewallTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApplyFirewallTemplateResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ApplyFirewallTemplateResponseBody
	GetTaskId() *string
}

type ApplyFirewallTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the template application task.
	//
	// example:
	//
	// aft-ikgt0bynitvs3****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ApplyFirewallTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyFirewallTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyFirewallTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyFirewallTemplateResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ApplyFirewallTemplateResponseBody) SetRequestId(v string) *ApplyFirewallTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyFirewallTemplateResponseBody) SetTaskId(v string) *ApplyFirewallTemplateResponseBody {
	s.TaskId = &v
	return s
}

func (s *ApplyFirewallTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
