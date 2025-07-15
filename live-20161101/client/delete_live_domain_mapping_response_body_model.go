// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveDomainMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveDomainMappingResponseBody
	GetRequestId() *string
}

type DeleteLiveDomainMappingResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveDomainMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveDomainMappingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveDomainMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveDomainMappingResponseBody) SetRequestId(v string) *DeleteLiveDomainMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveDomainMappingResponseBody) Validate() error {
	return dara.Validate(s)
}
