// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchDomainRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *SaveBatchDomainRemarkRequest
	GetInstanceIds() *string
	SetLang(v string) *SaveBatchDomainRemarkRequest
	GetLang() *string
	SetRemark(v string) *SaveBatchDomainRemarkRequest
	GetRemark() *string
	SetUserClientIp(v string) *SaveBatchDomainRemarkRequest
	GetUserClientIp() *string
}

type SaveBatchDomainRemarkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// S12344567
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// MyRemarkInfo
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveBatchDomainRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchDomainRemarkRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchDomainRemarkRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *SaveBatchDomainRemarkRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveBatchDomainRemarkRequest) GetRemark() *string {
	return s.Remark
}

func (s *SaveBatchDomainRemarkRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveBatchDomainRemarkRequest) SetInstanceIds(v string) *SaveBatchDomainRemarkRequest {
	s.InstanceIds = &v
	return s
}

func (s *SaveBatchDomainRemarkRequest) SetLang(v string) *SaveBatchDomainRemarkRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchDomainRemarkRequest) SetRemark(v string) *SaveBatchDomainRemarkRequest {
	s.Remark = &v
	return s
}

func (s *SaveBatchDomainRemarkRequest) SetUserClientIp(v string) *SaveBatchDomainRemarkRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveBatchDomainRemarkRequest) Validate() error {
	return dara.Validate(s)
}
