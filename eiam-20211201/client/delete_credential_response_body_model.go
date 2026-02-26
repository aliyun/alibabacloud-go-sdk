// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCredentialResponseBody
	GetRequestId() *string
}

type DeleteCredentialResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCredentialResponseBody) SetRequestId(v string) *DeleteCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
