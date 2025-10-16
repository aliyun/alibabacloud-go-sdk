// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkInstanceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribeNetworkInstanceListRequest
	GetCenId() *string
	SetConnectType(v string) *DescribeNetworkInstanceListRequest
	GetConnectType() *string
	SetLang(v string) *DescribeNetworkInstanceListRequest
	GetLang() *string
}

type DescribeNetworkInstanceListRequest struct {
	// example:
	//
	// cen-x5jayxou71ad73****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// example:
	//
	// expressconnect
	ConnectType *string `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeNetworkInstanceListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInstanceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInstanceListRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribeNetworkInstanceListRequest) GetConnectType() *string {
	return s.ConnectType
}

func (s *DescribeNetworkInstanceListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNetworkInstanceListRequest) SetCenId(v string) *DescribeNetworkInstanceListRequest {
	s.CenId = &v
	return s
}

func (s *DescribeNetworkInstanceListRequest) SetConnectType(v string) *DescribeNetworkInstanceListRequest {
	s.ConnectType = &v
	return s
}

func (s *DescribeNetworkInstanceListRequest) SetLang(v string) *DescribeNetworkInstanceListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNetworkInstanceListRequest) Validate() error {
	return dara.Validate(s)
}
