// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GrantApiKeyResponseBody
	GetRequestId() *string
}

type GrantApiKeyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GrantApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantApiKeyResponseBody) SetRequestId(v string) *GrantApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
