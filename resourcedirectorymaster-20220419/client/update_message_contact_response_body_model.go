// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMessageContactResponseBody
	GetRequestId() *string
}

type UpdateMessageContactResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMessageContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageContactResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMessageContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMessageContactResponseBody) SetRequestId(v string) *UpdateMessageContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMessageContactResponseBody) Validate() error {
	return dara.Validate(s)
}
