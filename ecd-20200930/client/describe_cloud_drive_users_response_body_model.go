// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDriveUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudDriveUsers(v []*DescribeCloudDriveUsersResponseBodyCloudDriveUsers) *DescribeCloudDriveUsersResponseBody
	GetCloudDriveUsers() []*DescribeCloudDriveUsersResponseBodyCloudDriveUsers
	SetNextToken(v string) *DescribeCloudDriveUsersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeCloudDriveUsersResponseBody
	GetRequestId() *string
}

type DescribeCloudDriveUsersResponseBody struct {
	// A list of user personal drives.
	CloudDriveUsers []*DescribeCloudDriveUsersResponseBodyCloudDriveUsers `json:"CloudDriveUsers,omitempty" xml:"CloudDriveUsers,omitempty" type:"Repeated"`
	// The token for the next page of results. An empty value indicates that all results have been returned.
	//
	// example:
	//
	// aGN4YzAxQGNuLWhhbmd6aG91LjExNzU5NTMyNjgzMTQ1****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F083AAE5-7AA9-53BB-9060-AFFB2C18****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCloudDriveUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDriveUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudDriveUsersResponseBody) GetCloudDriveUsers() []*DescribeCloudDriveUsersResponseBodyCloudDriveUsers {
	return s.CloudDriveUsers
}

func (s *DescribeCloudDriveUsersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCloudDriveUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudDriveUsersResponseBody) SetCloudDriveUsers(v []*DescribeCloudDriveUsersResponseBodyCloudDriveUsers) *DescribeCloudDriveUsersResponseBody {
	s.CloudDriveUsers = v
	return s
}

func (s *DescribeCloudDriveUsersResponseBody) SetNextToken(v string) *DescribeCloudDriveUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeCloudDriveUsersResponseBody) SetRequestId(v string) *DescribeCloudDriveUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudDriveUsersResponseBody) Validate() error {
	if s.CloudDriveUsers != nil {
		for _, item := range s.CloudDriveUsers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudDriveUsersResponseBodyCloudDriveUsers struct {
	// The ID of the user personal drive.
	//
	// example:
	//
	// 8
	DriveId *string `json:"DriveId,omitempty" xml:"DriveId,omitempty"`
	// The ID of the end user.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The status of the user personal drive.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Maximum storage capacity for the user’s personal drive, in bytes.
	//
	// example:
	//
	// 104857600
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
	// The used storage space, in bytes.
	//
	// example:
	//
	// 10485760
	UsedSize *int64 `json:"UsedSize,omitempty" xml:"UsedSize,omitempty"`
	// The internal ID of the user.
	//
	// example:
	//
	// alice@cn-shanghai.148875033399****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The name of the end user.
	//
	// example:
	//
	// alice
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeCloudDriveUsersResponseBodyCloudDriveUsers) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDriveUsersResponseBodyCloudDriveUsers) GoString() string {
	return s.String()
}

func (s *DescribeCloudDriveUsersResponseBodyCloudDriveUsers) GetDriveId() *string {
	return s.DriveId
}

func (s *DescribeCloudDriveUsersResponseBodyCloudDriveUsers) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeCloudDriveUsersResponseBodyCloudDriveUsers) GetStatus() *string {
	return s.Status
}

func (s *DescribeCloudDriveUsersResponseBodyCloudDriveUsers) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *DescribeCloudDriveUsersResponseBodyCloudDriveUsers) GetUsedSize() *int64 {
	return s.UsedSize
}

func (s *DescribeCloudDriveUsersResponseBodyCloudDriveUsers) GetUserId() *string {
	return s.UserId
}

func (s *DescribeCloudDriveUsersResponseBodyCloudDriveUsers) GetUserName() *string {
	return s.UserName
}

func (s *DescribeCloudDriveUsersResponseBodyCloudDriveUsers) SetDriveId(v string) *DescribeCloudDriveUsersResponseBodyCloudDriveUsers {
	s.DriveId = &v
	return s
}

func (s *DescribeCloudDriveUsersResponseBodyCloudDriveUsers) SetEndUserId(v string) *DescribeCloudDriveUsersResponseBodyCloudDriveUsers {
	s.EndUserId = &v
	return s
}

func (s *DescribeCloudDriveUsersResponseBodyCloudDriveUsers) SetStatus(v string) *DescribeCloudDriveUsersResponseBodyCloudDriveUsers {
	s.Status = &v
	return s
}

func (s *DescribeCloudDriveUsersResponseBodyCloudDriveUsers) SetTotalSize(v int64) *DescribeCloudDriveUsersResponseBodyCloudDriveUsers {
	s.TotalSize = &v
	return s
}

func (s *DescribeCloudDriveUsersResponseBodyCloudDriveUsers) SetUsedSize(v int64) *DescribeCloudDriveUsersResponseBodyCloudDriveUsers {
	s.UsedSize = &v
	return s
}

func (s *DescribeCloudDriveUsersResponseBodyCloudDriveUsers) SetUserId(v string) *DescribeCloudDriveUsersResponseBodyCloudDriveUsers {
	s.UserId = &v
	return s
}

func (s *DescribeCloudDriveUsersResponseBodyCloudDriveUsers) SetUserName(v string) *DescribeCloudDriveUsersResponseBodyCloudDriveUsers {
	s.UserName = &v
	return s
}

func (s *DescribeCloudDriveUsersResponseBodyCloudDriveUsers) Validate() error {
	return dara.Validate(s)
}
