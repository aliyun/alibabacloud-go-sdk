// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeUserLangResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeUserLangResponseBody
	GetRequestId() *string
}

type ChangeUserLangResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2F5AA940-9EBF-5948-ACE7-3EF0FE54****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeUserLangResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeUserLangResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeUserLangResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeUserLangResponseBody) SetRequestId(v string) *ChangeUserLangResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeUserLangResponseBody) Validate() error {
	return dara.Validate(s)
}
