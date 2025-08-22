// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDcdnIpaDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartDcdnIpaDomainResponseBody
	GetRequestId() *string
}

type StartDcdnIpaDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDcdnIpaDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDcdnIpaDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StartDcdnIpaDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDcdnIpaDomainResponseBody) SetRequestId(v string) *StartDcdnIpaDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDcdnIpaDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
