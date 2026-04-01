// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceKeyPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRCInstanceKeyPairResponseBody
	GetRequestId() *string
}

type ModifyRCInstanceKeyPairResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6C36770E-21AE-5689-BAA6-313DA58D5467
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRCInstanceKeyPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceKeyPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRCInstanceKeyPairResponseBody) SetRequestId(v string) *ModifyRCInstanceKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRCInstanceKeyPairResponseBody) Validate() error {
	return dara.Validate(s)
}
