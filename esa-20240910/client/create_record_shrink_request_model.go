// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecordShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfShrink(v string) *CreateRecordShrinkRequest
	GetAuthConfShrink() *string
	SetBizName(v string) *CreateRecordShrinkRequest
	GetBizName() *string
	SetComment(v string) *CreateRecordShrinkRequest
	GetComment() *string
	SetDataShrink(v string) *CreateRecordShrinkRequest
	GetDataShrink() *string
	SetHostPolicy(v string) *CreateRecordShrinkRequest
	GetHostPolicy() *string
	SetHttpPorts(v string) *CreateRecordShrinkRequest
	GetHttpPorts() *string
	SetHttpsPorts(v string) *CreateRecordShrinkRequest
	GetHttpsPorts() *string
	SetProxied(v bool) *CreateRecordShrinkRequest
	GetProxied() *bool
	SetRecordName(v string) *CreateRecordShrinkRequest
	GetRecordName() *string
	SetSiteId(v int64) *CreateRecordShrinkRequest
	GetSiteId() *int64
	SetSourceType(v string) *CreateRecordShrinkRequest
	GetSourceType() *string
	SetTtl(v int32) *CreateRecordShrinkRequest
	GetTtl() *int32
	SetType(v string) *CreateRecordShrinkRequest
	GetType() *string
}

type CreateRecordShrinkRequest struct {
	// The origin authentication information for the CNAME record.
	AuthConfShrink *string `json:"AuthConf,omitempty" xml:"AuthConf,omitempty"`
	// Used to tag the business scenario of the DNS record. This parameter is required when proxy acceleration is enabled for the DNS record (i.e., when the proxied parameter is set to true), and is not required when proxy acceleration is disabled (i.e., when the proxied parameter is set to false). Valid values:
	//
	// - **image_video**: Image and video.
	//
	// - **api**: API.
	//
	// - **web**: Web page.
	//
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The comment for the record, with a maximum length of 100 characters.
	//
	// example:
	//
	// This is a remark.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The DNS information of the record. Different types of records require different content for this field. For more information, see
	//
	// <props="china">[Documentation](https://help.aliyun.com/document_detail/2708761.html)<props="intl">[Documentation](https://www.alibabacloud.com/help/doc-detail/2708761.html)
	//
	// .
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "value":"2.2.2.2"
	//
	// }
	DataShrink *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The origin host policy. This takes effect when the record type is CNAME. It specifies the host policy for back-to-origin requests. Two modes are available:
	//
	// - **follow_hostname**: Follow the request host.
	//
	// - **follow_origin_domain**: Follow the origin domain.
	//
	// example:
	//
	// follow_origin_domain
	HostPolicy *string `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	HttpPorts  *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// Specifies whether to enable proxy acceleration for the record. Only CNAME records or A/AAAA records (i.e., when the type parameter is set to A/AAAA or CNAME) can enable proxy acceleration. Valid values:
	//
	// - **true**: Enable proxy acceleration.
	//
	// - **false**: Disable proxy acceleration.
	//
	// example:
	//
	// true
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// The record name.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// The site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The origin type of the CNAME record. This parameter is required when adding a CNAME record (i.e., when the type parameter is set to CNAME). Valid values:
	//
	// - **OSS**: OSS origin.
	//
	// - **S3**: S3 origin.
	//
	// - **LB**: Load balancer origin.
	//
	// - **OP**: Origin pool origin.
	//
	// - **Domain**: Standard domain origin.
	//
	// If this parameter is not specified or is left empty, it defaults to Domain, which is the standard domain origin type.
	//
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time-to-live (TTL) of the record, in seconds. When set to 1, the TTL is automatic.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The DNS type of the record, such as **A/AAAA**, **CNAME**, **TXT**, etc.
	//
	// This parameter is required.
	//
	// example:
	//
	// A/AAAA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRecordShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRecordShrinkRequest) GetAuthConfShrink() *string {
	return s.AuthConfShrink
}

func (s *CreateRecordShrinkRequest) GetBizName() *string {
	return s.BizName
}

func (s *CreateRecordShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateRecordShrinkRequest) GetDataShrink() *string {
	return s.DataShrink
}

func (s *CreateRecordShrinkRequest) GetHostPolicy() *string {
	return s.HostPolicy
}

func (s *CreateRecordShrinkRequest) GetHttpPorts() *string {
	return s.HttpPorts
}

func (s *CreateRecordShrinkRequest) GetHttpsPorts() *string {
	return s.HttpsPorts
}

func (s *CreateRecordShrinkRequest) GetProxied() *bool {
	return s.Proxied
}

func (s *CreateRecordShrinkRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *CreateRecordShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateRecordShrinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateRecordShrinkRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *CreateRecordShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateRecordShrinkRequest) SetAuthConfShrink(v string) *CreateRecordShrinkRequest {
	s.AuthConfShrink = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetBizName(v string) *CreateRecordShrinkRequest {
	s.BizName = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetComment(v string) *CreateRecordShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetDataShrink(v string) *CreateRecordShrinkRequest {
	s.DataShrink = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetHostPolicy(v string) *CreateRecordShrinkRequest {
	s.HostPolicy = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetHttpPorts(v string) *CreateRecordShrinkRequest {
	s.HttpPorts = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetHttpsPorts(v string) *CreateRecordShrinkRequest {
	s.HttpsPorts = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetProxied(v bool) *CreateRecordShrinkRequest {
	s.Proxied = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetRecordName(v string) *CreateRecordShrinkRequest {
	s.RecordName = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetSiteId(v int64) *CreateRecordShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetSourceType(v string) *CreateRecordShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetTtl(v int32) *CreateRecordShrinkRequest {
	s.Ttl = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetType(v string) *CreateRecordShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateRecordShrinkRequest) Validate() error {
	return dara.Validate(s)
}
