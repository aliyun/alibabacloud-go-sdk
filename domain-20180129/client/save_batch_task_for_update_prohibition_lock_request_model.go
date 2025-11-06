// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForUpdateProhibitionLockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v []*string) *SaveBatchTaskForUpdateProhibitionLockRequest
	GetDomainName() []*string
	SetLang(v string) *SaveBatchTaskForUpdateProhibitionLockRequest
	GetLang() *string
	SetStatus(v bool) *SaveBatchTaskForUpdateProhibitionLockRequest
	GetStatus() *bool
	SetUserClientIp(v string) *SaveBatchTaskForUpdateProhibitionLockRequest
	GetUserClientIp() *string
}

type SaveBatchTaskForUpdateProhibitionLockRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// aliyundoc.com
	DomainName []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveBatchTaskForUpdateProhibitionLockRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForUpdateProhibitionLockRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdateProhibitionLockRequest) GetDomainName() []*string {
	return s.DomainName
}

func (s *SaveBatchTaskForUpdateProhibitionLockRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveBatchTaskForUpdateProhibitionLockRequest) GetStatus() *bool {
	return s.Status
}

func (s *SaveBatchTaskForUpdateProhibitionLockRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveBatchTaskForUpdateProhibitionLockRequest) SetDomainName(v []*string) *SaveBatchTaskForUpdateProhibitionLockRequest {
	s.DomainName = v
	return s
}

func (s *SaveBatchTaskForUpdateProhibitionLockRequest) SetLang(v string) *SaveBatchTaskForUpdateProhibitionLockRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForUpdateProhibitionLockRequest) SetStatus(v bool) *SaveBatchTaskForUpdateProhibitionLockRequest {
	s.Status = &v
	return s
}

func (s *SaveBatchTaskForUpdateProhibitionLockRequest) SetUserClientIp(v string) *SaveBatchTaskForUpdateProhibitionLockRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveBatchTaskForUpdateProhibitionLockRequest) Validate() error {
	return dara.Validate(s)
}
