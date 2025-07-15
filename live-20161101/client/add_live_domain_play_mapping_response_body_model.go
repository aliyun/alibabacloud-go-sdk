// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveDomainPlayMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveDomainPlayMappingResponseBody
	GetRequestId() *string
}

type AddLiveDomainPlayMappingResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLiveDomainPlayMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveDomainPlayMappingResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveDomainPlayMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveDomainPlayMappingResponseBody) SetRequestId(v string) *AddLiveDomainPlayMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveDomainPlayMappingResponseBody) Validate() error {
	return dara.Validate(s)
}
