// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDomainItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchDeleteDomainItemsResponseBody
	GetRequestId() *string
}

type BatchDeleteDomainItemsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 0C76C4AD-5E46-555D-981B-CB004C37F41A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeleteDomainItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDomainItemsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteDomainItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteDomainItemsResponseBody) SetRequestId(v string) *BatchDeleteDomainItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteDomainItemsResponseBody) Validate() error {
	return dara.Validate(s)
}
