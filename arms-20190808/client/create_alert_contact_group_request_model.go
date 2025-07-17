// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertContactGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupName(v string) *CreateAlertContactGroupRequest
	GetContactGroupName() *string
	SetContactIds(v string) *CreateAlertContactGroupRequest
	GetContactIds() *string
	SetRegionId(v string) *CreateAlertContactGroupRequest
	GetRegionId() *string
}

type CreateAlertContactGroupRequest struct {
	// The name of the alert contact group.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestGroup
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	// The IDs of contacts in the contact group. Separate multiple contact IDs with spaces. You can call the SearchAlertContact operation to query the contact IDs. For more information, see [SearchAlertContact](https://help.aliyun.com/document_detail/130703.html).
	//
	// example:
	//
	// 12	- 23	- 34*
	ContactIds *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	// The region ID. Default value: `cn-hangzhou`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateAlertContactGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertContactGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertContactGroupRequest) GetContactGroupName() *string {
	return s.ContactGroupName
}

func (s *CreateAlertContactGroupRequest) GetContactIds() *string {
	return s.ContactIds
}

func (s *CreateAlertContactGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAlertContactGroupRequest) SetContactGroupName(v string) *CreateAlertContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

func (s *CreateAlertContactGroupRequest) SetContactIds(v string) *CreateAlertContactGroupRequest {
	s.ContactIds = &v
	return s
}

func (s *CreateAlertContactGroupRequest) SetRegionId(v string) *CreateAlertContactGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAlertContactGroupRequest) Validate() error {
	return dara.Validate(s)
}
