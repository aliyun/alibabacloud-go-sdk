// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateInstanceSettingsResponseBody
	GetRequestId() *string
}

type UpdateInstanceSettingsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// BB1C321A-211C-4FD7-BD8B-7F2FABE2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceSettingsResponseBody) SetRequestId(v string) *UpdateInstanceSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceSettingsResponseBody) Validate() error {
	return dara.Validate(s)
}
