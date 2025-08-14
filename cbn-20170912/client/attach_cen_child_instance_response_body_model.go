// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachCenChildInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachCenChildInstanceResponseBody
	GetRequestId() *string
}

type AttachCenChildInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A278B8A6-A5B8-4FDE-9F70-95F0F6A1D68A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachCenChildInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachCenChildInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AttachCenChildInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachCenChildInstanceResponseBody) SetRequestId(v string) *AttachCenChildInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachCenChildInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
