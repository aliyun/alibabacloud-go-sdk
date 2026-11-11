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
	// The origin server authentication information for a CNAME record.
	AuthConfShrink *string `json:"AuthConf,omitempty" xml:"AuthConf,omitempty"`
	// The business scenario used to tag the DNS record. This parameter is required when proxy acceleration is enabled for the DNS record. That is, the Proxied parameter is set to true. This parameter is not required when proxy acceleration is disabled. That is, the Proxied parameter is set to false. Valid values:
	//
	// - **image_video**: Images and videos.
	//
	// - **api**: API operations.
	//
	// - **web**: Web pages.
	//
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The description of the record. The description can be up to 100 characters in length.
	//
	// example:
	//
	// This is a remark.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The DNS information of the record. The required content for this field varies based on the record type. For more information, see
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
	// The back-to-origin host policy. This parameter takes effect when the record type is CNAME. Settings for the host policy in back-to-origin requests. Two patterns are available:
	//
	// - **follow_hostname**: Follows the request host.
	//
	// - **follow_origin_domain**: Follows the Origin Domain Name.
	//
	// example:
	//
	// follow_origin_domain
	HostPolicy *string `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	// The HTTP port. By default, port 80 is supported.
	//
	// example:
	//
	// 80,2052,2082,2086,2095,8080,8880
	HttpPorts *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	// The HTTPS port. By default, port 443 is supported.
	//
	// example:
	//
	// 443,2053,2083,2087,2096,8443
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// Specifies whether to enable proxy acceleration for the record. Proxy acceleration can be enabled only for CNAME, A, or AAAA records. That is, the Type parameter is set to CNAME, A, or AAAA. Valid values:
	//
	// - **true**: Enables proxy acceleration.
	//
	// - **false**: Disables proxy acceleration.
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
	// The site ID. You can obtain the site ID by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The origin server type for a CNAME record. This parameter is required when you add a CNAME record. That is, the Type parameter is set to CNAME. Valid values:
	//
	// - **OSS**: OSS origin server.
	//
	// - **S3**: S3 origin server.
	//
	// - **LB**: Load balancing origin server.
	//
	// - **OP**: IPAM pool origin server.
	//
	// - **Domain**: Common domain name origin server.
	//
	// If this parameter is not specified or is left empty, the default value is Domain, which indicates a common domain name origin server.
	//
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time-to-live (TTL) of the record, in seconds. A value of 1 indicates that the TTL is automatically determined.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The DNS record type, such as **A/AAAA**, **CNAME**, or **TXT**.
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
