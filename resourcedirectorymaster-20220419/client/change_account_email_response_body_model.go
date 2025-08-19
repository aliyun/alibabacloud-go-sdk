// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAccountEmailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeAccountEmailResponseBody
	GetRequestId() *string
}

type ChangeAccountEmailResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeAccountEmailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeAccountEmailResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeAccountEmailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeAccountEmailResponseBody) SetRequestId(v string) *ChangeAccountEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeAccountEmailResponseBody) Validate() error {
	return dara.Validate(s)
}
