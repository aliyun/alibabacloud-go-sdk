// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApplicationAccountFromUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveApplicationAccountFromUserResponseBody
	GetRequestId() *string
}

type RemoveApplicationAccountFromUserResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveApplicationAccountFromUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveApplicationAccountFromUserResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveApplicationAccountFromUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveApplicationAccountFromUserResponseBody) SetRequestId(v string) *RemoveApplicationAccountFromUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveApplicationAccountFromUserResponseBody) Validate() error {
	return dara.Validate(s)
}
