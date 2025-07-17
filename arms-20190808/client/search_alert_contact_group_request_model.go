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
	// The ID of the alert contact group. You can query multiple alert contact groups at a time. Separate multiple group IDs with commas (,).
	//
	// example:
	//
	// 746
	ContactGroupIds *string `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	// The name of the alert contact group.
	//
	// example:
	//
	// TestGroup
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	// The ID of the alert contact. You can call the SearchAlertContact operation to query the contact IDs. For more information, see [SearchAlertContact](https://help.aliyun.com/document_detail/130703.html).
	//
	// example:
	//
	// 123
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The name of the alert contact.
	//
	// example:
	//
	// John Doe
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// Specifies whether to return all alert contacts in the queried alert contact group. By default, not all alert contacts are returned.
	//
	// example:
	//
	// true
	IsDetail *bool `json:"IsDetail,omitempty" xml:"IsDetail,omitempty"`
	// The ID of the region. Default value: `cn-hangzhou`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
