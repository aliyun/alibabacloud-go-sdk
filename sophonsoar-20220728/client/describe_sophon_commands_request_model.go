// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSophonCommandsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DescribeSophonCommandsRequest
	GetName() *string
}

type DescribeSophonCommandsRequest struct {
	// The name of the command. Fuzzy match is supported.
	//
	// example:
	//
	// waf_process
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeSophonCommandsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSophonCommandsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSophonCommandsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeSophonCommandsRequest) SetName(v string) *DescribeSophonCommandsRequest {
	s.Name = &v
	return s
}

func (s *DescribeSophonCommandsRequest) Validate() error {
	return dara.Validate(s)
}
