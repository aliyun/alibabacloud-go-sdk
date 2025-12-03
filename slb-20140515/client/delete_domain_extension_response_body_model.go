// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainExtensionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDomainExtensionResponseBody
	GetRequestId() *string
}

type DeleteDomainExtensionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 149A2470-F010-4437-BF68-343D5099C19D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainExtensionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainExtensionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainExtensionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDomainExtensionResponseBody) SetRequestId(v string) *DeleteDomainExtensionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDomainExtensionResponseBody) Validate() error {
	return dara.Validate(s)
}
