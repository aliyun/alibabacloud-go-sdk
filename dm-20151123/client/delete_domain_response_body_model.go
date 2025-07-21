// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDomainResponseBody
	GetRequestId() *string
}

type DeleteDomainResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// E3DFF97B-00CF-5333-8125-3D6819471984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDomainResponseBody) SetRequestId(v string) *DeleteDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
