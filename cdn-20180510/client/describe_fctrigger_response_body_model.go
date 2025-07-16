// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFCTriggerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFCTrigger(v *DescribeFCTriggerResponseBodyFCTrigger) *DescribeFCTriggerResponseBody
	GetFCTrigger() *DescribeFCTriggerResponseBodyFCTrigger
	SetRequestId(v string) *DescribeFCTriggerResponseBody
	GetRequestId() *string
}

type DescribeFCTriggerResponseBody struct {
	// The Function Compute trigger that you want to query.
	FCTrigger *DescribeFCTriggerResponseBodyFCTrigger `json:"FCTrigger,omitempty" xml:"FCTrigger,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// EC046C5D-8CB4-4B6B-B7F8-B335E51EF90E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFCTriggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFCTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFCTriggerResponseBody) GetFCTrigger() *DescribeFCTriggerResponseBodyFCTrigger {
	return s.FCTrigger
}

func (s *DescribeFCTriggerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFCTriggerResponseBody) SetFCTrigger(v *DescribeFCTriggerResponseBodyFCTrigger) *DescribeFCTriggerResponseBody {
	s.FCTrigger = v
	return s
}

func (s *DescribeFCTriggerResponseBody) SetRequestId(v string) *DescribeFCTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFCTriggerResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeFCTriggerResponseBodyFCTrigger struct {
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
	// The remarks of the Function Compute trigger.
	//
	// example:
	//
	// test
	Notes *string `json:"Notes,omitempty" xml:"Notes,omitempty"`
	// The assigned Resource Access Management (RAM) role.
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

func (s DescribeFCTriggerResponseBodyFCTrigger) String() string {
	return dara.Prettify(s)
}

func (s DescribeFCTriggerResponseBodyFCTrigger) GoString() string {
	return s.String()
}

func (s *DescribeFCTriggerResponseBodyFCTrigger) GetEventMetaName() *string {
	return s.EventMetaName
}

func (s *DescribeFCTriggerResponseBodyFCTrigger) GetEventMetaVersion() *string {
	return s.EventMetaVersion
}

func (s *DescribeFCTriggerResponseBodyFCTrigger) GetNotes() *string {
	return s.Notes
}

func (s *DescribeFCTriggerResponseBodyFCTrigger) GetRoleARN() *string {
	return s.RoleARN
}

func (s *DescribeFCTriggerResponseBodyFCTrigger) GetSourceArn() *string {
	return s.SourceArn
}

func (s *DescribeFCTriggerResponseBodyFCTrigger) GetTriggerARN() *string {
	return s.TriggerARN
}

func (s *DescribeFCTriggerResponseBodyFCTrigger) SetEventMetaName(v string) *DescribeFCTriggerResponseBodyFCTrigger {
	s.EventMetaName = &v
	return s
}

func (s *DescribeFCTriggerResponseBodyFCTrigger) SetEventMetaVersion(v string) *DescribeFCTriggerResponseBodyFCTrigger {
	s.EventMetaVersion = &v
	return s
}

func (s *DescribeFCTriggerResponseBodyFCTrigger) SetNotes(v string) *DescribeFCTriggerResponseBodyFCTrigger {
	s.Notes = &v
	return s
}

func (s *DescribeFCTriggerResponseBodyFCTrigger) SetRoleARN(v string) *DescribeFCTriggerResponseBodyFCTrigger {
	s.RoleARN = &v
	return s
}

func (s *DescribeFCTriggerResponseBodyFCTrigger) SetSourceArn(v string) *DescribeFCTriggerResponseBodyFCTrigger {
	s.SourceArn = &v
	return s
}

func (s *DescribeFCTriggerResponseBodyFCTrigger) SetTriggerARN(v string) *DescribeFCTriggerResponseBodyFCTrigger {
	s.TriggerARN = &v
	return s
}

func (s *DescribeFCTriggerResponseBodyFCTrigger) Validate() error {
	return dara.Validate(s)
}
