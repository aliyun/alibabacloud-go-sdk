// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *DescribeUserRequest
	GetBusinessChannel() *string
	SetEndUserId(v string) *DescribeUserRequest
	GetEndUserId() *string
	SetRequireExtraAttributes(v []*string) *DescribeUserRequest
	GetRequireExtraAttributes() []*string
}

type DescribeUserRequest struct {
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// example:
	//
	// alex
	EndUserId              *string   `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	RequireExtraAttributes []*string `json:"RequireExtraAttributes,omitempty" xml:"RequireExtraAttributes,omitempty" type:"Repeated"`
}

func (s DescribeUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *DescribeUserRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeUserRequest) GetRequireExtraAttributes() []*string {
	return s.RequireExtraAttributes
}

func (s *DescribeUserRequest) SetBusinessChannel(v string) *DescribeUserRequest {
	s.BusinessChannel = &v
	return s
}

func (s *DescribeUserRequest) SetEndUserId(v string) *DescribeUserRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeUserRequest) SetRequireExtraAttributes(v []*string) *DescribeUserRequest {
	s.RequireExtraAttributes = v
	return s
}

func (s *DescribeUserRequest) Validate() error {
	return dara.Validate(s)
}
