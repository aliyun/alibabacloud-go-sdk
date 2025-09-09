// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyInstanceADAuthServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *VerifyInstanceADAuthServerResponseBody
	GetRequestId() *string
}

type VerifyInstanceADAuthServerResponseBody struct {
	// example:
	//
	// 8F1085E3-F048-5F34-B650-F145216E4AA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyInstanceADAuthServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyInstanceADAuthServerResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyInstanceADAuthServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyInstanceADAuthServerResponseBody) SetRequestId(v string) *VerifyInstanceADAuthServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyInstanceADAuthServerResponseBody) Validate() error {
	return dara.Validate(s)
}
