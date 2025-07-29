// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAlertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMsg(v string) *StartAlertResponseBody
	GetMsg() *string
	SetStatus(v bool) *StartAlertResponseBody
	GetStatus() *bool
}

type StartAlertResponseBody struct {
	// The message returned.
	//
	// example:
	//
	// success
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// The status.
	//
	// example:
	//
	// true
	Status *bool `json:"status,omitempty" xml:"status,omitempty"`
}

func (s StartAlertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartAlertResponseBody) GoString() string {
	return s.String()
}

func (s *StartAlertResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *StartAlertResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *StartAlertResponseBody) SetMsg(v string) *StartAlertResponseBody {
	s.Msg = &v
	return s
}

func (s *StartAlertResponseBody) SetStatus(v bool) *StartAlertResponseBody {
	s.Status = &v
	return s
}

func (s *StartAlertResponseBody) Validate() error {
	return dara.Validate(s)
}
