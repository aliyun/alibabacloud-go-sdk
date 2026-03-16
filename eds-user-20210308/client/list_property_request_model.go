// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPropertyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *ListPropertyRequest
	GetBusinessChannel() *string
}

type ListPropertyRequest struct {
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
}

func (s ListPropertyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPropertyRequest) GoString() string {
	return s.String()
}

func (s *ListPropertyRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *ListPropertyRequest) SetBusinessChannel(v string) *ListPropertyRequest {
	s.BusinessChannel = &v
	return s
}

func (s *ListPropertyRequest) Validate() error {
	return dara.Validate(s)
}
