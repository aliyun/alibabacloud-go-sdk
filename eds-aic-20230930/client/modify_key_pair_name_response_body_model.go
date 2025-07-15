// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyKeyPairNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyKeyPairNameResponseBody
	GetRequestId() *string
}

type ModifyKeyPairNameResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyKeyPairNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyKeyPairNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyKeyPairNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyKeyPairNameResponseBody) SetRequestId(v string) *ModifyKeyPairNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyKeyPairNameResponseBody) Validate() error {
	return dara.Validate(s)
}
