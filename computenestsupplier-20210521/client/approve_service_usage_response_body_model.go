// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveServiceUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApproveServiceUsageResponseBody
	GetRequestId() *string
}

type ApproveServiceUsageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveServiceUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApproveServiceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveServiceUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApproveServiceUsageResponseBody) SetRequestId(v string) *ApproveServiceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApproveServiceUsageResponseBody) Validate() error {
	return dara.Validate(s)
}
