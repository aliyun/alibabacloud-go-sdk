// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetLiveDomainConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchSetLiveDomainConfigsResponseBody
	GetRequestId() *string
}

type BatchSetLiveDomainConfigsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSetLiveDomainConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchSetLiveDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSetLiveDomainConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchSetLiveDomainConfigsResponseBody) SetRequestId(v string) *BatchSetLiveDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSetLiveDomainConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}
