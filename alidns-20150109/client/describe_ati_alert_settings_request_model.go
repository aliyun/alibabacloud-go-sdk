// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAtiAlertSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeAtiAlertSettingsRequest
	GetClientToken() *string
}

type DescribeAtiAlertSettingsRequest struct {
	// The token that ensures the idempotence of the request. Generate a unique value from your client. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cC.....
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeAtiAlertSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiAlertSettingsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAtiAlertSettingsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeAtiAlertSettingsRequest) SetClientToken(v string) *DescribeAtiAlertSettingsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeAtiAlertSettingsRequest) Validate() error {
	return dara.Validate(s)
}
