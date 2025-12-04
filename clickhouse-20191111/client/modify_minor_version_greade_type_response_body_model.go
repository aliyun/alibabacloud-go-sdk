// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMinorVersionGreadeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyMinorVersionGreadeTypeResponseBody
	GetRequestId() *string
}

type ModifyMinorVersionGreadeTypeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMinorVersionGreadeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMinorVersionGreadeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMinorVersionGreadeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMinorVersionGreadeTypeResponseBody) SetRequestId(v string) *ModifyMinorVersionGreadeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMinorVersionGreadeTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
