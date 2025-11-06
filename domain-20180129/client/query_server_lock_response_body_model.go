// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryServerLockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainInstanceId(v string) *QueryServerLockResponseBody
	GetDomainInstanceId() *string
	SetDomainName(v string) *QueryServerLockResponseBody
	GetDomainName() *string
	SetExpireDate(v string) *QueryServerLockResponseBody
	GetExpireDate() *string
	SetGmtCreate(v string) *QueryServerLockResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *QueryServerLockResponseBody
	GetGmtModified() *string
	SetLockInstanceId(v string) *QueryServerLockResponseBody
	GetLockInstanceId() *string
	SetLockProductId(v string) *QueryServerLockResponseBody
	GetLockProductId() *string
	SetRequestId(v string) *QueryServerLockResponseBody
	GetRequestId() *string
	SetServerLockStatus(v int32) *QueryServerLockResponseBody
	GetServerLockStatus() *int32
	SetStartDate(v string) *QueryServerLockResponseBody
	GetStartDate() *string
	SetUserId(v string) *QueryServerLockResponseBody
	GetUserId() *string
}

type QueryServerLockResponseBody struct {
	// example:
	//
	// S20190N1DAI4****
	DomainInstanceId *string `json:"DomainInstanceId,omitempty" xml:"DomainInstanceId,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 2030-07-10 17:37:36
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// example:
	//
	// 2021-07-10 17:37:36
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2021-07-10 17:37:36
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// S2021591IQ28****
	LockInstanceId *string `json:"LockInstanceId,omitempty" xml:"LockInstanceId,omitempty"`
	// example:
	//
	// 1807**
	LockProductId *string `json:"LockProductId,omitempty" xml:"LockProductId,omitempty"`
	// example:
	//
	// 9DFCF6F8-243C-****-8035-4B12FEFD7D48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	ServerLockStatus *int32 `json:"ServerLockStatus,omitempty" xml:"ServerLockStatus,omitempty"`
	// example:
	//
	// 2021-07-10 17:37:36
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// 121000000****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryServerLockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryServerLockResponseBody) GoString() string {
	return s.String()
}

func (s *QueryServerLockResponseBody) GetDomainInstanceId() *string {
	return s.DomainInstanceId
}

func (s *QueryServerLockResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryServerLockResponseBody) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *QueryServerLockResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryServerLockResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryServerLockResponseBody) GetLockInstanceId() *string {
	return s.LockInstanceId
}

func (s *QueryServerLockResponseBody) GetLockProductId() *string {
	return s.LockProductId
}

func (s *QueryServerLockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryServerLockResponseBody) GetServerLockStatus() *int32 {
	return s.ServerLockStatus
}

func (s *QueryServerLockResponseBody) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryServerLockResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *QueryServerLockResponseBody) SetDomainInstanceId(v string) *QueryServerLockResponseBody {
	s.DomainInstanceId = &v
	return s
}

func (s *QueryServerLockResponseBody) SetDomainName(v string) *QueryServerLockResponseBody {
	s.DomainName = &v
	return s
}

func (s *QueryServerLockResponseBody) SetExpireDate(v string) *QueryServerLockResponseBody {
	s.ExpireDate = &v
	return s
}

func (s *QueryServerLockResponseBody) SetGmtCreate(v string) *QueryServerLockResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *QueryServerLockResponseBody) SetGmtModified(v string) *QueryServerLockResponseBody {
	s.GmtModified = &v
	return s
}

func (s *QueryServerLockResponseBody) SetLockInstanceId(v string) *QueryServerLockResponseBody {
	s.LockInstanceId = &v
	return s
}

func (s *QueryServerLockResponseBody) SetLockProductId(v string) *QueryServerLockResponseBody {
	s.LockProductId = &v
	return s
}

func (s *QueryServerLockResponseBody) SetRequestId(v string) *QueryServerLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryServerLockResponseBody) SetServerLockStatus(v int32) *QueryServerLockResponseBody {
	s.ServerLockStatus = &v
	return s
}

func (s *QueryServerLockResponseBody) SetStartDate(v string) *QueryServerLockResponseBody {
	s.StartDate = &v
	return s
}

func (s *QueryServerLockResponseBody) SetUserId(v string) *QueryServerLockResponseBody {
	s.UserId = &v
	return s
}

func (s *QueryServerLockResponseBody) Validate() error {
	return dara.Validate(s)
}
