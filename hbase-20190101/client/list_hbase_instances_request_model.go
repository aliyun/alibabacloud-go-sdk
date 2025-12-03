// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHBaseInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVpcId(v string) *ListHBaseInstancesRequest
	GetVpcId() *string
}

type ListHBaseInstancesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// vpc-t4nx81tmlixcq5i****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListHBaseInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHBaseInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListHBaseInstancesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListHBaseInstancesRequest) SetVpcId(v string) *ListHBaseInstancesRequest {
	s.VpcId = &v
	return s
}

func (s *ListHBaseInstancesRequest) Validate() error {
	return dara.Validate(s)
}
