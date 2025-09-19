// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventOperationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSecurityEventOperationsRequest
	GetLang() *string
	SetResourceOwnerId(v int64) *DescribeSecurityEventOperationsRequest
	GetResourceOwnerId() *int64
	SetSecurityEventId(v int64) *DescribeSecurityEventOperationsRequest
	GetSecurityEventId() *int64
	SetSourceIp(v string) *DescribeSecurityEventOperationsRequest
	GetSourceIp() *string
}

type DescribeSecurityEventOperationsRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the alert event that you want to handle.
	//
	// This parameter is required.
	//
	// example:
	//
	// 61352054
	SecurityEventId *int64 `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeSecurityEventOperationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventOperationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSecurityEventOperationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSecurityEventOperationsRequest) GetSecurityEventId() *int64 {
	return s.SecurityEventId
}

func (s *DescribeSecurityEventOperationsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSecurityEventOperationsRequest) SetLang(v string) *DescribeSecurityEventOperationsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSecurityEventOperationsRequest) SetResourceOwnerId(v int64) *DescribeSecurityEventOperationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSecurityEventOperationsRequest) SetSecurityEventId(v int64) *DescribeSecurityEventOperationsRequest {
	s.SecurityEventId = &v
	return s
}

func (s *DescribeSecurityEventOperationsRequest) SetSourceIp(v string) *DescribeSecurityEventOperationsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSecurityEventOperationsRequest) Validate() error {
	return dara.Validate(s)
}
