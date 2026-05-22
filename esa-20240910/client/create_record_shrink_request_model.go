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
	// The origin authentication information of the CNAME record.
	AuthConfShrink *string `json:"AuthConf,omitempty" xml:"AuthConf,omitempty"`
	// The business scenario of the record for acceleration. Leave the parameter empty if your record is not proxied. Valid values:
	//
	// 	- **image_video**: video and image.
	//
	// 	- **api**: API.
	//
	// 	- **web**: web page.
	//
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The comment of the record. The maximum length is 100 characters.
	//
	// example:
	//
	// This is a remark.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The DNS record information. The format of this field varies based on the record type. For more information, see [References](https://www.alibabacloud.com/help/doc-detail/2708761.html) .
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
	// The origin host policy. This policy takes effect when the record type is CNAME. You can set the policy in two modes:
	//
	// 	- follow_hostname: Follow the host record.
	//
	// 	- follow_origin_domain: match the origin\\"s domain name.
	//
	// example:
	//
	// follow_origin_domain
	HostPolicy *string `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	// Specifies whether to proxy the record. Only CNAME and A/AAAA records can be proxied. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
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
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The origin type for the CNAME record. This parameter is required when you add a CNAME record. Valid values:
	//
	// 	- **OSS**: OSS bucket.
	//
	// 	- **S3**: S3 bucket.
	//
	// 	- **LB**: load balancer.
	//
	// 	- **OP**: origin pool.
	//
	// 	- **Domain**: domain name.
	//
	// If you do not pass this parameter or if you leave its value empty, Domain is used by default.
	//
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The TTL of the record. Unit: seconds. If the value is 1, the TTL of the record is determined by the system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type of the DNS record. For example, A/AAAA, TXT, MX, or CNAME.
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
