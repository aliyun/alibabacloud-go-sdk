// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *DescribeInstanceAccountsRequest
	GetDrdsInstanceId() *string
}

type DescribeInstanceAccountsRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds*********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeInstanceAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeInstanceAccountsRequest) SetDrdsInstanceId(v string) *DescribeInstanceAccountsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeInstanceAccountsRequest) Validate() error {
	return dara.Validate(s)
}
