// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGroupedSecurityEventMarkMissListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryGroupedSecurityEventMarkMissListRequest
	GetCurrentPage() *int32
	SetDisposalWay(v string) *QueryGroupedSecurityEventMarkMissListRequest
	GetDisposalWay() *string
	SetEventName(v string) *QueryGroupedSecurityEventMarkMissListRequest
	GetEventName() *string
	SetFrom(v string) *QueryGroupedSecurityEventMarkMissListRequest
	GetFrom() *string
	SetLang(v string) *QueryGroupedSecurityEventMarkMissListRequest
	GetLang() *string
	SetPageSize(v int32) *QueryGroupedSecurityEventMarkMissListRequest
	GetPageSize() *int32
	SetRemark(v string) *QueryGroupedSecurityEventMarkMissListRequest
	GetRemark() *string
	SetSourceIp(v string) *QueryGroupedSecurityEventMarkMissListRequest
	GetSourceIp() *string
}

type QueryGroupedSecurityEventMarkMissListRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The handling method. Valid values:
	//
	// 	- **1**: Automatically Added to Whitelist
	//
	// 	- **2**: Defense Without Notification
	//
	// example:
	//
	// 1
	DisposalWay *string `json:"DisposalWay,omitempty" xml:"DisposalWay,omitempty"`
	// The name of the alert event. The value indicates a subtype.
	//
	// example:
	//
	// Login with unusual location
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The ID of the request source. Set the value to sas.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The condition that is used to query alert events by asset. You can specify a value of the following types:
	//
	// 	- The IP address of the asset.
	//
	// 	- The public IP address of the asset.
	//
	// 	- The private IP address of the asset.
	//
	// 	- The name of the asset.
	//
	// example:
	//
	// 222.185.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 113.66.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s QueryGroupedSecurityEventMarkMissListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupedSecurityEventMarkMissListRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupedSecurityEventMarkMissListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryGroupedSecurityEventMarkMissListRequest) GetDisposalWay() *string {
	return s.DisposalWay
}

func (s *QueryGroupedSecurityEventMarkMissListRequest) GetEventName() *string {
	return s.EventName
}

func (s *QueryGroupedSecurityEventMarkMissListRequest) GetFrom() *string {
	return s.From
}

func (s *QueryGroupedSecurityEventMarkMissListRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryGroupedSecurityEventMarkMissListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryGroupedSecurityEventMarkMissListRequest) GetRemark() *string {
	return s.Remark
}

func (s *QueryGroupedSecurityEventMarkMissListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *QueryGroupedSecurityEventMarkMissListRequest) SetCurrentPage(v int32) *QueryGroupedSecurityEventMarkMissListRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListRequest) SetDisposalWay(v string) *QueryGroupedSecurityEventMarkMissListRequest {
	s.DisposalWay = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListRequest) SetEventName(v string) *QueryGroupedSecurityEventMarkMissListRequest {
	s.EventName = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListRequest) SetFrom(v string) *QueryGroupedSecurityEventMarkMissListRequest {
	s.From = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListRequest) SetLang(v string) *QueryGroupedSecurityEventMarkMissListRequest {
	s.Lang = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListRequest) SetPageSize(v int32) *QueryGroupedSecurityEventMarkMissListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListRequest) SetRemark(v string) *QueryGroupedSecurityEventMarkMissListRequest {
	s.Remark = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListRequest) SetSourceIp(v string) *QueryGroupedSecurityEventMarkMissListRequest {
	s.SourceIp = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListRequest) Validate() error {
	return dara.Validate(s)
}
