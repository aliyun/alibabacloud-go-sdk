// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelServiceUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelServiceUsageResponseBody
	GetRequestId() *string
}

type CancelServiceUsageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelServiceUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelServiceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *CancelServiceUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelServiceUsageResponseBody) SetRequestId(v string) *CancelServiceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelServiceUsageResponseBody) Validate() error {
	return dara.Validate(s)
}
