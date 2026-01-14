// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DisassociateResourcesResponseBody
	GetAcceleratorId() *string
	SetAssociatedResourceId(v string) *DisassociateResourcesResponseBody
	GetAssociatedResourceId() *string
	SetRequestId(v string) *DisassociateResourcesResponseBody
	GetRequestId() *string
}

type DisassociateResourcesResponseBody struct {
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// example:
	//
	// c66d65f411b9143bab253bfef61c03c48
	AssociatedResourceId *string `json:"AssociatedResourceId,omitempty" xml:"AssociatedResourceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DE77A7F3-3B74-41C0-A5BC-CAFD188C28B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisassociateResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisassociateResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateResourcesResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DisassociateResourcesResponseBody) GetAssociatedResourceId() *string {
	return s.AssociatedResourceId
}

func (s *DisassociateResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisassociateResourcesResponseBody) SetAcceleratorId(v string) *DisassociateResourcesResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DisassociateResourcesResponseBody) SetAssociatedResourceId(v string) *DisassociateResourcesResponseBody {
	s.AssociatedResourceId = &v
	return s
}

func (s *DisassociateResourcesResponseBody) SetRequestId(v string) *DisassociateResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisassociateResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
