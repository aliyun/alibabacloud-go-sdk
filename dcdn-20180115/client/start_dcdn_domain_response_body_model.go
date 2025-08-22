// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDcdnDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartDcdnDomainResponseBody
	GetRequestId() *string
}

type StartDcdnDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDcdnDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDcdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StartDcdnDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDcdnDomainResponseBody) SetRequestId(v string) *StartDcdnDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDcdnDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
