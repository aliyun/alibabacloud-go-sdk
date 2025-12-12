// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventMetaInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeEventMetaInfoRequest
	GetRegionId() *string
	SetSecurityToken(v string) *DescribeEventMetaInfoRequest
	GetSecurityToken() *string
	SetSourceCode(v string) *DescribeEventMetaInfoRequest
	GetSourceCode() *string
}

type DescribeEventMetaInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// Event.EventCode
	SourceCode *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
}

func (s DescribeEventMetaInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventMetaInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventMetaInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEventMetaInfoRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeEventMetaInfoRequest) GetSourceCode() *string {
	return s.SourceCode
}

func (s *DescribeEventMetaInfoRequest) SetRegionId(v string) *DescribeEventMetaInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEventMetaInfoRequest) SetSecurityToken(v string) *DescribeEventMetaInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeEventMetaInfoRequest) SetSourceCode(v string) *DescribeEventMetaInfoRequest {
	s.SourceCode = &v
	return s
}

func (s *DescribeEventMetaInfoRequest) Validate() error {
	return dara.Validate(s)
}
