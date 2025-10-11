// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFilterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateFilterResponseBody
	GetRequestId() *string
}

type CreateFilterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EEF1EE1F-50F6-5494-B3DA-8F597DEB31BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFilterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFilterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFilterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFilterResponseBody) SetRequestId(v string) *CreateFilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFilterResponseBody) Validate() error {
	return dara.Validate(s)
}
