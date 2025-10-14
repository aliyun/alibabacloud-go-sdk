// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateLogStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ValidateLogStoreResponseBody
	GetRequestId() *string
}

type ValidateLogStoreResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ValidateLogStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateLogStoreResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateLogStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateLogStoreResponseBody) SetRequestId(v string) *ValidateLogStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateLogStoreResponseBody) Validate() error {
	return dara.Validate(s)
}
