// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCdnDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartCdnDomainResponseBody
	GetRequestId() *string
}

type StartCdnDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartCdnDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartCdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StartCdnDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartCdnDomainResponseBody) SetRequestId(v string) *StartCdnDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCdnDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
