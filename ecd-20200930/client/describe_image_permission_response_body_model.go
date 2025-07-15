// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImagePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliUids(v []*string) *DescribeImagePermissionResponseBody
	GetAliUids() []*string
	SetRequestId(v string) *DescribeImagePermissionResponseBody
	GetRequestId() *string
}

type DescribeImagePermissionResponseBody struct {
	// The IDs of the Alibaba Cloud accounts with which the image is shared.
	AliUids []*string `json:"AliUids,omitempty" xml:"AliUids,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImagePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImagePermissionResponseBody) GetAliUids() []*string {
	return s.AliUids
}

func (s *DescribeImagePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImagePermissionResponseBody) SetAliUids(v []*string) *DescribeImagePermissionResponseBody {
	s.AliUids = v
	return s
}

func (s *DescribeImagePermissionResponseBody) SetRequestId(v string) *DescribeImagePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImagePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
