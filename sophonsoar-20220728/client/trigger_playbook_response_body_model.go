// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerPlaybookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TriggerPlaybookResponseBody
	GetRequestId() *string
	SetTriggerUuid(v string) *TriggerPlaybookResponseBody
	GetTriggerUuid() *string
}

type TriggerPlaybookResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD5A8DB6-A42C-532B-BCE8-83E69550CD59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The running UUID of the playbook. This parameter is used to query the running result of the playbook.
	//
	// example:
	//
	// 55E63C57-D6C8-5036-A770-5CB10AC807AA
	TriggerUuid *string `json:"TriggerUuid,omitempty" xml:"TriggerUuid,omitempty"`
}

func (s TriggerPlaybookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TriggerPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerPlaybookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TriggerPlaybookResponseBody) GetTriggerUuid() *string {
	return s.TriggerUuid
}

func (s *TriggerPlaybookResponseBody) SetRequestId(v string) *TriggerPlaybookResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerPlaybookResponseBody) SetTriggerUuid(v string) *TriggerPlaybookResponseBody {
	s.TriggerUuid = &v
	return s
}

func (s *TriggerPlaybookResponseBody) Validate() error {
	return dara.Validate(s)
}
