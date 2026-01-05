// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSlotsResponseBody
	GetRequestId() *string
	SetSlotIds(v string) *CreateSlotsResponseBody
	GetSlotIds() *string
	SetSummary(v map[string]*string) *CreateSlotsResponseBody
	GetSummary() map[string]*string
}

type CreateSlotsResponseBody struct {
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// slot-5zk866779me51jgu3w
	SlotIds *string            `json:"SlotIds,omitempty" xml:"SlotIds,omitempty"`
	Summary map[string]*string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s CreateSlotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSlotsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSlotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSlotsResponseBody) GetSlotIds() *string {
	return s.SlotIds
}

func (s *CreateSlotsResponseBody) GetSummary() map[string]*string {
	return s.Summary
}

func (s *CreateSlotsResponseBody) SetRequestId(v string) *CreateSlotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSlotsResponseBody) SetSlotIds(v string) *CreateSlotsResponseBody {
	s.SlotIds = &v
	return s
}

func (s *CreateSlotsResponseBody) SetSummary(v map[string]*string) *CreateSlotsResponseBody {
	s.Summary = v
	return s
}

func (s *CreateSlotsResponseBody) Validate() error {
	return dara.Validate(s)
}
