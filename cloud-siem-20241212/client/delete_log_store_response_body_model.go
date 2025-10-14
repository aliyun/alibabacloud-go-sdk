// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLogStoreResponseBody
	GetRequestId() *string
}

type DeleteLogStoreResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLogStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogStoreResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLogStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLogStoreResponseBody) SetRequestId(v string) *DeleteLogStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLogStoreResponseBody) Validate() error {
	return dara.Validate(s)
}
