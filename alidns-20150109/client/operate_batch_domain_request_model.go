// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateBatchDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainRecordInfo(v []*OperateBatchDomainRequestDomainRecordInfo) *OperateBatchDomainRequest
	GetDomainRecordInfo() []*OperateBatchDomainRequestDomainRecordInfo
	SetLang(v string) *OperateBatchDomainRequest
	GetLang() *string
	SetType(v string) *OperateBatchDomainRequest
	GetType() *string
}

type OperateBatchDomainRequest struct {
	// The DNS records. You can submit up to 1,000 DNS records.
	//
	// This parameter is required.
	DomainRecordInfo []*OperateBatchDomainRequestDomainRecordInfo `json:"DomainRecordInfo,omitempty" xml:"DomainRecordInfo,omitempty" type:"Repeated"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: zh
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of the batch operation. Valid values:
	//
	// 	- **DOMAIN_ADD**: adds domain names in batches.
	//
	// 	- **DOMAIN_DEL**: deletes domain names in batches.
	//
	// 	- **RR_ADD**: adds DNS records in batches.
	//
	// 	- **RR_DEL**: deletes DNS records in batches. This operation deletes the DNS records with the specified hostname or record value. If you do not specify the Rr and Value parameters, this operation deletes the DNS records that are added for the specified domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// RR_ADD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OperateBatchDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateBatchDomainRequest) GoString() string {
	return s.String()
}

func (s *OperateBatchDomainRequest) GetDomainRecordInfo() []*OperateBatchDomainRequestDomainRecordInfo {
	return s.DomainRecordInfo
}

func (s *OperateBatchDomainRequest) GetLang() *string {
	return s.Lang
}

func (s *OperateBatchDomainRequest) GetType() *string {
	return s.Type
}

func (s *OperateBatchDomainRequest) SetDomainRecordInfo(v []*OperateBatchDomainRequestDomainRecordInfo) *OperateBatchDomainRequest {
	s.DomainRecordInfo = v
	return s
}

func (s *OperateBatchDomainRequest) SetLang(v string) *OperateBatchDomainRequest {
	s.Lang = &v
	return s
}

func (s *OperateBatchDomainRequest) SetType(v string) *OperateBatchDomainRequest {
	s.Type = &v
	return s
}

func (s *OperateBatchDomainRequest) Validate() error {
	if s.DomainRecordInfo != nil {
		for _, item := range s.DomainRecordInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OperateBatchDomainRequestDomainRecordInfo struct {
	// The domain name.
	//
	// >  You can submit 1 to 1,000 domain names. Due to the limit on the length of HTTP request headers, excessive domain names are ignored. Do not enter more than 1,000 domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The DNS request source. Default value: default.
	//
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The new hostname (used only for modification operations, not for external users).
	//
	// example:
	//
	// mail
	NewRr *string `json:"NewRr,omitempty" xml:"NewRr,omitempty"`
	// The new type of the DNS record (used only for modification operations, not for external users).
	//
	// example:
	//
	// AAAA
	NewType *string `json:"NewType,omitempty" xml:"NewType,omitempty"`
	// The new value of the DNS record (used only for modification operations, not for external users).
	//
	// example:
	//
	// 114.92.XX.XX
	NewValue *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	// The priority of the mail exchanger (MX) record.
	//
	// This parameter is required if the type of the DNS record is MX. Default value: 10.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The hostname.
	//
	// >  This parameter is required if you set Type to **RR_ADD*	- or **RR_DEL**.
	//
	// example:
	//
	// zhaohui
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The time-to-live (TTL) value of the cached DNS record. Unit: seconds. Default value: ***600***.
	//
	// example:
	//
	// 600
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type of the DNS record. Valid values: A, AAAA, TXT, MX, and CNAME.
	//
	// >  This parameter is required if you set Type to **RR_ADD*	- or **RR_DEL**.
	//
	// example:
	//
	// MX
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the DNS record.
	//
	// >  This parameter is required if you set Type to **RR_ADD*	- or **RR_DEL**.
	//
	// example:
	//
	// fd87da3c4528844d45af39200155a905
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OperateBatchDomainRequestDomainRecordInfo) String() string {
	return dara.Prettify(s)
}

func (s OperateBatchDomainRequestDomainRecordInfo) GoString() string {
	return s.String()
}

func (s *OperateBatchDomainRequestDomainRecordInfo) GetDomain() *string {
	return s.Domain
}

func (s *OperateBatchDomainRequestDomainRecordInfo) GetLine() *string {
	return s.Line
}

func (s *OperateBatchDomainRequestDomainRecordInfo) GetNewRr() *string {
	return s.NewRr
}

func (s *OperateBatchDomainRequestDomainRecordInfo) GetNewType() *string {
	return s.NewType
}

func (s *OperateBatchDomainRequestDomainRecordInfo) GetNewValue() *string {
	return s.NewValue
}

func (s *OperateBatchDomainRequestDomainRecordInfo) GetPriority() *int32 {
	return s.Priority
}

func (s *OperateBatchDomainRequestDomainRecordInfo) GetRr() *string {
	return s.Rr
}

func (s *OperateBatchDomainRequestDomainRecordInfo) GetTtl() *int32 {
	return s.Ttl
}

func (s *OperateBatchDomainRequestDomainRecordInfo) GetType() *string {
	return s.Type
}

func (s *OperateBatchDomainRequestDomainRecordInfo) GetValue() *string {
	return s.Value
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetDomain(v string) *OperateBatchDomainRequestDomainRecordInfo {
	s.Domain = &v
	return s
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetLine(v string) *OperateBatchDomainRequestDomainRecordInfo {
	s.Line = &v
	return s
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetNewRr(v string) *OperateBatchDomainRequestDomainRecordInfo {
	s.NewRr = &v
	return s
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetNewType(v string) *OperateBatchDomainRequestDomainRecordInfo {
	s.NewType = &v
	return s
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetNewValue(v string) *OperateBatchDomainRequestDomainRecordInfo {
	s.NewValue = &v
	return s
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetPriority(v int32) *OperateBatchDomainRequestDomainRecordInfo {
	s.Priority = &v
	return s
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetRr(v string) *OperateBatchDomainRequestDomainRecordInfo {
	s.Rr = &v
	return s
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetTtl(v int32) *OperateBatchDomainRequestDomainRecordInfo {
	s.Ttl = &v
	return s
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetType(v string) *OperateBatchDomainRequestDomainRecordInfo {
	s.Type = &v
	return s
}

func (s *OperateBatchDomainRequestDomainRecordInfo) SetValue(v string) *OperateBatchDomainRequestDomainRecordInfo {
	s.Value = &v
	return s
}

func (s *OperateBatchDomainRequestDomainRecordInfo) Validate() error {
	return dara.Validate(s)
}
