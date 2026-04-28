// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDefaultDriveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *GetDefaultDriveRequest
	GetUserId() *string
}

type GetDefaultDriveRequest struct {
	// The user ID. If you use an AccessKey pair for authentication, you must specify this parameter. If you use an access token for authentication, this parameter is optional. By default, the user ID associated with the access token is used.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s GetDefaultDriveRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDefaultDriveRequest) GoString() string {
	return s.String()
}

func (s *GetDefaultDriveRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetDefaultDriveRequest) SetUserId(v string) *GetDefaultDriveRequest {
	s.UserId = &v
	return s
}

func (s *GetDefaultDriveRequest) Validate() error {
	return dara.Validate(s)
}
