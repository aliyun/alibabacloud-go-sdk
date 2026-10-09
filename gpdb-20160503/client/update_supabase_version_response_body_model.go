// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSupabaseVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSupabaseVersionResponseBody
	GetRequestId() *string
}

type UpdateSupabaseVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSupabaseVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSupabaseVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSupabaseVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSupabaseVersionResponseBody) SetRequestId(v string) *UpdateSupabaseVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSupabaseVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
