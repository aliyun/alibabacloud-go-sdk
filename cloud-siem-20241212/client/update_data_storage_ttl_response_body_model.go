// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataStorageTtlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataStorageTtlResponseBody
	GetRequestId() *string
}

type UpdateDataStorageTtlResponseBody struct {
	// example:
	//
	// D92E4FCF-4584-5E50-9C02-26B79A9C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDataStorageTtlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataStorageTtlResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataStorageTtlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataStorageTtlResponseBody) SetRequestId(v string) *UpdateDataStorageTtlResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataStorageTtlResponseBody) Validate() error {
	return dara.Validate(s)
}
