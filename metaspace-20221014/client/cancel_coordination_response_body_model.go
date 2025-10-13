// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCoordinationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelCoordinationResponseBody
	GetRequestId() *string
}

type CancelCoordinationResponseBody struct {
	// example:
	//
	// AD2D0761-1FE5-549D-B169******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelCoordinationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelCoordinationResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCoordinationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelCoordinationResponseBody) SetRequestId(v string) *CancelCoordinationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelCoordinationResponseBody) Validate() error {
	return dara.Validate(s)
}
