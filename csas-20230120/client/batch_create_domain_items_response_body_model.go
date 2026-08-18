// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateDomainItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchCreateDomainItemsResponseBody
	GetRequestId() *string
}

type BatchCreateDomainItemsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1A4859B9-0DAD-5B40-B603-254445DC6D45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchCreateDomainItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateDomainItemsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateDomainItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCreateDomainItemsResponseBody) SetRequestId(v string) *BatchCreateDomainItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateDomainItemsResponseBody) Validate() error {
	return dara.Validate(s)
}
