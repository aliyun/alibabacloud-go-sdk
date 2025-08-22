// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnIpaDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDcdnIpaDomainResponseBody
	GetRequestId() *string
}

type DeleteDcdnIpaDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 94E3559F-7B6A-4A5E-AFFD-44E2A208A249
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnIpaDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnIpaDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnIpaDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDcdnIpaDomainResponseBody) SetRequestId(v string) *DeleteDcdnIpaDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDcdnIpaDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
