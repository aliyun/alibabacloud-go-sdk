// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddCdnDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchAddCdnDomainResponseBody
	GetRequestId() *string
}

type BatchAddCdnDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchAddCdnDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchAddCdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddCdnDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchAddCdnDomainResponseBody) SetRequestId(v string) *BatchAddCdnDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchAddCdnDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
