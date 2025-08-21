// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetVsDomainConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchSetVsDomainConfigsResponseBody
	GetRequestId() *string
}

type BatchSetVsDomainConfigsResponseBody struct {
	// example:
	//
	// 9BEC5E85-C76B-56EF-A922-860EFDB8B64B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSetVsDomainConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchSetVsDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSetVsDomainConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchSetVsDomainConfigsResponseBody) SetRequestId(v string) *BatchSetVsDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSetVsDomainConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}
