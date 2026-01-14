// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpSetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIpSetsResponseBody
	GetRequestId() *string
}

type UpdateIpSetsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 685662FF-B1CE-4D5C-A4C8-2FF3C2146BFC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpSetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpSetsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpSetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIpSetsResponseBody) SetRequestId(v string) *UpdateIpSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIpSetsResponseBody) Validate() error {
	return dara.Validate(s)
}
