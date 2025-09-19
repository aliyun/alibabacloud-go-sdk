// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetLogShipperResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetLogShipperResponseBody
	GetRequestId() *string
}

type ResetLogShipperResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D2E31293-DAAC-574B-B816-A18EA0A6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetLogShipperResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetLogShipperResponseBody) GoString() string {
	return s.String()
}

func (s *ResetLogShipperResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetLogShipperResponseBody) SetRequestId(v string) *ResetLogShipperResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetLogShipperResponseBody) Validate() error {
	return dara.Validate(s)
}
