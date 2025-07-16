// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCdnDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCdnDomainResponseBody
	GetRequestId() *string
}

type DeleteCdnDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCdnDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCdnDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCdnDomainResponseBody) SetRequestId(v string) *DeleteCdnDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCdnDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
