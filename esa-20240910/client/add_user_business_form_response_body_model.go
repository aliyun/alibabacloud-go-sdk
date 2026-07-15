// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserBusinessFormResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddUserBusinessFormResponseBody
	GetRequestId() *string
}

type AddUserBusinessFormResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserBusinessFormResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUserBusinessFormResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserBusinessFormResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUserBusinessFormResponseBody) SetRequestId(v string) *AddUserBusinessFormResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserBusinessFormResponseBody) Validate() error {
	return dara.Validate(s)
}
