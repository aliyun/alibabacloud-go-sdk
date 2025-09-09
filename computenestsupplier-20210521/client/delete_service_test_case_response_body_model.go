// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceTestCaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteServiceTestCaseResponseBody
	GetRequestId() *string
}

type DeleteServiceTestCaseResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceTestCaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceTestCaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceTestCaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceTestCaseResponseBody) SetRequestId(v string) *DeleteServiceTestCaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceTestCaseResponseBody) Validate() error {
	return dara.Validate(s)
}
