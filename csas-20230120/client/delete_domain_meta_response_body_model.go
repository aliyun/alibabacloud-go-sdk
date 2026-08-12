// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDomainMetaResponseBody
	GetRequestId() *string
}

type DeleteDomainMetaResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 91DAC192-F069-5CE6-B53F-41683D6A9555
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDomainMetaResponseBody) SetRequestId(v string) *DeleteDomainMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDomainMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
