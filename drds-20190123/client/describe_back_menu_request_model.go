// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackMenuRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *DescribeBackMenuRequest
	GetDrdsInstanceId() *string
}

type DescribeBackMenuRequest struct {
	// The ID of the DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds***********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeBackMenuRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackMenuRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackMenuRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeBackMenuRequest) SetDrdsInstanceId(v string) *DescribeBackMenuRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeBackMenuRequest) Validate() error {
	return dara.Validate(s)
}
