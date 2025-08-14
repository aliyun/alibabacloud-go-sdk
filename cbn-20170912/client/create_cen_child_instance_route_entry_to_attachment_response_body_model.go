// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenChildInstanceRouteEntryToAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateCenChildInstanceRouteEntryToAttachmentResponseBody
	GetRequestId() *string
}

type CreateCenChildInstanceRouteEntryToAttachmentResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A95A100B-3F3A-56F4-A5DE-19DB7E994807
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCenChildInstanceRouteEntryToAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCenChildInstanceRouteEntryToAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentResponseBody) SetRequestId(v string) *CreateCenChildInstanceRouteEntryToAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
