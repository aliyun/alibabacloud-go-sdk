// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomImageShareAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddCustomImageShareAccountResponseBody
	GetRequestId() *string
}

type AddCustomImageShareAccountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCustomImageShareAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCustomImageShareAccountResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomImageShareAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCustomImageShareAccountResponseBody) SetRequestId(v string) *AddCustomImageShareAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCustomImageShareAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
