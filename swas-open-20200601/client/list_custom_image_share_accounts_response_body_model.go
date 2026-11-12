// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomImageShareAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageShareUsers(v []*ListCustomImageShareAccountsResponseBodyImageShareUsers) *ListCustomImageShareAccountsResponseBody
	GetImageShareUsers() []*ListCustomImageShareAccountsResponseBodyImageShareUsers
	SetPageNumber(v int32) *ListCustomImageShareAccountsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomImageShareAccountsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCustomImageShareAccountsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCustomImageShareAccountsResponseBody
	GetTotalCount() *int32
}

type ListCustomImageShareAccountsResponseBody struct {
	// The information about the shared custom images.
	ImageShareUsers []*ListCustomImageShareAccountsResponseBodyImageShareUsers `json:"ImageShareUsers,omitempty" xml:"ImageShareUsers,omitempty" type:"Repeated"`
	// The page number of the paged query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instances.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomImageShareAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomImageShareAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomImageShareAccountsResponseBody) GetImageShareUsers() []*ListCustomImageShareAccountsResponseBodyImageShareUsers {
	return s.ImageShareUsers
}

func (s *ListCustomImageShareAccountsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomImageShareAccountsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomImageShareAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomImageShareAccountsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCustomImageShareAccountsResponseBody) SetImageShareUsers(v []*ListCustomImageShareAccountsResponseBodyImageShareUsers) *ListCustomImageShareAccountsResponseBody {
	s.ImageShareUsers = v
	return s
}

func (s *ListCustomImageShareAccountsResponseBody) SetPageNumber(v int32) *ListCustomImageShareAccountsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomImageShareAccountsResponseBody) SetPageSize(v int32) *ListCustomImageShareAccountsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomImageShareAccountsResponseBody) SetRequestId(v string) *ListCustomImageShareAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomImageShareAccountsResponseBody) SetTotalCount(v int32) *ListCustomImageShareAccountsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCustomImageShareAccountsResponseBody) Validate() error {
	if s.ImageShareUsers != nil {
		for _, item := range s.ImageShareUsers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomImageShareAccountsResponseBodyImageShareUsers struct {
	// The time when the custom image was shared, in the format of YYYY-MM-DD HH:mm:ss.S.
	//
	// example:
	//
	// 2024-07-15 13:59:21.0
	SharedTime *string `json:"SharedTime,omitempty" xml:"SharedTime,omitempty"`
	// The Alibaba Cloud account ID with which the custom image is shared.
	//
	// example:
	//
	// 125111425233****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListCustomImageShareAccountsResponseBodyImageShareUsers) String() string {
	return dara.Prettify(s)
}

func (s ListCustomImageShareAccountsResponseBodyImageShareUsers) GoString() string {
	return s.String()
}

func (s *ListCustomImageShareAccountsResponseBodyImageShareUsers) GetSharedTime() *string {
	return s.SharedTime
}

func (s *ListCustomImageShareAccountsResponseBodyImageShareUsers) GetUserId() *int64 {
	return s.UserId
}

func (s *ListCustomImageShareAccountsResponseBodyImageShareUsers) SetSharedTime(v string) *ListCustomImageShareAccountsResponseBodyImageShareUsers {
	s.SharedTime = &v
	return s
}

func (s *ListCustomImageShareAccountsResponseBodyImageShareUsers) SetUserId(v int64) *ListCustomImageShareAccountsResponseBodyImageShareUsers {
	s.UserId = &v
	return s
}

func (s *ListCustomImageShareAccountsResponseBodyImageShareUsers) Validate() error {
	return dara.Validate(s)
}
