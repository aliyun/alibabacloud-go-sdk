// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTerminalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *ListTerminalRequest
	GetAlias() *string
	SetBuildId(v string) *ListTerminalRequest
	GetBuildId() *string
	SetClientType(v int32) *ListTerminalRequest
	GetClientType() *int32
	SetInManage(v bool) *ListTerminalRequest
	GetInManage() *bool
	SetIpv4(v string) *ListTerminalRequest
	GetIpv4() *string
	SetLocationInfo(v string) *ListTerminalRequest
	GetLocationInfo() *string
	SetMaxResults(v int32) *ListTerminalRequest
	GetMaxResults() *int32
	SetModel(v string) *ListTerminalRequest
	GetModel() *string
	SetNextToken(v string) *ListTerminalRequest
	GetNextToken() *string
	SetSearchKeyword(v string) *ListTerminalRequest
	GetSearchKeyword() *string
	SetSerialNumber(v string) *ListTerminalRequest
	GetSerialNumber() *string
	SetTerminalGroupId(v string) *ListTerminalRequest
	GetTerminalGroupId() *string
	SetUuid(v string) *ListTerminalRequest
	GetUuid() *string
}

type ListTerminalRequest struct {
	Alias           *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	BuildId         *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	ClientType      *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	InManage        *bool   `json:"InManage,omitempty" xml:"InManage,omitempty"`
	Ipv4            *string `json:"Ipv4,omitempty" xml:"Ipv4,omitempty"`
	LocationInfo    *string `json:"LocationInfo,omitempty" xml:"LocationInfo,omitempty"`
	MaxResults      *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Model           *string `json:"Model,omitempty" xml:"Model,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SearchKeyword   *string `json:"SearchKeyword,omitempty" xml:"SearchKeyword,omitempty"`
	SerialNumber    *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	TerminalGroupId *string `json:"TerminalGroupId,omitempty" xml:"TerminalGroupId,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListTerminalRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTerminalRequest) GoString() string {
	return s.String()
}

func (s *ListTerminalRequest) GetAlias() *string {
	return s.Alias
}

func (s *ListTerminalRequest) GetBuildId() *string {
	return s.BuildId
}

func (s *ListTerminalRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *ListTerminalRequest) GetInManage() *bool {
	return s.InManage
}

func (s *ListTerminalRequest) GetIpv4() *string {
	return s.Ipv4
}

func (s *ListTerminalRequest) GetLocationInfo() *string {
	return s.LocationInfo
}

func (s *ListTerminalRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTerminalRequest) GetModel() *string {
	return s.Model
}

func (s *ListTerminalRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTerminalRequest) GetSearchKeyword() *string {
	return s.SearchKeyword
}

func (s *ListTerminalRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ListTerminalRequest) GetTerminalGroupId() *string {
	return s.TerminalGroupId
}

func (s *ListTerminalRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListTerminalRequest) SetAlias(v string) *ListTerminalRequest {
	s.Alias = &v
	return s
}

func (s *ListTerminalRequest) SetBuildId(v string) *ListTerminalRequest {
	s.BuildId = &v
	return s
}

func (s *ListTerminalRequest) SetClientType(v int32) *ListTerminalRequest {
	s.ClientType = &v
	return s
}

func (s *ListTerminalRequest) SetInManage(v bool) *ListTerminalRequest {
	s.InManage = &v
	return s
}

func (s *ListTerminalRequest) SetIpv4(v string) *ListTerminalRequest {
	s.Ipv4 = &v
	return s
}

func (s *ListTerminalRequest) SetLocationInfo(v string) *ListTerminalRequest {
	s.LocationInfo = &v
	return s
}

func (s *ListTerminalRequest) SetMaxResults(v int32) *ListTerminalRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTerminalRequest) SetModel(v string) *ListTerminalRequest {
	s.Model = &v
	return s
}

func (s *ListTerminalRequest) SetNextToken(v string) *ListTerminalRequest {
	s.NextToken = &v
	return s
}

func (s *ListTerminalRequest) SetSearchKeyword(v string) *ListTerminalRequest {
	s.SearchKeyword = &v
	return s
}

func (s *ListTerminalRequest) SetSerialNumber(v string) *ListTerminalRequest {
	s.SerialNumber = &v
	return s
}

func (s *ListTerminalRequest) SetTerminalGroupId(v string) *ListTerminalRequest {
	s.TerminalGroupId = &v
	return s
}

func (s *ListTerminalRequest) SetUuid(v string) *ListTerminalRequest {
	s.Uuid = &v
	return s
}

func (s *ListTerminalRequest) Validate() error {
	return dara.Validate(s)
}
