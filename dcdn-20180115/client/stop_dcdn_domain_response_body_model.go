// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDcdnDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopDcdnDomainResponseBody
	GetRequestId() *string
}

type StopDcdnDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDcdnDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDcdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StopDcdnDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDcdnDomainResponseBody) SetRequestId(v string) *StopDcdnDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDcdnDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
