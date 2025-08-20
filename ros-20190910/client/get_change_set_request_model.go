// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChangeSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeSetId(v string) *GetChangeSetRequest
	GetChangeSetId() *string
	SetRegionId(v string) *GetChangeSetRequest
	GetRegionId() *string
	SetShowTemplate(v bool) *GetChangeSetRequest
	GetShowTemplate() *bool
}

type GetChangeSetRequest struct {
	// The ID of the change set.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4c11658d-bd47-4dd0-ba64-727edc62****
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	// The region ID of the change set. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to obtain the template. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	ShowTemplate *bool `json:"ShowTemplate,omitempty" xml:"ShowTemplate,omitempty"`
}

func (s GetChangeSetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChangeSetRequest) GoString() string {
	return s.String()
}

func (s *GetChangeSetRequest) GetChangeSetId() *string {
	return s.ChangeSetId
}

func (s *GetChangeSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetChangeSetRequest) GetShowTemplate() *bool {
	return s.ShowTemplate
}

func (s *GetChangeSetRequest) SetChangeSetId(v string) *GetChangeSetRequest {
	s.ChangeSetId = &v
	return s
}

func (s *GetChangeSetRequest) SetRegionId(v string) *GetChangeSetRequest {
	s.RegionId = &v
	return s
}

func (s *GetChangeSetRequest) SetShowTemplate(v bool) *GetChangeSetRequest {
	s.ShowTemplate = &v
	return s
}

func (s *GetChangeSetRequest) Validate() error {
	return dara.Validate(s)
}
