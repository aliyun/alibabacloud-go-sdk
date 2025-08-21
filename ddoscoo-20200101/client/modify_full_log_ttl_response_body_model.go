// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFullLogTtlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyFullLogTtlResponseBody
	GetRequestId() *string
}

type ModifyFullLogTtlResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFullLogTtlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyFullLogTtlResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFullLogTtlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyFullLogTtlResponseBody) SetRequestId(v string) *ModifyFullLogTtlResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFullLogTtlResponseBody) Validate() error {
	return dara.Validate(s)
}
