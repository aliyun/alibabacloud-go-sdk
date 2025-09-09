// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectServiceUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RejectServiceUsageResponseBody
	GetRequestId() *string
}

type RejectServiceUsageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RejectServiceUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RejectServiceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *RejectServiceUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RejectServiceUsageResponseBody) SetRequestId(v string) *RejectServiceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RejectServiceUsageResponseBody) Validate() error {
	return dara.Validate(s)
}
