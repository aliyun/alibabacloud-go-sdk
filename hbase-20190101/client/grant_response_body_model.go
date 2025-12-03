// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GrantResponseBody
	GetRequestId() *string
}

type GrantResponseBody struct {
	// example:
	//
	// 9CBF8DF0-4931-4A54-9B60-4C6E1AB5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantResponseBody) GoString() string {
	return s.String()
}

func (s *GrantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantResponseBody) SetRequestId(v string) *GrantResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantResponseBody) Validate() error {
	return dara.Validate(s)
}
