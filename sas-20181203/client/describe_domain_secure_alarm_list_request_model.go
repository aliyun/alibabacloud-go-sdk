// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecureAlarmListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *DescribeDomainSecureAlarmListRequest
	GetFrom() *string
	SetLang(v string) *DescribeDomainSecureAlarmListRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeDomainSecureAlarmListRequest
	GetSourceIp() *string
}

type DescribeDomainSecureAlarmListRequest struct {
	// The identifier of the request source. Set the value to sas.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 139.227.*.*
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDomainSecureAlarmListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureAlarmListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureAlarmListRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeDomainSecureAlarmListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDomainSecureAlarmListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDomainSecureAlarmListRequest) SetFrom(v string) *DescribeDomainSecureAlarmListRequest {
	s.From = &v
	return s
}

func (s *DescribeDomainSecureAlarmListRequest) SetLang(v string) *DescribeDomainSecureAlarmListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainSecureAlarmListRequest) SetSourceIp(v string) *DescribeDomainSecureAlarmListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainSecureAlarmListRequest) Validate() error {
	return dara.Validate(s)
}
