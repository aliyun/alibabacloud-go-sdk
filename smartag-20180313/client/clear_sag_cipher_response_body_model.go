// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearSagCipherResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ClearSagCipherResponseBody
	GetRequestId() *string
}

type ClearSagCipherResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3712F0B2-721E-4FBF-BBEF-888E3BFE0A20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClearSagCipherResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClearSagCipherResponseBody) GoString() string {
	return s.String()
}

func (s *ClearSagCipherResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClearSagCipherResponseBody) SetRequestId(v string) *ClearSagCipherResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearSagCipherResponseBody) Validate() error {
	return dara.Validate(s)
}
