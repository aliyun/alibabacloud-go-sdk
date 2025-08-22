// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDcdnDomainConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchDeleteDcdnDomainConfigsResponseBody
	GetRequestId() *string
}

type BatchDeleteDcdnDomainConfigsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeleteDcdnDomainConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDcdnDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteDcdnDomainConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteDcdnDomainConfigsResponseBody) SetRequestId(v string) *BatchDeleteDcdnDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteDcdnDomainConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}
