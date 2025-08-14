// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CreateEventResponseBody
	GetData() *bool
	SetRequestId(v string) *CreateEventResponseBody
	GetRequestId() *string
}

type CreateEventResponseBody struct {
	// Return result.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEventResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEventResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEventResponseBody) SetData(v bool) *CreateEventResponseBody {
	s.Data = &v
	return s
}

func (s *CreateEventResponseBody) SetRequestId(v string) *CreateEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEventResponseBody) Validate() error {
	return dara.Validate(s)
}
