// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateServiceUsageResponseBody
	GetRequestId() *string
}

type UpdateServiceUsageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceUsageResponseBody) SetRequestId(v string) *UpdateServiceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceUsageResponseBody) Validate() error {
	return dara.Validate(s)
}
