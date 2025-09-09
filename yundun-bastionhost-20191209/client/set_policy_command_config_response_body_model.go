// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyCommandConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetPolicyCommandConfigResponseBody
	GetRequestId() *string
}

type SetPolicyCommandConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4F6C075F-FC86-476E-943B-097BD4E12948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPolicyCommandConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyCommandConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetPolicyCommandConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetPolicyCommandConfigResponseBody) SetRequestId(v string) *SetPolicyCommandConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetPolicyCommandConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
