// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAlertContactGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupIds(v string) *SearchAlertContactGroupRequest
	GetContactGroupIds() *string
	SetContactGroupName(v string) *SearchAlertContactGroupRequest
	GetContactGroupName() *string
	SetContactId(v int64) *SearchAlertContactGroupRequest
	GetContactId() *int64
	SetContactName(v string) *SearchAlertContactGroupRequest
	GetContactName() *string
	SetIsDetail(v bool) *SearchAlertContactGroupRequest
	GetIsDetail() *bool
	SetRegionId(v string) *SearchAlertContactGroupRequest
	GetRegionId() *string
}

type SearchAlertContactGroupRequest struct {
	ContactGroupIds  *string `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	ContactId        *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName      *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	IsDetail         *bool   `json:"IsDetail,omitempty" xml:"IsDetail,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SearchAlertContactGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertContactGroupRequest) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupRequest) GetContactGroupIds() *string {
	return s.ContactGroupIds
}

func (s *SearchAlertContactGroupRequest) GetContactGroupName() *string {
	return s.ContactGroupName
}

func (s *SearchAlertContactGroupRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *SearchAlertContactGroupRequest) GetContactName() *string {
	return s.ContactName
}

func (s *SearchAlertContactGroupRequest) GetIsDetail() *bool {
	return s.IsDetail
}

func (s *SearchAlertContactGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchAlertContactGroupRequest) SetContactGroupIds(v string) *SearchAlertContactGroupRequest {
	s.ContactGroupIds = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetContactGroupName(v string) *SearchAlertContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetContactId(v int64) *SearchAlertContactGroupRequest {
	s.ContactId = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetContactName(v string) *SearchAlertContactGroupRequest {
	s.ContactName = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetIsDetail(v bool) *SearchAlertContactGroupRequest {
	s.IsDetail = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetRegionId(v string) *SearchAlertContactGroupRequest {
	s.RegionId = &v
	return s
}

func (s *SearchAlertContactGroupRequest) Validate() error {
	return dara.Validate(s)
}
