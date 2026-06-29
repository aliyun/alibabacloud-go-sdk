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
}

type DescribeRegionsRequest struct {
	// A parameter to ensure idempotence.
	//
	// > This parameter ensures that each request is processed only once. Generate a unique value on the client for each request. The value can be up to 64 ASCII characters long and must not contain non-ASCII characters.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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

func (s *DescribeRegionsRequest) SetClientToken(v string) *DescribeRegionsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
