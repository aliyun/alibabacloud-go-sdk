// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateCdnDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchUpdateCdnDomainResponseBody
	GetRequestId() *string
}

type BatchUpdateCdnDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchUpdateCdnDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateCdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateCdnDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchUpdateCdnDomainResponseBody) SetRequestId(v string) *BatchUpdateCdnDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUpdateCdnDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
