// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTableMetaResponseBody
	GetRequestId() *string
}

type UpdateTableMetaResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTableMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableMetaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTableMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTableMetaResponseBody) SetRequestId(v string) *UpdateTableMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTableMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
