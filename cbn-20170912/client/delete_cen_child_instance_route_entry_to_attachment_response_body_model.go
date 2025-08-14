// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenChildInstanceRouteEntryToAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCenChildInstanceRouteEntryToAttachmentResponseBody
	GetRequestId() *string
}

type DeleteCenChildInstanceRouteEntryToAttachmentResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 250E717B-9823-5FD8-A1C6-5714234FB825
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCenChildInstanceRouteEntryToAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenChildInstanceRouteEntryToAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentResponseBody) SetRequestId(v string) *DeleteCenChildInstanceRouteEntryToAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
