// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *AssociateResourcesResponseBody
	GetAcceleratorId() *string
	SetAssociatedResourceId(v string) *AssociateResourcesResponseBody
	GetAssociatedResourceId() *string
	SetRequestId(v string) *AssociateResourcesResponseBody
	GetRequestId() *string
}

type AssociateResourcesResponseBody struct {
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// example:
	//
	// waf_xx
	AssociatedResourceId *string `json:"AssociatedResourceId,omitempty" xml:"AssociatedResourceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DE77A7F3-3B74-41C0-A5BC-CAFD188C28B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateResourcesResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *AssociateResourcesResponseBody) GetAssociatedResourceId() *string {
	return s.AssociatedResourceId
}

func (s *AssociateResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateResourcesResponseBody) SetAcceleratorId(v string) *AssociateResourcesResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *AssociateResourcesResponseBody) SetAssociatedResourceId(v string) *AssociateResourcesResponseBody {
	s.AssociatedResourceId = &v
	return s
}

func (s *AssociateResourcesResponseBody) SetRequestId(v string) *AssociateResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
