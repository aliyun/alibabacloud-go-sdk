// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSummaryInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSummaryInfoRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeSummaryInfoRequest
	GetSourceIp() *string
}

type DescribeSummaryInfoRequest struct {
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
	// The source IP address of the request.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeSummaryInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSummaryInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeSummaryInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSummaryInfoRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSummaryInfoRequest) SetLang(v string) *DescribeSummaryInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSummaryInfoRequest) SetSourceIp(v string) *DescribeSummaryInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSummaryInfoRequest) Validate() error {
	return dara.Validate(s)
}
