// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteServiceResponseBody
	GetRequestId() *string
}

type DeleteServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceResponseBody) SetRequestId(v string) *DeleteServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
