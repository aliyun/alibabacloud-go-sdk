// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTableMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTableMetaResponseBody
	GetRequestId() *string
}

type DeleteTableMetaResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTableMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTableMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTableMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTableMetaResponseBody) SetRequestId(v string) *DeleteTableMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTableMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
