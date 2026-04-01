// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePostgresExtensionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreatePostgresExtensionsResponseBody
	GetRequestId() *string
}

type CreatePostgresExtensionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7E4448A6-9FE6-4474-A0C1-AA7CFC772CAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePostgresExtensionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePostgresExtensionsResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePostgresExtensionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePostgresExtensionsResponseBody) SetRequestId(v string) *CreatePostgresExtensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePostgresExtensionsResponseBody) Validate() error {
	return dara.Validate(s)
}
