// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateStackResponseBody
	GetRequestId() *string
}

type UpdateStackResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 0D298375-F92F-5B65-82E4-EA68F02521F1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateStackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateStackResponseBody) SetRequestId(v string) *UpdateStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStackResponseBody) Validate() error {
	return dara.Validate(s)
}
