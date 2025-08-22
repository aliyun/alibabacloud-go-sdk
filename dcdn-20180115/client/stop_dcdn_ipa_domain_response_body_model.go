// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDcdnIpaDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopDcdnIpaDomainResponseBody
	GetRequestId() *string
}

type StopDcdnIpaDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDcdnIpaDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDcdnIpaDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StopDcdnIpaDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDcdnIpaDomainResponseBody) SetRequestId(v string) *StopDcdnIpaDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDcdnIpaDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
