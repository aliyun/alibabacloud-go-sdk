// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForApplyQuickTransferOutOpenlyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v []*string) *SaveBatchTaskForApplyQuickTransferOutOpenlyRequest
	GetDomainNames() []*string
	SetLang(v string) *SaveBatchTaskForApplyQuickTransferOutOpenlyRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveBatchTaskForApplyQuickTransferOutOpenlyRequest
	GetUserClientIp() *string
}

type SaveBatchTaskForApplyQuickTransferOutOpenlyRequest struct {
	DomainNames []*string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty" type:"Repeated"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveBatchTaskForApplyQuickTransferOutOpenlyRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForApplyQuickTransferOutOpenlyRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForApplyQuickTransferOutOpenlyRequest) GetDomainNames() []*string {
	return s.DomainNames
}

func (s *SaveBatchTaskForApplyQuickTransferOutOpenlyRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveBatchTaskForApplyQuickTransferOutOpenlyRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveBatchTaskForApplyQuickTransferOutOpenlyRequest) SetDomainNames(v []*string) *SaveBatchTaskForApplyQuickTransferOutOpenlyRequest {
	s.DomainNames = v
	return s
}

func (s *SaveBatchTaskForApplyQuickTransferOutOpenlyRequest) SetLang(v string) *SaveBatchTaskForApplyQuickTransferOutOpenlyRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForApplyQuickTransferOutOpenlyRequest) SetUserClientIp(v string) *SaveBatchTaskForApplyQuickTransferOutOpenlyRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveBatchTaskForApplyQuickTransferOutOpenlyRequest) Validate() error {
	return dara.Validate(s)
}
