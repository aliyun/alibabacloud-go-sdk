// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyInfoInRecycleBinRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserAccessKeyId(v string) *GetAccessKeyInfoInRecycleBinRequest
	GetUserAccessKeyId() *string
}

type GetAccessKeyInfoInRecycleBinRequest struct {
	// The AccessKey ID of the Resource Access Management (RAM) user.
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAI*******************
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
}

func (s GetAccessKeyInfoInRecycleBinRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyInfoInRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *GetAccessKeyInfoInRecycleBinRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *GetAccessKeyInfoInRecycleBinRequest) SetUserAccessKeyId(v string) *GetAccessKeyInfoInRecycleBinRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinRequest) Validate() error {
	return dara.Validate(s)
}
