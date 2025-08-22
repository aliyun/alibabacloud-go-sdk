// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDcdnDomainResponseBody
	GetRequestId() *string
}

type DeleteDcdnDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDcdnDomainResponseBody) SetRequestId(v string) *DeleteDcdnDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDcdnDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
