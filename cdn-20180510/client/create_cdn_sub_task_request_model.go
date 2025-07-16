// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCdnSubTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *CreateCdnSubTaskRequest
	GetDomainName() *string
	SetReportIds(v string) *CreateCdnSubTaskRequest
	GetReportIds() *string
}

type CreateCdnSubTaskRequest struct {
	// The domain names to be tracked. Separate multiple domain names with commas (,). You can specify up to 500 domain names. If you want to specify more than 500 domain names, [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.2020520001.aliyun_topbar.18.dbd44bd3e4f845#/ticket/createIndex).
	//
	// > If you do not specify a domain name, the custom operations report is created for all domain names that belong to your Alibaba Cloud account.
	//
	// example:
	//
	// www.example1.com,www.example2.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The IDs of the metrics that you want to include in the report. Separate multiple IDs with commas (,). Valid values:
	//
	// 	- **1**: frequently requested URLs (ranked by the number of requests)
	//
	// 	- **3**: frequently requested URLs (ranked by the amount of network traffic)
	//
	// 	- **5**: frequently used Referer headers (ranked by the number of requests)
	//
	// 	- **7**: frequently used Referer headers (ranked by the amount of network traffic)
	//
	// 	- **9**: frequently requested URLs that are redirected to the origin (ranked by the number of requests)
	//
	// 	- **11**: frequently requested URLs that are redirected to the origin (ranked by the amount of network traffic)
	//
	// 	- **13**: top client IP addresses (ranked by the number of requests)
	//
	// 	- **15**: top client IP addresses (ranked by the amount of network traffic)
	//
	// 	- **17**: domain names ranked by the amount of network traffic
	//
	// 	- **19**: page views and unique visitors
	//
	// 	- **21**: regions from which requests are initiated
	//
	// 	- **23**: Internet service providers (ISPs)
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,3,5
	ReportIds *string `json:"ReportIds,omitempty" xml:"ReportIds,omitempty"`
}

func (s CreateCdnSubTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCdnSubTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateCdnSubTaskRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateCdnSubTaskRequest) GetReportIds() *string {
	return s.ReportIds
}

func (s *CreateCdnSubTaskRequest) SetDomainName(v string) *CreateCdnSubTaskRequest {
	s.DomainName = &v
	return s
}

func (s *CreateCdnSubTaskRequest) SetReportIds(v string) *CreateCdnSubTaskRequest {
	s.ReportIds = &v
	return s
}

func (s *CreateCdnSubTaskRequest) Validate() error {
	return dara.Validate(s)
}
