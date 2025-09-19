// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogShipperStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *DescribeLogShipperStatusRequest
	GetFrom() *string
}

type DescribeLogShipperStatusRequest struct {
	// The ID of the request source. Set the value to **sas**.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
}

func (s DescribeLogShipperStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogShipperStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogShipperStatusRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeLogShipperStatusRequest) SetFrom(v string) *DescribeLogShipperStatusRequest {
	s.From = &v
	return s
}

func (s *DescribeLogShipperStatusRequest) Validate() error {
	return dara.Validate(s)
}
