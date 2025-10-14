// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDataStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetDataStorageResponseBody
	GetRequestId() *string
}

type ResetDataStorageResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetDataStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetDataStorageResponseBody) GoString() string {
	return s.String()
}

func (s *ResetDataStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetDataStorageResponseBody) SetRequestId(v string) *ResetDataStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetDataStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
