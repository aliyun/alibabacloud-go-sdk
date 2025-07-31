// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAccountPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetAccountPasswordResponseBody
	GetRequestId() *string
}

type ResetAccountPasswordResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 06CBD06E-ABC9-4121-AB93-3C3820B3E7E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAccountPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetAccountPasswordResponseBody) SetRequestId(v string) *ResetAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAccountPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
