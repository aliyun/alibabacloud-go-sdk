// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAlertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsSuccess(v bool) *StartAlertResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *StartAlertResponseBody
	GetRequestId() *string
}

type StartAlertResponseBody struct {
	// Indicates whether the call was successful.
	//
	// 	- `true`: The call was successful.
	//
	// 	- `false`: The call failed.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 27E653FA-5958-45BE-8AA9-14D884DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartAlertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartAlertResponseBody) GoString() string {
	return s.String()
}

func (s *StartAlertResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *StartAlertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartAlertResponseBody) SetIsSuccess(v bool) *StartAlertResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *StartAlertResponseBody) SetRequestId(v string) *StartAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAlertResponseBody) Validate() error {
	return dara.Validate(s)
}
