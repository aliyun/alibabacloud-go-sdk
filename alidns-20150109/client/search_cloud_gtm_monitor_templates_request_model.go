// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCloudGtmMonitorTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *SearchCloudGtmMonitorTemplatesRequest
	GetAcceptLanguage() *string
	SetIpVersion(v string) *SearchCloudGtmMonitorTemplatesRequest
	GetIpVersion() *string
	SetName(v string) *SearchCloudGtmMonitorTemplatesRequest
	GetName() *string
	SetPageNumber(v int32) *SearchCloudGtmMonitorTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchCloudGtmMonitorTemplatesRequest
	GetPageSize() *int32
	SetProtocol(v string) *SearchCloudGtmMonitorTemplatesRequest
	GetProtocol() *string
}

type SearchCloudGtmMonitorTemplatesRequest struct {
	// The language of the response. Valid values:
	//
	// - zh-CN: Chinese
	//
	// - en-US (default): English
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The IP version of the detection points. An exact search is performed. Valid values:
	//
	// - IPv4: This version is applicable when the target address is an IPv4 address.
	//
	// - IPv6: This version is applicable when the target address is an IPv6 address.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The name of the health check template. A fuzzy search is performed based on the keyword that you enter.
	//
	// example:
	//
	// IPv4-Ping
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The current page number. The value starts from 1. The default value is 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page for a paged query. The maximum value is **100**. The default value is **20**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The protocol used for detection. An exact search is performed. Valid values:
	//
	// - ping
	//
	// - tcp
	//
	// - http
	//
	// - https
	//
	// example:
	//
	// ping
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s SearchCloudGtmMonitorTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmMonitorTemplatesRequest) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmMonitorTemplatesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *SearchCloudGtmMonitorTemplatesRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *SearchCloudGtmMonitorTemplatesRequest) GetName() *string {
	return s.Name
}

func (s *SearchCloudGtmMonitorTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchCloudGtmMonitorTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchCloudGtmMonitorTemplatesRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *SearchCloudGtmMonitorTemplatesRequest) SetAcceptLanguage(v string) *SearchCloudGtmMonitorTemplatesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesRequest) SetIpVersion(v string) *SearchCloudGtmMonitorTemplatesRequest {
	s.IpVersion = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesRequest) SetName(v string) *SearchCloudGtmMonitorTemplatesRequest {
	s.Name = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesRequest) SetPageNumber(v int32) *SearchCloudGtmMonitorTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesRequest) SetPageSize(v int32) *SearchCloudGtmMonitorTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesRequest) SetProtocol(v string) *SearchCloudGtmMonitorTemplatesRequest {
	s.Protocol = &v
	return s
}

func (s *SearchCloudGtmMonitorTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
