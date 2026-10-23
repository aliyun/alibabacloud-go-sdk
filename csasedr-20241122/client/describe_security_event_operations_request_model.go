// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventOperationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityEventId(v int64) *DescribeSecurityEventOperationsRequest
	GetSecurityEventId() *int64
}

type DescribeSecurityEventOperationsRequest struct {
	// The ID of the security alert event to be handled.
	//
	// This parameter is required.
	//
	// example:
	//
	// 61352054
	SecurityEventId *int64 `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
}

func (s DescribeSecurityEventOperationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventOperationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsRequest) GetSecurityEventId() *int64 {
	return s.SecurityEventId
}

func (s *DescribeSecurityEventOperationsRequest) SetSecurityEventId(v int64) *DescribeSecurityEventOperationsRequest {
	s.SecurityEventId = &v
	return s
}

func (s *DescribeSecurityEventOperationsRequest) Validate() error {
	return dara.Validate(s)
}
