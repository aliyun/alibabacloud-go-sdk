// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUIAccountPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ModifyUIAccountPasswordRequest
	GetAccountName() *string
	SetAccountPassword(v string) *ModifyUIAccountPasswordRequest
	GetAccountPassword() *string
	SetClusterId(v string) *ModifyUIAccountPasswordRequest
	GetClusterId() *string
}

type ModifyUIAccountPasswordRequest struct {
	// The username of the cluster management system to be reset. If no user has been created for the HBase instance, the default username is **root**.
	//
	// This parameter is required.
	//
	// example:
	//
	// test01
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The new password for the target username of the cluster management system. The password must be 2 to 30 characters in length and must contain one or more of the following character types: uppercase letters, lowercase letters, special characters, and digits. Supported special characters are underscores (_) and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// **********
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The ID of the target instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/144595.html) operation to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ModifyUIAccountPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUIAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyUIAccountPasswordRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyUIAccountPasswordRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *ModifyUIAccountPasswordRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyUIAccountPasswordRequest) SetAccountName(v string) *ModifyUIAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyUIAccountPasswordRequest) SetAccountPassword(v string) *ModifyUIAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ModifyUIAccountPasswordRequest) SetClusterId(v string) *ModifyUIAccountPasswordRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyUIAccountPasswordRequest) Validate() error {
	return dara.Validate(s)
}
