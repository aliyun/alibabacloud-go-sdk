// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignalResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SignalResourceResponseBody
	GetRequestId() *string
}

type SignalResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SignalResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SignalResourceResponseBody) GoString() string {
	return s.String()
}

func (s *SignalResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SignalResourceResponseBody) SetRequestId(v string) *SignalResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SignalResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
