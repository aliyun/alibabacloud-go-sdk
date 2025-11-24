// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsersWithPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUsersWithPermissionsResponseBody
	GetRequestId() *string
	SetUIDs(v []*string) *DescribeUsersWithPermissionsResponseBody
	GetUIDs() []*string
}

type DescribeUsersWithPermissionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 12B94024-C241-13CD-BA44-6106DF1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of the IDs of the RAM users or RAM roles to which an RBAC role is assigned.
	UIDs []*string `json:"UIDs,omitempty" xml:"UIDs,omitempty" type:"Repeated"`
}

func (s DescribeUsersWithPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersWithPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsersWithPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUsersWithPermissionsResponseBody) GetUIDs() []*string {
	return s.UIDs
}

func (s *DescribeUsersWithPermissionsResponseBody) SetRequestId(v string) *DescribeUsersWithPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsersWithPermissionsResponseBody) SetUIDs(v []*string) *DescribeUsersWithPermissionsResponseBody {
	s.UIDs = v
	return s
}

func (s *DescribeUsersWithPermissionsResponseBody) Validate() error {
	return dara.Validate(s)
}
