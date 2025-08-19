// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDestination interface {
	dara.Model
	String() string
	GoString() string
	SetDestination(v string) *Destination
	GetDestination() *string
}

type Destination struct {
	// example:
	//
	// acs:fc:cn-shanghai:xxx:functions/f1
	Destination *string `json:"destination,omitempty" xml:"destination,omitempty"`
}

func (s Destination) String() string {
	return dara.Prettify(s)
}

func (s Destination) GoString() string {
	return s.String()
}

func (s *Destination) GetDestination() *string {
	return s.Destination
}

func (s *Destination) SetDestination(v string) *Destination {
	s.Destination = &v
	return s
}

func (s *Destination) Validate() error {
	return dara.Validate(s)
}
