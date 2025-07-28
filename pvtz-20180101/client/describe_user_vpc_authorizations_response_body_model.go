// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserVpcAuthorizationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeUserVpcAuthorizationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeUserVpcAuthorizationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeUserVpcAuthorizationsResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *DescribeUserVpcAuthorizationsResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeUserVpcAuthorizationsResponseBody
	GetTotalPages() *int32
	SetUsers(v []*DescribeUserVpcAuthorizationsResponseBodyUsers) *DescribeUserVpcAuthorizationsResponseBody
	GetUsers() []*DescribeUserVpcAuthorizationsResponseBodyUsers
}

type DescribeUserVpcAuthorizationsResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 46973D4C-E3E4-4ABA-9190-9A9DE406C7E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 5
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	// The Alibaba Cloud accounts to which the permissions on the resources are granted.
	Users []*DescribeUserVpcAuthorizationsResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s DescribeUserVpcAuthorizationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserVpcAuthorizationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserVpcAuthorizationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeUserVpcAuthorizationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeUserVpcAuthorizationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserVpcAuthorizationsResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeUserVpcAuthorizationsResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeUserVpcAuthorizationsResponseBody) GetUsers() []*DescribeUserVpcAuthorizationsResponseBodyUsers {
	return s.Users
}

func (s *DescribeUserVpcAuthorizationsResponseBody) SetPageNumber(v int32) *DescribeUserVpcAuthorizationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBody) SetPageSize(v int32) *DescribeUserVpcAuthorizationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBody) SetRequestId(v string) *DescribeUserVpcAuthorizationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBody) SetTotalItems(v int32) *DescribeUserVpcAuthorizationsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBody) SetTotalPages(v int32) *DescribeUserVpcAuthorizationsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBody) SetUsers(v []*DescribeUserVpcAuthorizationsResponseBodyUsers) *DescribeUserVpcAuthorizationsResponseBody {
	s.Users = v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUserVpcAuthorizationsResponseBodyUsers struct {
	// The authorization scope. Valid values:
	//
	// 	- NORMAL: general authorization
	//
	// 	- CLOUD_PRODUCT: cloud service-related authorization
	//
	// example:
	//
	// NORMAL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The name of the Alibaba Cloud account to which the permissions on the resources are granted.
	//
	// example:
	//
	// alidn****@test.com
	AuthorizedAliyunId *string `json:"AuthorizedAliyunId,omitempty" xml:"AuthorizedAliyunId,omitempty"`
	// The ID of the Alibaba Cloud account to which the permissions on the resources are granted.
	//
	// example:
	//
	// 141339776561****
	AuthorizedUserId *int64 `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
	// The time when the authorization was performed. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-05-08T02:31Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the authorization was performed. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1672740294000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribeUserVpcAuthorizationsResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserVpcAuthorizationsResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *DescribeUserVpcAuthorizationsResponseBodyUsers) GetAuthType() *string {
	return s.AuthType
}

func (s *DescribeUserVpcAuthorizationsResponseBodyUsers) GetAuthorizedAliyunId() *string {
	return s.AuthorizedAliyunId
}

func (s *DescribeUserVpcAuthorizationsResponseBodyUsers) GetAuthorizedUserId() *int64 {
	return s.AuthorizedUserId
}

func (s *DescribeUserVpcAuthorizationsResponseBodyUsers) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeUserVpcAuthorizationsResponseBodyUsers) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeUserVpcAuthorizationsResponseBodyUsers) SetAuthType(v string) *DescribeUserVpcAuthorizationsResponseBodyUsers {
	s.AuthType = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBodyUsers) SetAuthorizedAliyunId(v string) *DescribeUserVpcAuthorizationsResponseBodyUsers {
	s.AuthorizedAliyunId = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBodyUsers) SetAuthorizedUserId(v int64) *DescribeUserVpcAuthorizationsResponseBodyUsers {
	s.AuthorizedUserId = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBodyUsers) SetCreateTime(v string) *DescribeUserVpcAuthorizationsResponseBodyUsers {
	s.CreateTime = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBodyUsers) SetCreateTimestamp(v int64) *DescribeUserVpcAuthorizationsResponseBodyUsers {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBodyUsers) Validate() error {
	return dara.Validate(s)
}
