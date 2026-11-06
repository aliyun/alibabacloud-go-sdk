// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNetListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMainNet(v string) *QueryNetListRequest
	GetMainNet() *string
	SetMode(v string) *QueryNetListRequest
	GetMode() *string
	SetNet(v string) *QueryNetListRequest
	GetNet() *string
	SetNum(v int64) *QueryNetListRequest
	GetNum() *int64
	SetPage(v int64) *QueryNetListRequest
	GetPage() *int64
	SetSaleId(v string) *QueryNetListRequest
	GetSaleId() *string
}

type QueryNetListRequest struct {
	// The primary CIDR block of the anti-DDoS diversion instance for which an extended CIDR block is configured. If no extended CIDR blocks are configured for the anti-DDoS diversion instance, leave this parameter empty.
	//
	// example:
	//
	// 192.168.XX.XX/22
	MainNet *string `json:"MainNet,omitempty" xml:"MainNet,omitempty"`
	// The scheduling mode. Valid values:
	//
	// - manual: manual scheduling.
	//
	// - netflow-auto: automatic scheduling.
	//
	// example:
	//
	// netflow-auto
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The CIDR block of the anti-DDoS diversion instance.
	//
	// > If no extended CIDR blocks are configured for the anti-DDoS diversion instance, this parameter specifies the CIDR block of the instance. If an extended CIDR block is configured for the anti-DDoS diversion instance, this parameter specifies the extended CIDR block that is configured for the instance. If this parameter is specified, the MainNet parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX/24
	Net *string `json:"Net,omitempty" xml:"Net,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 100
	Num *int64 `json:"Num,omitempty" xml:"Num,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The ID of the anti-DDoS diversion instance.
	//
	// example:
	//
	// ddos_diversion_public_cn-xxxxxxxxxxxxx
	SaleId *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
}

func (s QueryNetListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryNetListRequest) GoString() string {
	return s.String()
}

func (s *QueryNetListRequest) GetMainNet() *string {
	return s.MainNet
}

func (s *QueryNetListRequest) GetMode() *string {
	return s.Mode
}

func (s *QueryNetListRequest) GetNet() *string {
	return s.Net
}

func (s *QueryNetListRequest) GetNum() *int64 {
	return s.Num
}

func (s *QueryNetListRequest) GetPage() *int64 {
	return s.Page
}

func (s *QueryNetListRequest) GetSaleId() *string {
	return s.SaleId
}

func (s *QueryNetListRequest) SetMainNet(v string) *QueryNetListRequest {
	s.MainNet = &v
	return s
}

func (s *QueryNetListRequest) SetMode(v string) *QueryNetListRequest {
	s.Mode = &v
	return s
}

func (s *QueryNetListRequest) SetNet(v string) *QueryNetListRequest {
	s.Net = &v
	return s
}

func (s *QueryNetListRequest) SetNum(v int64) *QueryNetListRequest {
	s.Num = &v
	return s
}

func (s *QueryNetListRequest) SetPage(v int64) *QueryNetListRequest {
	s.Page = &v
	return s
}

func (s *QueryNetListRequest) SetSaleId(v string) *QueryNetListRequest {
	s.SaleId = &v
	return s
}

func (s *QueryNetListRequest) Validate() error {
	return dara.Validate(s)
}
