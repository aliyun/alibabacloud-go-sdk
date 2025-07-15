// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEipForwardModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyEipForwardModeResponseBody
	GetRequestId() *string
}

type ModifyEipForwardModeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 62C6A6A5-1534-53D9-AB1E-C9307A147ED5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEipForwardModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyEipForwardModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEipForwardModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyEipForwardModeResponseBody) SetRequestId(v string) *ModifyEipForwardModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyEipForwardModeResponseBody) Validate() error {
	return dara.Validate(s)
}
