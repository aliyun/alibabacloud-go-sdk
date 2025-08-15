// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLoggingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *StopLoggingRequest
	GetName() *string
}

type StopLoggingRequest struct {
	// The name of the trail that you want to disable.
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

func (s StopLoggingRequest) String() string {
	return dara.Prettify(s)
}

func (s StopLoggingRequest) GoString() string {
	return s.String()
}

func (s *StopLoggingRequest) GetName() *string {
	return s.Name
}

func (s *StopLoggingRequest) SetName(v string) *StopLoggingRequest {
	s.Name = &v
	return s
}

func (s *StopLoggingRequest) Validate() error {
	return dara.Validate(s)
}
