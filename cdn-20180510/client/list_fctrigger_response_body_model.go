// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFCTriggerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFCTriggers(v []*ListFCTriggerResponseBodyFCTriggers) *ListFCTriggerResponseBody
	GetFCTriggers() []*ListFCTriggerResponseBodyFCTriggers
	SetRequestId(v string) *ListFCTriggerResponseBody
	GetRequestId() *string
}

type ListFCTriggerResponseBody struct {
	// The Function Compute triggers that are set for Alibaba Cloud CDN events.
	FCTriggers []*ListFCTriggerResponseBodyFCTriggers `json:"FCTriggers,omitempty" xml:"FCTriggers,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// EC046C5D-8CB4-4B6B-B7F8-B335E51EF90E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFCTriggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFCTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *ListFCTriggerResponseBody) GetFCTriggers() []*ListFCTriggerResponseBodyFCTriggers {
	return s.FCTriggers
}

func (s *ListFCTriggerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFCTriggerResponseBody) SetFCTriggers(v []*ListFCTriggerResponseBodyFCTriggers) *ListFCTriggerResponseBody {
	s.FCTriggers = v
	return s
}

func (s *ListFCTriggerResponseBody) SetRequestId(v string) *ListFCTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFCTriggerResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFCTriggerResponseBodyFCTriggers struct {
	// The name of the event.
	//
	// example:
	//
	// LogFileCreated
	EventMetaName *string `json:"EventMetaName,omitempty" xml:"EventMetaName,omitempty"`
	// The version of the event.
	//
	// example:
	//
	// 1.0.0
	EventMetaVersion *string `json:"EventMetaVersion,omitempty" xml:"EventMetaVersion,omitempty"`
	// The remarks.
	//
	// example:
	//
	// Test
	Notes *string `json:"Notes,omitempty" xml:"Notes,omitempty"`
	// The Resource Access Management (RAM) role.
	//
	// example:
	//
	// acs:ram:: 1234567890:role/aliyuncdneventnotificationrole
	RoleARN *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	// The resources and filters for event listening.
	//
	// example:
	//
	// acs:cdn:*:1234567890:domain/example.com
	SourceArn *string `json:"SourceArn,omitempty" xml:"SourceArn,omitempty"`
	// The trigger that corresponds to the Function Compute service.
	//
	// example:
	//
	// acs:fc:cn-beijing: 1234567890:services/FCTestService/functions/printEvent/triggers/testtrigger
	TriggerARN *string `json:"TriggerARN,omitempty" xml:"TriggerARN,omitempty"`
}

func (s ListFCTriggerResponseBodyFCTriggers) String() string {
	return dara.Prettify(s)
}

func (s ListFCTriggerResponseBodyFCTriggers) GoString() string {
	return s.String()
}

func (s *ListFCTriggerResponseBodyFCTriggers) GetEventMetaName() *string {
	return s.EventMetaName
}

func (s *ListFCTriggerResponseBodyFCTriggers) GetEventMetaVersion() *string {
	return s.EventMetaVersion
}

func (s *ListFCTriggerResponseBodyFCTriggers) GetNotes() *string {
	return s.Notes
}

func (s *ListFCTriggerResponseBodyFCTriggers) GetRoleARN() *string {
	return s.RoleARN
}

func (s *ListFCTriggerResponseBodyFCTriggers) GetSourceArn() *string {
	return s.SourceArn
}

func (s *ListFCTriggerResponseBodyFCTriggers) GetTriggerARN() *string {
	return s.TriggerARN
}

func (s *ListFCTriggerResponseBodyFCTriggers) SetEventMetaName(v string) *ListFCTriggerResponseBodyFCTriggers {
	s.EventMetaName = &v
	return s
}

func (s *ListFCTriggerResponseBodyFCTriggers) SetEventMetaVersion(v string) *ListFCTriggerResponseBodyFCTriggers {
	s.EventMetaVersion = &v
	return s
}

func (s *ListFCTriggerResponseBodyFCTriggers) SetNotes(v string) *ListFCTriggerResponseBodyFCTriggers {
	s.Notes = &v
	return s
}

func (s *ListFCTriggerResponseBodyFCTriggers) SetRoleARN(v string) *ListFCTriggerResponseBodyFCTriggers {
	s.RoleARN = &v
	return s
}

func (s *ListFCTriggerResponseBodyFCTriggers) SetSourceArn(v string) *ListFCTriggerResponseBodyFCTriggers {
	s.SourceArn = &v
	return s
}

func (s *ListFCTriggerResponseBodyFCTriggers) SetTriggerARN(v string) *ListFCTriggerResponseBodyFCTriggers {
	s.TriggerARN = &v
	return s
}

func (s *ListFCTriggerResponseBodyFCTriggers) Validate() error {
	return dara.Validate(s)
}
