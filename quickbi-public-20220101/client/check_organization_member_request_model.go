// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckOrganizationMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *CheckOrganizationMemberRequest
	GetUserId() *string
}

type CheckOrganizationMemberRequest struct {
	// The ID of the Quick BI user. This is not your Alibaba Cloud account ID. Call the [QueryUserInfoByAccount](https://next.api.aliyun.com/api/quickbi-public/2022-01-01/QueryUserInfoByAccount?spm=api-workbench.api_explorer.0.0.672f50daGq9ooV\\&params=%7B%7D\\&tab=DOC\\&sdkStyle=old\\&RegionId=cn-hangzhou) operation to obtain the user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// adfssd-sdf****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CheckOrganizationMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckOrganizationMemberRequest) GoString() string {
	return s.String()
}

func (s *CheckOrganizationMemberRequest) GetUserId() *string {
	return s.UserId
}

func (s *CheckOrganizationMemberRequest) SetUserId(v string) *CheckOrganizationMemberRequest {
	s.UserId = &v
	return s
}

func (s *CheckOrganizationMemberRequest) Validate() error {
	return dara.Validate(s)
}
