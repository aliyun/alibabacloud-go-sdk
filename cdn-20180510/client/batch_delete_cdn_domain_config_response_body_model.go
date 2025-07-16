// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteCdnDomainConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchDeleteCdnDomainConfigResponseBody
	GetRequestId() *string
}

type BatchDeleteCdnDomainConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeleteCdnDomainConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteCdnDomainConfigResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteCdnDomainConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteCdnDomainConfigResponseBody) SetRequestId(v string) *BatchDeleteCdnDomainConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteCdnDomainConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
