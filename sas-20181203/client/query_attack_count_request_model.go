// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAttackCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *QueryAttackCountRequest
	GetFrom() *string
	SetLang(v string) *QueryAttackCountRequest
	GetLang() *string
	SetSourceIp(v string) *QueryAttackCountRequest
	GetSourceIp() *string
	SetUuids(v string) *QueryAttackCountRequest
	GetUuids() *string
}

type QueryAttackCountRequest struct {
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
	// The source IP address of the request.
	//
	// example:
	//
	// 175.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The UUID of the asset.
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of assets.
	//
	// example:
	//
	// 0c1714dc-f7a3-4265-8364-7aa3fce8****,1cc45e7d-7698-4b2c-89d8-e8cba407****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s QueryAttackCountRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAttackCountRequest) GoString() string {
	return s.String()
}

func (s *QueryAttackCountRequest) GetFrom() *string {
	return s.From
}

func (s *QueryAttackCountRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryAttackCountRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *QueryAttackCountRequest) GetUuids() *string {
	return s.Uuids
}

func (s *QueryAttackCountRequest) SetFrom(v string) *QueryAttackCountRequest {
	s.From = &v
	return s
}

func (s *QueryAttackCountRequest) SetLang(v string) *QueryAttackCountRequest {
	s.Lang = &v
	return s
}

func (s *QueryAttackCountRequest) SetSourceIp(v string) *QueryAttackCountRequest {
	s.SourceIp = &v
	return s
}

func (s *QueryAttackCountRequest) SetUuids(v string) *QueryAttackCountRequest {
	s.Uuids = &v
	return s
}

func (s *QueryAttackCountRequest) Validate() error {
	return dara.Validate(s)
}
