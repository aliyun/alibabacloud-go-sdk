// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopWebApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *StopWebApplicationRequest
	GetNamespaceId() *string
}

type StopWebApplicationRequest struct {
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s StopWebApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s StopWebApplicationRequest) GoString() string {
	return s.String()
}

func (s *StopWebApplicationRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *StopWebApplicationRequest) SetNamespaceId(v string) *StopWebApplicationRequest {
	s.NamespaceId = &v
	return s
}

func (s *StopWebApplicationRequest) Validate() error {
	return dara.Validate(s)
}
