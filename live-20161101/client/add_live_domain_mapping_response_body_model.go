// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveDomainMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveDomainMappingResponseBody
	GetRequestId() *string
}

type AddLiveDomainMappingResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLiveDomainMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveDomainMappingResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveDomainMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveDomainMappingResponseBody) SetRequestId(v string) *AddLiveDomainMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveDomainMappingResponseBody) Validate() error {
	return dara.Validate(s)
}
