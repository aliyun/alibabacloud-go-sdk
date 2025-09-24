// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRunResponseBody
	GetRequestId() *string
}

type UpdateRunResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ADF6D849-*****-7E7030F0CE53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRunResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRunResponseBody) SetRequestId(v string) *UpdateRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRunResponseBody) Validate() error {
	return dara.Validate(s)
}
