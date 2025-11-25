// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvadeEventDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetsInstanceId(v string) *DescribeInvadeEventDetailRequest
	GetAssetsInstanceId() *string
	SetEventUuid(v string) *DescribeInvadeEventDetailRequest
	GetEventUuid() *string
	SetLang(v string) *DescribeInvadeEventDetailRequest
	GetLang() *string
	SetPublicIP(v string) *DescribeInvadeEventDetailRequest
	GetPublicIP() *string
	SetSourceIp(v string) *DescribeInvadeEventDetailRequest
	GetSourceIp() *string
}

type DescribeInvadeEventDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// i-bp135d2rmbwpt****
	AssetsInstanceId *string `json:"AssetsInstanceId,omitempty" xml:"AssetsInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 92498578-7c42-4845-8c73-7e824782****
	EventUuid *string `json:"EventUuid,omitempty" xml:"EventUuid,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 182.92.103.XXX
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// example:
	//
	// 218.76.30.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeInvadeEventDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventDetailRequest) GetAssetsInstanceId() *string {
	return s.AssetsInstanceId
}

func (s *DescribeInvadeEventDetailRequest) GetEventUuid() *string {
	return s.EventUuid
}

func (s *DescribeInvadeEventDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInvadeEventDetailRequest) GetPublicIP() *string {
	return s.PublicIP
}

func (s *DescribeInvadeEventDetailRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInvadeEventDetailRequest) SetAssetsInstanceId(v string) *DescribeInvadeEventDetailRequest {
	s.AssetsInstanceId = &v
	return s
}

func (s *DescribeInvadeEventDetailRequest) SetEventUuid(v string) *DescribeInvadeEventDetailRequest {
	s.EventUuid = &v
	return s
}

func (s *DescribeInvadeEventDetailRequest) SetLang(v string) *DescribeInvadeEventDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInvadeEventDetailRequest) SetPublicIP(v string) *DescribeInvadeEventDetailRequest {
	s.PublicIP = &v
	return s
}

func (s *DescribeInvadeEventDetailRequest) SetSourceIp(v string) *DescribeInvadeEventDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInvadeEventDetailRequest) Validate() error {
	return dara.Validate(s)
}
