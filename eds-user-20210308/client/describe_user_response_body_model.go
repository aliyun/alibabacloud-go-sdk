// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUserResponseBody
	GetRequestId() *string
	SetUser(v *DescribeUserResponseBodyUser) *DescribeUserResponseBody
	GetUser() *DescribeUserResponseBodyUser
}

type DescribeUserResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	User      *DescribeUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s DescribeUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserResponseBody) GetUser() *DescribeUserResponseBodyUser {
	return s.User
}

func (s *DescribeUserResponseBody) SetRequestId(v string) *DescribeUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserResponseBody) SetUser(v *DescribeUserResponseBodyUser) *DescribeUserResponseBody {
	s.User = v
	return s
}

func (s *DescribeUserResponseBody) Validate() error {
	if s.User != nil {
		if err := s.User.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUserResponseBodyUser struct {
	// example:
	//
	// alex@test-email.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// alex
	EndUserId    *string                                   `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ExternalInfo *DescribeUserResponseBodyUserExternalInfo `json:"ExternalInfo,omitempty" xml:"ExternalInfo,omitempty" type:"Struct"`
	Extras       map[string]*string                        `json:"Extras,omitempty" xml:"Extras,omitempty"`
	// example:
	//
	// 2025-01-01 12:00:00
	GmtCreate *int64    `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	NickName  *string   `json:"NickName,omitempty" xml:"NickName,omitempty"`
	OrgIds    []*string `json:"OrgIds,omitempty" xml:"OrgIds,omitempty" type:"Repeated"`
	OrgPaths  []*string `json:"OrgPaths,omitempty" xml:"OrgPaths,omitempty" type:"Repeated"`
	// example:
	//
	// 1888888****
	Phone      *string                                   `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Properties []*DescribeUserResponseBodyUserProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
	// example:
	//
	// remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// aisdfumj****
	WyId *string `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s DescribeUserResponseBodyUser) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *DescribeUserResponseBodyUser) GetEmail() *string {
	return s.Email
}

func (s *DescribeUserResponseBodyUser) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeUserResponseBodyUser) GetExternalInfo() *DescribeUserResponseBodyUserExternalInfo {
	return s.ExternalInfo
}

func (s *DescribeUserResponseBodyUser) GetExtras() map[string]*string {
	return s.Extras
}

func (s *DescribeUserResponseBodyUser) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeUserResponseBodyUser) GetNickName() *string {
	return s.NickName
}

func (s *DescribeUserResponseBodyUser) GetOrgIds() []*string {
	return s.OrgIds
}

func (s *DescribeUserResponseBodyUser) GetOrgPaths() []*string {
	return s.OrgPaths
}

func (s *DescribeUserResponseBodyUser) GetPhone() *string {
	return s.Phone
}

func (s *DescribeUserResponseBodyUser) GetProperties() []*DescribeUserResponseBodyUserProperties {
	return s.Properties
}

func (s *DescribeUserResponseBodyUser) GetRemark() *string {
	return s.Remark
}

func (s *DescribeUserResponseBodyUser) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeUserResponseBodyUser) GetWyId() *string {
	return s.WyId
}

func (s *DescribeUserResponseBodyUser) SetEmail(v string) *DescribeUserResponseBodyUser {
	s.Email = &v
	return s
}

func (s *DescribeUserResponseBodyUser) SetEndUserId(v string) *DescribeUserResponseBodyUser {
	s.EndUserId = &v
	return s
}

func (s *DescribeUserResponseBodyUser) SetExternalInfo(v *DescribeUserResponseBodyUserExternalInfo) *DescribeUserResponseBodyUser {
	s.ExternalInfo = v
	return s
}

func (s *DescribeUserResponseBodyUser) SetExtras(v map[string]*string) *DescribeUserResponseBodyUser {
	s.Extras = v
	return s
}

func (s *DescribeUserResponseBodyUser) SetGmtCreate(v int64) *DescribeUserResponseBodyUser {
	s.GmtCreate = &v
	return s
}

func (s *DescribeUserResponseBodyUser) SetNickName(v string) *DescribeUserResponseBodyUser {
	s.NickName = &v
	return s
}

