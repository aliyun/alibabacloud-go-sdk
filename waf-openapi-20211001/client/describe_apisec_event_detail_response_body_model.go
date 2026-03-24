// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecEventDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttackCnt(v string) *DescribeApisecEventDetailResponseBody
	GetAttackCnt() *string
	SetAttackerList(v []*string) *DescribeApisecEventDetailResponseBody
	GetAttackerList() []*string
	SetDetailValue(v string) *DescribeApisecEventDetailResponseBody
	GetDetailValue() *string
	SetEndTs(v string) *DescribeApisecEventDetailResponseBody
	GetEndTs() *string
	SetEventId(v string) *DescribeApisecEventDetailResponseBody
	GetEventId() *string
	SetEventLevel(v string) *DescribeApisecEventDetailResponseBody
	GetEventLevel() *string
	SetEventScope(v string) *DescribeApisecEventDetailResponseBody
	GetEventScope() *string
	SetEventTag(v string) *DescribeApisecEventDetailResponseBody
	GetEventTag() *string
	SetNote(v string) *DescribeApisecEventDetailResponseBody
	GetNote() *string
	SetOrigin(v string) *DescribeApisecEventDetailResponseBody
	GetOrigin() *string
	SetRequestId(v string) *DescribeApisecEventDetailResponseBody
	GetRequestId() *string
	SetStartTs(v string) *DescribeApisecEventDetailResponseBody
	GetStartTs() *string
	SetUserStatus(v string) *DescribeApisecEventDetailResponseBody
	GetUserStatus() *string
}

type DescribeApisecEventDetailResponseBody struct {
	// The total number of attacks in the security event.
	//
	// example:
	//
	// 345
	AttackCnt *string `json:"AttackCnt,omitempty" xml:"AttackCnt,omitempty"`
	// The list of attackers associated with the security event.
	AttackerList []*string `json:"AttackerList,omitempty" xml:"AttackerList,omitempty" type:"Repeated"`
	// The details of the security event.
	//
	// example:
	//
	// {\\"location\\":[\\"FR\\",\\"CN\\"],\\"location_type\\":\\"country\\"}
	DetailValue *string `json:"DetailValue,omitempty" xml:"DetailValue,omitempty"`
	// The end of the time range during which the security event occurred. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1683703260
	EndTs *string `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// The ID of the API security event.
	//
	// example:
	//
	// 18ba94fea9***e66ba0557b7b91
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The severity level of the security event. Valid values:
	//
	// - **high**: high severity.
	//
	// - **medium**: medium severity.
	//
	// - **low**: low severity.
	//
	// example:
	//
	// low
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// The dimension of the security event. Valid values:
	//
	// - **ip*	- (default): IP security event.
	//
	// - **account**: account security event.
	//
	// example:
	//
	// ip
	EventScope *string `json:"EventScope,omitempty" xml:"EventScope,omitempty"`
	// The type of the security event.
	//
	// > You can call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported event types.
	//
	// example:
	//
	// ObtainSensitiveUnauthorized
	EventTag *string `json:"EventTag,omitempty" xml:"EventTag,omitempty"`
	// The remarks added to the security event.
	//
	// example:
	//
	// already confirmed.
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The source of the event type rule. Valid values:
	//
	// - **custom**: a user-defined rule.
	//
	// - **default**: a built-in rule.
	//
	// example:
	//
	// custom
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range during which the security event occurred. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1683648000
	StartTs *string `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	// The event status. Valid values:
	//
	// - **toBeConfirmed**: to be confirmed.
	//
	// - **confirmed**: confirmed.
	//
	// - **actioned**: handled.
	//
	// - **ignored**: ignored.
	//
	// example:
	//
	// toBeConfirmed
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s DescribeApisecEventDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecEventDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisecEventDetailResponseBody) GetAttackCnt() *string {
	return s.AttackCnt
}

func (s *DescribeApisecEventDetailResponseBody) GetAttackerList() []*string {
	return s.AttackerList
}

func (s *DescribeApisecEventDetailResponseBody) GetDetailValue() *string {
	return s.DetailValue
}

func (s *DescribeApisecEventDetailResponseBody) GetEndTs() *string {
	return s.EndTs
}

func (s *DescribeApisecEventDetailResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *DescribeApisecEventDetailResponseBody) GetEventLevel() *string {
	return s.EventLevel
}

func (s *DescribeApisecEventDetailResponseBody) GetEventScope() *string {
	return s.EventScope
}

func (s *DescribeApisecEventDetailResponseBody) GetEventTag() *string {
	return s.EventTag
}

func (s *DescribeApisecEventDetailResponseBody) GetNote() *string {
	return s.Note
}

func (s *DescribeApisecEventDetailResponseBody) GetOrigin() *string {
	return s.Origin
}

func (s *DescribeApisecEventDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisecEventDetailResponseBody) GetStartTs() *string {
	return s.StartTs
}

func (s *DescribeApisecEventDetailResponseBody) GetUserStatus() *string {
	return s.UserStatus
}

func (s *DescribeApisecEventDetailResponseBody) SetAttackCnt(v string) *DescribeApisecEventDetailResponseBody {
	s.AttackCnt = &v
	return s
}

func (s *DescribeApisecEventDetailResponseBody) SetAttackerList(v []*string) *DescribeApisecEventDetailResponseBody {
	s.AttackerList = v
	return s
}

func (s *DescribeApisecEventDetailResponseBody) SetDetailValue(v string) *DescribeApisecEventDetailResponseBody {
	s.DetailValue = &v
	return s
}

func (s *DescribeApisecEventDetailResponseBody) SetEndTs(v string) *DescribeApisecEventDetailResponseBody {
	s.EndTs = &v
	return s
}

func (s *DescribeApisecEventDetailResponseBody) SetEventId(v string) *DescribeApisecEventDetailResponseBody {
	s.EventId = &v
	return s
}

func (s *DescribeApisecEventDetailResponseBody) SetEventLevel(v string) *DescribeApisecEventDetailResponseBody {
	s.EventLevel = &v
	return s
}

func (s *DescribeApisecEventDetailResponseBody) SetEventScope(v string) *DescribeApisecEventDetailResponseBody {
	s.EventScope = &v
	return s
}

func (s *DescribeApisecEventDetailResponseBody) SetEventTag(v string) *DescribeApisecEventDetailResponseBody {
	s.EventTag = &v
	return s
}

func (s *DescribeApisecEventDetailResponseBody) SetNote(v string) *DescribeApisecEventDetailResponseBody {
	s.Note = &v
	return s
}

func (s *DescribeApisecEventDetailResponseBody) SetOrigin(v string) *DescribeApisecEventDetailResponseBody {
	s.Origin = &v
	return s
}

func (s *DescribeApisecEventDetailResponseBody) SetRequestId(v string) *DescribeApisecEventDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisecEventDetailResponseBody) SetStartTs(v string) *DescribeApisecEventDetailResponseBody {
	s.StartTs = &v
	return s
}

func (s *DescribeApisecEventDetailResponseBody) SetUserStatus(v string) *DescribeApisecEventDetailResponseBody {
	s.UserStatus = &v
	return s
}

func (s *DescribeApisecEventDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
