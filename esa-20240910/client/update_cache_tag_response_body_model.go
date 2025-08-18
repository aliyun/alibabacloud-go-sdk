// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCacheTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCacheTagResponseBody
	GetRequestId() *string
}

type UpdateCacheTagResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-280B-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCacheTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCacheTagResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCacheTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCacheTagResponseBody) SetRequestId(v string) *UpdateCacheTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCacheTagResponseBody) Validate() error {
	return dara.Validate(s)
}
