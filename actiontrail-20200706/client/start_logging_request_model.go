// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLoggingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *StartLoggingRequest
	GetName() *string
}

type StartLoggingRequest struct {
	// The name of the trail that you want to enable.
	//
	// The name must be 6 to 36 characters in length, and can contain lowercase letters, digits, hyphens (-), and underscores (_). It must start with a lowercase letter.
	//
	// > The name must be unique within your Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// trail-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s StartLoggingRequest) String() string {
	return dara.Prettify(s)
}

func (s StartLoggingRequest) GoString() string {
	return s.String()
}

func (s *StartLoggingRequest) GetName() *string {
	return s.Name
}

func (s *StartLoggingRequest) SetName(v string) *StartLoggingRequest {
	s.Name = &v
	return s
}

func (s *StartLoggingRequest) Validate() error {
	return dara.Validate(s)
}