func (s *DescribeUserResponseBodyUser) SetOrgIds(v []*string) *DescribeUserResponseBodyUser {
	s.OrgIds = v
	return s
}

func (s *DescribeUserResponseBodyUser) SetOrgPaths(v []*string) *DescribeUserResponseBodyUser {
	s.OrgPaths = v
	return s
}

func (s *DescribeUserResponseBodyUser) SetPhone(v string) *DescribeUserResponseBodyUser {
	s.Phone = &v
	return s
}

func (s *DescribeUserResponseBodyUser) SetProperties(v []*DescribeUserResponseBodyUserProperties) *DescribeUserResponseBodyUser {
	s.Properties = v
	return s
}

func (s *DescribeUserResponseBodyUser) SetRemark(v string) *DescribeUserResponseBodyUser {
	s.Remark = &v
	return s
}

func (s *DescribeUserResponseBodyUser) SetStatus(v int32) *DescribeUserResponseBodyUser {
	s.Status = &v
	return s
}

func (s *DescribeUserResponseBodyUser) SetWyId(v string) *DescribeUserResponseBodyUser {
	s.WyId = &v
	return s
}

func (s *DescribeUserResponseBodyUser) Validate() error {
	if s.ExternalInfo != nil {
		if err := s.ExternalInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Properties != nil {
		for _, item := range s.Properties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserResponseBodyUserExternalInfo struct {
	// example:
	//
	// oijjnabsf****
	ExternalId *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	// example:
	//
	// Alex
	ExternalName *string `json:"ExternalName,omitempty" xml:"ExternalName,omitempty"`
	// example:
	//
	// 15412***
	JobNumber *string `json:"JobNumber,omitempty" xml:"JobNumber,omitempty"`
	// example:
	//
	// OIDC
	SsoType *string `json:"SsoType,omitempty" xml:"SsoType,omitempty"`
}

func (s DescribeUserResponseBodyUserExternalInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResponseBodyUserExternalInfo) GoString() string {
	return s.String()
}

func (s *DescribeUserResponseBodyUserExternalInfo) GetExternalId() *string {
	return s.ExternalId
}

func (s *DescribeUserResponseBodyUserExternalInfo) GetExternalName() *string {
	return s.ExternalName
}

func (s *DescribeUserResponseBodyUserExternalInfo) GetJobNumber() *string {
	return s.JobNumber
}

func (s *DescribeUserResponseBodyUserExternalInfo) GetSsoType() *string {
	return s.SsoType
}

func (s *DescribeUserResponseBodyUserExternalInfo) SetExternalId(v string) *DescribeUserResponseBodyUserExternalInfo {
	s.ExternalId = &v
	return s
}

func (s *DescribeUserResponseBodyUserExternalInfo) SetExternalName(v string) *DescribeUserResponseBodyUserExternalInfo {
	s.ExternalName = &v
	return s
}

func (s *DescribeUserResponseBodyUserExternalInfo) SetJobNumber(v string) *DescribeUserResponseBodyUserExternalInfo {
	s.JobNumber = &v
	return s
}

func (s *DescribeUserResponseBodyUserExternalInfo) SetSsoType(v string) *DescribeUserResponseBodyUserExternalInfo {
	s.SsoType = &v
	return s
}

func (s *DescribeUserResponseBodyUserExternalInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeUserResponseBodyUserProperties struct {
	// example:
	//
	// role
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// teacher
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeUserResponseBodyUserProperties) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResponseBodyUserProperties) GoString() string {
	return s.String()
}

func (s *DescribeUserResponseBodyUserProperties) GetKey() *string {
	return s.Key
}

func (s *DescribeUserResponseBodyUserProperties) GetValue() *string {
	return s.Value
}

func (s *DescribeUserResponseBodyUserProperties) SetKey(v string) *DescribeUserResponseBodyUserProperties {
	s.Key = &v
	return s
}

func (s *DescribeUserResponseBodyUserProperties) SetValue(v string) *DescribeUserResponseBodyUserProperties {
	s.Value = &v
	return s
}

func (s *DescribeUserResponseBodyUserProperties) Validate() error {
	return dara.Validate(s)
}
