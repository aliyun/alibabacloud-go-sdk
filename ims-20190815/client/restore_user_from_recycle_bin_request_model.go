// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreUserFromRecycleBinRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *RestoreUserFromRecycleBinRequest
	GetUserId() *string
}

type RestoreUserFromRecycleBinRequest struct {
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RestoreUserFromRecycleBinRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreUserFromRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *RestoreUserFromRecycleBinRequest) GetUserId() *string {
	return s.UserId
}

func (s *RestoreUserFromRecycleBinRequest) SetUserId(v string) *RestoreUserFromRecycleBinRequest {
	s.UserId = &v
	return s
}

func (s *RestoreUserFromRecycleBinRequest) Validate() error {
	return dara.Validate(s)
}
