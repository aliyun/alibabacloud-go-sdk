// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainToDomainGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDomainToDomainGroupResponseBody
	GetRequestId() *string
}

type UpdateDomainToDomainGroupResponseBody struct {
	// example:
	//
	// 40F46D3D-F4F3-4CCB-AC30-2DD20E32E528
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDomainToDomainGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainToDomainGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDomainToDomainGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDomainToDomainGroupResponseBody) SetRequestId(v string) *UpdateDomainToDomainGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDomainToDomainGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
