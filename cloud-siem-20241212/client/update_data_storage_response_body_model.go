// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataStorageResponseBody
	GetRequestId() *string
}

type UpdateDataStorageResponseBody struct {
	// example:
	//
	// EA7FC160-8D86-5ABE-A08A-7962FDC1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDataStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataStorageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataStorageResponseBody) SetRequestId(v string) *UpdateDataStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
