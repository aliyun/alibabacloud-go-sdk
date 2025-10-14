// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateLogStoreResponseBody
	GetRequestId() *string
}

type CreateLogStoreResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLogStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLogStoreResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLogStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLogStoreResponseBody) SetRequestId(v string) *CreateLogStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLogStoreResponseBody) Validate() error {
	return dara.Validate(s)
}
