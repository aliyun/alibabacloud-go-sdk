// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOfficeSiteAcceleratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *CreateOfficeSiteAcceleratorResponseBody
	GetAcceleratorId() *string
	SetRequestId(v string) *CreateOfficeSiteAcceleratorResponseBody
	GetRequestId() *string
}

type CreateOfficeSiteAcceleratorResponseBody struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp1qxuk10jceqw3zb***p
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOfficeSiteAcceleratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOfficeSiteAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOfficeSiteAcceleratorResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateOfficeSiteAcceleratorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOfficeSiteAcceleratorResponseBody) SetAcceleratorId(v string) *CreateOfficeSiteAcceleratorResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *CreateOfficeSiteAcceleratorResponseBody) SetRequestId(v string) *CreateOfficeSiteAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOfficeSiteAcceleratorResponseBody) Validate() error {
	return dara.Validate(s)
}
