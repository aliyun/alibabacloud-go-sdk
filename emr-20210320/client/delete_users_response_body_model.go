// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteUsersResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteUsersResponseBody
	GetRequestId() *string
}

type DeleteUsersResponseBody struct {
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUsersResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUsersResponseBody) SetData(v bool) *DeleteUsersResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteUsersResponseBody) SetRequestId(v string) *DeleteUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
