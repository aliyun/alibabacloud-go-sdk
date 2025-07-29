// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactGroupForAlertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMsg(v string) *UpdateContactGroupForAlertResponseBody
	GetMsg() *string
	SetStatus(v bool) *UpdateContactGroupForAlertResponseBody
	GetStatus() *bool
}

type UpdateContactGroupForAlertResponseBody struct {
	// The error message returned if the call fails.
	//
	// example:
	//
	// contact group illegal.
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// The update status.
	//
	// 	- true: The update is successful.
	//
	// 	- false: The update failed.
	//
	// example:
	//
	// true
	Status *bool `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateContactGroupForAlertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactGroupForAlertResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContactGroupForAlertResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *UpdateContactGroupForAlertResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *UpdateContactGroupForAlertResponseBody) SetMsg(v string) *UpdateContactGroupForAlertResponseBody {
	s.Msg = &v
	return s
}

func (s *UpdateContactGroupForAlertResponseBody) SetStatus(v bool) *UpdateContactGroupForAlertResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateContactGroupForAlertResponseBody) Validate() error {
	return dara.Validate(s)
}
