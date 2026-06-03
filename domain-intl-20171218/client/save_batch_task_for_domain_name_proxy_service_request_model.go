// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForDomainNameProxyServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v []*string) *SaveBatchTaskForDomainNameProxyServiceRequest
	GetDomainName() []*string
	SetLang(v string) *SaveBatchTaskForDomainNameProxyServiceRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveBatchTaskForDomainNameProxyServiceRequest
	GetUserClientIp() *string
	SetStatus(v bool) *SaveBatchTaskForDomainNameProxyServiceRequest
	GetStatus() *bool
}

type SaveBatchTaskForDomainNameProxyServiceRequest struct {
	// This parameter is required.
	DomainName   []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	Lang         *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// This parameter is required.
	Status *bool `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SaveBatchTaskForDomainNameProxyServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForDomainNameProxyServiceRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) GetDomainName() []*string {
	return s.DomainName
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) GetStatus() *bool {
	return s.Status
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) SetDomainName(v []*string) *SaveBatchTaskForDomainNameProxyServiceRequest {
	s.DomainName = v
	return s
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) SetLang(v string) *SaveBatchTaskForDomainNameProxyServiceRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) SetUserClientIp(v string) *SaveBatchTaskForDomainNameProxyServiceRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) SetStatus(v bool) *SaveBatchTaskForDomainNameProxyServiceRequest {
	s.Status = &v
	return s
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) Validate() error {
	return dara.Validate(s)
}
