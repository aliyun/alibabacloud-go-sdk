// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachCenChildInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachCenChildInstanceResponseBody
	GetRequestId() *string
}

type DetachCenChildInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0C2EE7A8-74D4-4081-8236-CEBDE3BBCF50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachCenChildInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachCenChildInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DetachCenChildInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachCenChildInstanceResponseBody) SetRequestId(v string) *DetachCenChildInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachCenChildInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
