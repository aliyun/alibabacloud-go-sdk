// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserConfigResponseBody
	GetRequestId() *string
	SetUser(v *GetUserConfigResponseBodyUser) *GetUserConfigResponseBody
	GetUser() *GetUserConfigResponseBodyUser
}

type GetUserConfigResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	User      *GetUserConfigResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s GetUserConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserConfigResponseBody) GetUser() *GetUserConfigResponseBodyUser {
	return s.User
}

func (s *GetUserConfigResponseBody) SetRequestId(v string) *GetUserConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserConfigResponseBody) SetUser(v *GetUserConfigResponseBodyUser) *GetUserConfigResponseBody {
	s.User = v
	return s
}

func (s *GetUserConfigResponseBody) Validate() error {
	if s.User != nil {
		if err := s.User.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserConfigResponseBodyUser struct {
	// example:
	//
	// v2
	CtdrVersion *string `json:"CtdrVersion,omitempty" xml:"CtdrVersion,omitempty"`
	// example:
	//
	// pending
	DataStorageVersion *string `json:"DataStorageVersion,omitempty" xml:"DataStorageVersion,omitempty"`
	// example:
	//
	// v2
	UpgradeCtdrVersion *string `json:"UpgradeCtdrVersion,omitempty" xml:"UpgradeCtdrVersion,omitempty"`
	// example:
	//
	// v2
	UpgradeStatus *string `json:"UpgradeStatus,omitempty" xml:"UpgradeStatus,omitempty"`
}

func (s GetUserConfigResponseBodyUser) String() string {
	return dara.Prettify(s)
}

func (s GetUserConfigResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetUserConfigResponseBodyUser) GetCtdrVersion() *string {
	return s.CtdrVersion
}

func (s *GetUserConfigResponseBodyUser) GetDataStorageVersion() *string {
	return s.DataStorageVersion
}

func (s *GetUserConfigResponseBodyUser) GetUpgradeCtdrVersion() *string {
	return s.UpgradeCtdrVersion
}

func (s *GetUserConfigResponseBodyUser) GetUpgradeStatus() *string {
	return s.UpgradeStatus
}

func (s *GetUserConfigResponseBodyUser) SetCtdrVersion(v string) *GetUserConfigResponseBodyUser {
	s.CtdrVersion = &v
	return s
}

func (s *GetUserConfigResponseBodyUser) SetDataStorageVersion(v string) *GetUserConfigResponseBodyUser {
	s.DataStorageVersion = &v
	return s
}

func (s *GetUserConfigResponseBodyUser) SetUpgradeCtdrVersion(v string) *GetUserConfigResponseBodyUser {
	s.UpgradeCtdrVersion = &v
	return s
}

func (s *GetUserConfigResponseBodyUser) SetUpgradeStatus(v string) *GetUserConfigResponseBodyUser {
	s.UpgradeStatus = &v
	return s
}

func (s *GetUserConfigResponseBodyUser) Validate() error {
	return dara.Validate(s)
}
