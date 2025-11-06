// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryServerLockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryServerLockRequest
	GetInstanceId() *string
	SetLang(v string) *QueryServerLockRequest
	GetLang() *string
	SetUserClientIp(v string) *QueryServerLockRequest
	GetUserClientIp() *string
}

type QueryServerLockRequest struct {
	// example:
	//
	// S20181*****85212
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryServerLockRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryServerLockRequest) GoString() string {
	return s.String()
}

func (s *QueryServerLockRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryServerLockRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryServerLockRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryServerLockRequest) SetInstanceId(v string) *QueryServerLockRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryServerLockRequest) SetLang(v string) *QueryServerLockRequest {
	s.Lang = &v
	return s
}

func (s *QueryServerLockRequest) SetUserClientIp(v string) *QueryServerLockRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryServerLockRequest) Validate() error {
	return dara.Validate(s)
}
