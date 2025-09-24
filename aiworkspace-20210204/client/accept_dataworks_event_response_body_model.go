// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptDataworksEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AcceptDataworksEventResponseBody
	GetRequestId() *string
}

type AcceptDataworksEventResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ADF6D849-*****-7E7030F0CE53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AcceptDataworksEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AcceptDataworksEventResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptDataworksEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AcceptDataworksEventResponseBody) SetRequestId(v string) *AcceptDataworksEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *AcceptDataworksEventResponseBody) Validate() error {
	return dara.Validate(s)
}
