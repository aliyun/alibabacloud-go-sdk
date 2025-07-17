// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertContactGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsSuccess(v bool) *DeleteAlertContactGroupResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteAlertContactGroupResponseBody
	GetRequestId() *string
}

type DeleteAlertContactGroupResponseBody struct {
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
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C21AB7CF-B7AF-410F-BD61-82D1567F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAlertContactGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactGroupResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteAlertContactGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAlertContactGroupResponseBody) SetIsSuccess(v bool) *DeleteAlertContactGroupResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteAlertContactGroupResponseBody) SetRequestId(v string) *DeleteAlertContactGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlertContactGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
