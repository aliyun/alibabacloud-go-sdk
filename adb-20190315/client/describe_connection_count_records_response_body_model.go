// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConnectionCountRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessIpRecords(v []*DescribeConnectionCountRecordsResponseBodyAccessIpRecords) *DescribeConnectionCountRecordsResponseBody
	GetAccessIpRecords() []*DescribeConnectionCountRecordsResponseBodyAccessIpRecords
	SetDBClusterId(v string) *DescribeConnectionCountRecordsResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *DescribeConnectionCountRecordsResponseBody
	GetRequestId() *string
	SetUserRecords(v []*DescribeConnectionCountRecordsResponseBodyUserRecords) *DescribeConnectionCountRecordsResponseBody
	GetUserRecords() []*DescribeConnectionCountRecordsResponseBodyUserRecords
}

type DescribeConnectionCountRecordsResponseBody struct {
	// The queried client IP addresses.
	AccessIpRecords []*DescribeConnectionCountRecordsResponseBodyAccessIpRecords `json:"AccessIpRecords,omitempty" xml:"AccessIpRecords,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// example:
	//
	// am-bp1jj9xqft1po****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 562C7F89-FBE6-4A04-AAAA-7EBC25******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried database accounts.
	UserRecords []*DescribeConnectionCountRecordsResponseBodyUserRecords `json:"UserRecords,omitempty" xml:"UserRecords,omitempty" type:"Repeated"`
}

func (s DescribeConnectionCountRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConnectionCountRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConnectionCountRecordsResponseBody) GetAccessIpRecords() []*DescribeConnectionCountRecordsResponseBodyAccessIpRecords {
	return s.AccessIpRecords
}

func (s *DescribeConnectionCountRecordsResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeConnectionCountRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConnectionCountRecordsResponseBody) GetUserRecords() []*DescribeConnectionCountRecordsResponseBodyUserRecords {
	return s.UserRecords
}

func (s *DescribeConnectionCountRecordsResponseBody) SetAccessIpRecords(v []*DescribeConnectionCountRecordsResponseBodyAccessIpRecords) *DescribeConnectionCountRecordsResponseBody {
	s.AccessIpRecords = v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBody) SetDBClusterId(v string) *DescribeConnectionCountRecordsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBody) SetRequestId(v string) *DescribeConnectionCountRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBody) SetUserRecords(v []*DescribeConnectionCountRecordsResponseBodyUserRecords) *DescribeConnectionCountRecordsResponseBody {
	s.UserRecords = v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBody) Validate() error {
	if s.AccessIpRecords != nil {
		for _, item := range s.AccessIpRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserRecords != nil {
		for _, item := range s.UserRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeConnectionCountRecordsResponseBodyAccessIpRecords struct {
	// The IP address of the client.
	//
	// example:
	//
	// 42.120.XX.XX
	AccessIp *string `json:"AccessIp,omitempty" xml:"AccessIp,omitempty"`
	// The number of connections.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeConnectionCountRecordsResponseBodyAccessIpRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeConnectionCountRecordsResponseBodyAccessIpRecords) GoString() string {
	return s.String()
}

func (s *DescribeConnectionCountRecordsResponseBodyAccessIpRecords) GetAccessIp() *string {
	return s.AccessIp
}

func (s *DescribeConnectionCountRecordsResponseBodyAccessIpRecords) GetCount() *int64 {
	return s.Count
}

func (s *DescribeConnectionCountRecordsResponseBodyAccessIpRecords) SetAccessIp(v string) *DescribeConnectionCountRecordsResponseBodyAccessIpRecords {
	s.AccessIp = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBodyAccessIpRecords) SetCount(v int64) *DescribeConnectionCountRecordsResponseBodyAccessIpRecords {
	s.Count = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBodyAccessIpRecords) Validate() error {
	return dara.Validate(s)
}

type DescribeConnectionCountRecordsResponseBodyUserRecords struct {
	// The number of connections.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The username of the database account.
	//
	// example:
	//
	// test
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeConnectionCountRecordsResponseBodyUserRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeConnectionCountRecordsResponseBodyUserRecords) GoString() string {
	return s.String()
}

func (s *DescribeConnectionCountRecordsResponseBodyUserRecords) GetCount() *int64 {
	return s.Count
}

func (s *DescribeConnectionCountRecordsResponseBodyUserRecords) GetUser() *string {
	return s.User
}

func (s *DescribeConnectionCountRecordsResponseBodyUserRecords) SetCount(v int64) *DescribeConnectionCountRecordsResponseBodyUserRecords {
	s.Count = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBodyUserRecords) SetUser(v string) *DescribeConnectionCountRecordsResponseBodyUserRecords {
	s.User = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBodyUserRecords) Validate() error {
	return dara.Validate(s)
}
