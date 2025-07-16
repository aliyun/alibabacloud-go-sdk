// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCdnDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopCdnDomainResponseBody
	GetRequestId() *string
}

type StopCdnDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopCdnDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopCdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StopCdnDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopCdnDomainResponseBody) SetRequestId(v string) *StopCdnDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopCdnDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
