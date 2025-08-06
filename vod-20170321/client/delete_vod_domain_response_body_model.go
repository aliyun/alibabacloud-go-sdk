// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVodDomainResponseBody
	GetRequestId() *string
}

type DeleteVodDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-****-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVodDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVodDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVodDomainResponseBody) SetRequestId(v string) *DeleteVodDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVodDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
