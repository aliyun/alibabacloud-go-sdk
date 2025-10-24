// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterCustomViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomViewId(v string) *RegisterCustomViewResponseBody
	GetCustomViewId() *string
	SetRequestId(v string) *RegisterCustomViewResponseBody
	GetRequestId() *string
}

type RegisterCustomViewResponseBody struct {
	// example:
	//
	// 1
	CustomViewId *string `json:"CustomViewId,omitempty" xml:"CustomViewId,omitempty"`
	// example:
	//
	// 580e8ce3-3b80-44c5-9f3f-36ac3cc5bdd5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterCustomViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterCustomViewResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterCustomViewResponseBody) GetCustomViewId() *string {
	return s.CustomViewId
}

func (s *RegisterCustomViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterCustomViewResponseBody) SetCustomViewId(v string) *RegisterCustomViewResponseBody {
	s.CustomViewId = &v
	return s
}

func (s *RegisterCustomViewResponseBody) SetRequestId(v string) *RegisterCustomViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterCustomViewResponseBody) Validate() error {
	return dara.Validate(s)
}
