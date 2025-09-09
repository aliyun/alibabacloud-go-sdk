// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSwitchAzoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *DescribeInstanceSwitchAzoneRequest
	GetDrdsInstanceId() *string
}

type DescribeInstanceSwitchAzoneRequest struct {
	// The ID of the DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdsxxxxxxxxxxxx
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeInstanceSwitchAzoneRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSwitchAzoneRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchAzoneRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeInstanceSwitchAzoneRequest) SetDrdsInstanceId(v string) *DescribeInstanceSwitchAzoneRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeInstanceSwitchAzoneRequest) Validate() error {
	return dara.Validate(s)
}
