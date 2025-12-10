// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeForwardEntryAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForwardEntryId(v string) *DescribeForwardEntryAttributeRequest
	GetForwardEntryId() *string
}

type DescribeForwardEntryAttributeRequest struct {
	// The ID of the DNAT entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// fwd-5tfi6f0rutmd00xrhkag7****
	ForwardEntryId *string `json:"ForwardEntryId,omitempty" xml:"ForwardEntryId,omitempty"`
}

func (s DescribeForwardEntryAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeForwardEntryAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeForwardEntryAttributeRequest) GetForwardEntryId() *string {
	return s.ForwardEntryId
}

func (s *DescribeForwardEntryAttributeRequest) SetForwardEntryId(v string) *DescribeForwardEntryAttributeRequest {
	s.ForwardEntryId = &v
	return s
}

func (s *DescribeForwardEntryAttributeRequest) Validate() error {
	return dara.Validate(s)
}
