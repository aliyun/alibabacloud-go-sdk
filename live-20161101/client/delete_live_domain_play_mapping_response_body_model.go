// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveDomainPlayMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveDomainPlayMappingResponseBody
	GetRequestId() *string
}

type DeleteLiveDomainPlayMappingResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveDomainPlayMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveDomainPlayMappingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveDomainPlayMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveDomainPlayMappingResponseBody) SetRequestId(v string) *DeleteLiveDomainPlayMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveDomainPlayMappingResponseBody) Validate() error {
	return dara.Validate(s)
}
