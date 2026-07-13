// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmMonitorTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListCloudGtmMonitorTemplatesRequest
	GetAcceptLanguage() *string
	SetIpVersion(v string) *ListCloudGtmMonitorTemplatesRequest
	GetIpVersion() *string
	SetName(v string) *ListCloudGtmMonitorTemplatesRequest
	GetName() *string
	SetPageNumber(v int32) *ListCloudGtmMonitorTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCloudGtmMonitorTemplatesRequest
	GetPageSize() *int32
	SetProtocol(v string) *ListCloudGtmMonitorTemplatesRequest
	GetProtocol() *string
}

type ListCloudGtmMonitorTemplatesRequest struct {
	// The language of the response. Valid values:
	//
	// - zh-CN: Chinese.
	//
	// - en-US (default): English.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The IP version of the detection points.
	//
	// - IPv4: The destination address is an IPv4 address.
	//
	// - IPv6: The destination address is an IPv6 address.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The name of the health check template. Name the template in a way that helps you distinguish between different health check protocols.
	//
	// example:
	//
	// IPv4-Ping
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number. The value starts from **1**. The default value is **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The protocol used for the health check.
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

func (s ListCloudGtmMonitorTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorTemplatesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListCloudGtmMonitorTemplatesRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *ListCloudGtmMonitorTemplatesRequest) GetName() *string {
	return s.Name
}

func (s *ListCloudGtmMonitorTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCloudGtmMonitorTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudGtmMonitorTemplatesRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *ListCloudGtmMonitorTemplatesRequest) SetAcceptLanguage(v string) *ListCloudGtmMonitorTemplatesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesRequest) SetIpVersion(v string) *ListCloudGtmMonitorTemplatesRequest {
	s.IpVersion = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesRequest) SetName(v string) *ListCloudGtmMonitorTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesRequest) SetPageNumber(v int32) *ListCloudGtmMonitorTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesRequest) SetPageSize(v int32) *ListCloudGtmMonitorTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesRequest) SetProtocol(v string) *ListCloudGtmMonitorTemplatesRequest {
	s.Protocol = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
