// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeLiveDomainResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeLiveDomainResourceGroupResponseBody
	GetRequestId() *string
}

type ChangeLiveDomainResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******8F-F82F-10E2-BAE1-A036FD******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeLiveDomainResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeLiveDomainResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeLiveDomainResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeLiveDomainResourceGroupResponseBody) SetRequestId(v string) *ChangeLiveDomainResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeLiveDomainResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
