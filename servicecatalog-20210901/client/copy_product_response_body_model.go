// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CopyProductResponseBody
	GetRequestId() *string
}

type CopyProductResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyProductResponseBody) GoString() string {
	return s.String()
}

func (s *CopyProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyProductResponseBody) SetRequestId(v string) *CopyProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyProductResponseBody) Validate() error {
	return dara.Validate(s)
}
