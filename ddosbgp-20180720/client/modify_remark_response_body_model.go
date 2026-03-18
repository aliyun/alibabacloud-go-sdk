// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRemarkResponseBody
	GetRequestId() *string
}

type ModifyRemarkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6AC3597B-7FD5-5E68-97C3-E11F4D010732
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRemarkResponseBody) SetRequestId(v string) *ModifyRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
