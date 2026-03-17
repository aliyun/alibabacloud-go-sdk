// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeInstanceFromCbnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeInstanceFromCbnResponseBody
	GetRequestId() *string
}

type RevokeInstanceFromCbnResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 07D73949-2508-4169-8C64-7CCDB33871C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeInstanceFromCbnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeInstanceFromCbnResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromCbnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeInstanceFromCbnResponseBody) SetRequestId(v string) *RevokeInstanceFromCbnResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeInstanceFromCbnResponseBody) Validate() error {
	return dara.Validate(s)
}
