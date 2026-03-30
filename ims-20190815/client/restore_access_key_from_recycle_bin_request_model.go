// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreAccessKeyFromRecycleBinRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserAccessKeyId(v string) *RestoreAccessKeyFromRecycleBinRequest
	GetUserAccessKeyId() *string
	SetUserId(v string) *RestoreAccessKeyFromRecycleBinRequest
	GetUserId() *string
}

type RestoreAccessKeyFromRecycleBinRequest struct {
	// The AccessKey ID of the RAM user.
	//
	// example:
	//
	// LTAI*******************
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// The ID of the Resource Access Management (RAM) user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RestoreAccessKeyFromRecycleBinRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreAccessKeyFromRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *RestoreAccessKeyFromRecycleBinRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *RestoreAccessKeyFromRecycleBinRequest) GetUserId() *string {
	return s.UserId
}

func (s *RestoreAccessKeyFromRecycleBinRequest) SetUserAccessKeyId(v string) *RestoreAccessKeyFromRecycleBinRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *RestoreAccessKeyFromRecycleBinRequest) SetUserId(v string) *RestoreAccessKeyFromRecycleBinRequest {
	s.UserId = &v
	return s
}

func (s *RestoreAccessKeyFromRecycleBinRequest) Validate() error {
	return dara.Validate(s)
}
