// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCacheReserveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCacheReserveResponseBody
	GetRequestId() *string
}

type UpdateCacheReserveResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EDBD3EB3-97DA-5465-AEF5-8DCA5DC5E395
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCacheReserveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCacheReserveResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCacheReserveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCacheReserveResponseBody) SetRequestId(v string) *UpdateCacheReserveResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCacheReserveResponseBody) Validate() error {
	return dara.Validate(s)
}
