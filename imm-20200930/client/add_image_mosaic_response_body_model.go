// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageMosaicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddImageMosaicResponseBody
	GetRequestId() *string
}

type AddImageMosaicResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CA995EFD-083D-4F40-BE8A-BDF75FF*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddImageMosaicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddImageMosaicResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageMosaicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddImageMosaicResponseBody) SetRequestId(v string) *AddImageMosaicResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddImageMosaicResponseBody) Validate() error {
	return dara.Validate(s)
}
