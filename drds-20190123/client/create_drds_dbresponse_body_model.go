// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDrdsDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDrdsDBResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDrdsDBResponseBody
	GetSuccess() *bool
}

type CreateDrdsDBResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// FF13E47D-4E38-4A5A-BA68-32A554AD45T6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDrdsDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDrdsDBResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDrdsDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDrdsDBResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDrdsDBResponseBody) SetRequestId(v string) *CreateDrdsDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDrdsDBResponseBody) SetSuccess(v bool) *CreateDrdsDBResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDrdsDBResponseBody) Validate() error {
	return dara.Validate(s)
}
