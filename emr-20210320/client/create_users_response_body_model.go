// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CreateUsersResponseBody
	GetData() *bool
	SetRequestId(v string) *CreateUsersResponseBody
	GetRequestId() *string
}

type CreateUsersResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUsersResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUsersResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUsersResponseBody) SetData(v bool) *CreateUsersResponseBody {
	s.Data = &v
	return s
}

func (s *CreateUsersResponseBody) SetRequestId(v string) *CreateUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
