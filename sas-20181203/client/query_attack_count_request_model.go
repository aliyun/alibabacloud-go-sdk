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
	// The source identifier of the request. Set this parameter to sas.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 175.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The UUID of the server. Separate multiple UUIDs with commas (,).
	//
	// > Call the [DescribeCloudCenterInstances](https://help.aliyun.com/document_detail/141932.html) operation to obtain this parameter.
	//
	// example:
	//
	// 1587bedb-fdb4-48c4-9330-************
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
