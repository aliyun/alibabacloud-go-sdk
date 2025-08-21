// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteVsDomainConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchDeleteVsDomainConfigsResponseBody
	GetRequestId() *string
}

type BatchDeleteVsDomainConfigsResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeleteVsDomainConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteVsDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteVsDomainConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteVsDomainConfigsResponseBody) SetRequestId(v string) *BatchDeleteVsDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteVsDomainConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}
