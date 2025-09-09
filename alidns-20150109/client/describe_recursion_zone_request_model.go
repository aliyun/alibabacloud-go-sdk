// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecursionZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetZoneId(v string) *DescribeRecursionZoneRequest
	GetZoneId() *string
}

type DescribeRecursionZoneRequest struct {
	// example:
	//
	// 169783221000012
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRecursionZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecursionZoneRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecursionZoneRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRecursionZoneRequest) SetZoneId(v string) *DescribeRecursionZoneRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeRecursionZoneRequest) Validate() error {
	return dara.Validate(s)
}
