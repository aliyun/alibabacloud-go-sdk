// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachCenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachCenResponseBody
	GetRequestId() *string
}

type AttachCenResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachCenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachCenResponseBody) GoString() string {
	return s.String()
}

func (s *AttachCenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachCenResponseBody) SetRequestId(v string) *AttachCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachCenResponseBody) Validate() error {
	return dara.Validate(s)
}
