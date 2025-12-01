// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeRegionsRequest
	GetClientToken() *string
	SetOwnerId(v string) *DescribeRegionsRequest
	GetOwnerId() *string
}

type DescribeRegionsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeRegionsRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeRegionsRequest) SetClientToken(v string) *DescribeRegionsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerId(v string) *DescribeRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
