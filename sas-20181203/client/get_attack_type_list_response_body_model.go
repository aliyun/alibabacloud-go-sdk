// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackTypeListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttackTypeList(v []*GetAttackTypeListResponseBodyAttackTypeList) *GetAttackTypeListResponseBody
	GetAttackTypeList() []*GetAttackTypeListResponseBodyAttackTypeList
	SetRequestId(v string) *GetAttackTypeListResponseBody
	GetRequestId() *string
}

type GetAttackTypeListResponseBody struct {
	// The attack types.
	AttackTypeList []*GetAttackTypeListResponseBodyAttackTypeList `json:"AttackTypeList,omitempty" xml:"AttackTypeList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// CE500770-42D3-442E-9DDD-156E0F9F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAttackTypeListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAttackTypeListResponseBody) GoString() string {
	return s.String()
}

func (s *GetAttackTypeListResponseBody) GetAttackTypeList() []*GetAttackTypeListResponseBodyAttackTypeList {
	return s.AttackTypeList
}

func (s *GetAttackTypeListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAttackTypeListResponseBody) SetAttackTypeList(v []*GetAttackTypeListResponseBodyAttackTypeList) *GetAttackTypeListResponseBody {
	s.AttackTypeList = v
	return s
}

func (s *GetAttackTypeListResponseBody) SetRequestId(v string) *GetAttackTypeListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAttackTypeListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAttackTypeListResponseBodyAttackTypeList struct {
	// The description of the attack type.
	//
	// example:
	//
	// sas.attack.type.type12
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The attack source. Valid values:
	//
	// 	- **cfw**: Cloud Firewall
	//
	// 	- **alinet**: network defense plug-in
	//
	// 	- **waf**: Web Application Firewall (WAF)
	//
	// example:
	//
	// alinet
	StatusType *string `json:"Status_Type,omitempty" xml:"Status_Type,omitempty"`
	// The value of the attack type.
	//
	// example:
	//
	// upload
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAttackTypeListResponseBodyAttackTypeList) String() string {
	return dara.Prettify(s)
}

func (s GetAttackTypeListResponseBodyAttackTypeList) GoString() string {
	return s.String()
}

func (s *GetAttackTypeListResponseBodyAttackTypeList) GetLabel() *string {
	return s.Label
}

func (s *GetAttackTypeListResponseBodyAttackTypeList) GetStatusType() *string {
	return s.StatusType
}

func (s *GetAttackTypeListResponseBodyAttackTypeList) GetValue() *string {
	return s.Value
}

func (s *GetAttackTypeListResponseBodyAttackTypeList) SetLabel(v string) *GetAttackTypeListResponseBodyAttackTypeList {
	s.Label = &v
	return s
}

func (s *GetAttackTypeListResponseBodyAttackTypeList) SetStatusType(v string) *GetAttackTypeListResponseBodyAttackTypeList {
	s.StatusType = &v
	return s
}

func (s *GetAttackTypeListResponseBodyAttackTypeList) SetValue(v string) *GetAttackTypeListResponseBodyAttackTypeList {
	s.Value = &v
	return s
}

func (s *GetAttackTypeListResponseBodyAttackTypeList) Validate() error {
	return dara.Validate(s)
}
