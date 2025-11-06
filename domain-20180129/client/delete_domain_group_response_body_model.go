// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDomainGroupResponseBody
	GetRequestId() *string
}

type DeleteDomainGroupResponseBody struct {
	// example:
	//
	// 40F46D3D-F4F3-4CCB-AC30-2DD20E32E528
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDomainGroupResponseBody) SetRequestId(v string) *DeleteDomainGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDomainGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
