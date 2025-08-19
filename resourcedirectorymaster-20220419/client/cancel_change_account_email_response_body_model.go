// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelChangeAccountEmailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelChangeAccountEmailResponseBody
	GetRequestId() *string
}

type CancelChangeAccountEmailResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelChangeAccountEmailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelChangeAccountEmailResponseBody) GoString() string {
	return s.String()
}

func (s *CancelChangeAccountEmailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelChangeAccountEmailResponseBody) SetRequestId(v string) *CancelChangeAccountEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelChangeAccountEmailResponseBody) Validate() error {
	return dara.Validate(s)
}
