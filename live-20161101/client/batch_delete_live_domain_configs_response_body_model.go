// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteLiveDomainConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchDeleteLiveDomainConfigsResponseBody
	GetRequestId() *string
}

type BatchDeleteLiveDomainConfigsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeleteLiveDomainConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteLiveDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteLiveDomainConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteLiveDomainConfigsResponseBody) SetRequestId(v string) *BatchDeleteLiveDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteLiveDomainConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}
