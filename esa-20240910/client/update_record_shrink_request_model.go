// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecordShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfShrink(v string) *UpdateRecordShrinkRequest
	GetAuthConfShrink() *string
	SetBizName(v string) *UpdateRecordShrinkRequest
	GetBizName() *string
	SetComment(v string) *UpdateRecordShrinkRequest
	GetComment() *string
	SetDataShrink(v string) *UpdateRecordShrinkRequest
	GetDataShrink() *string
	SetHostPolicy(v string) *UpdateRecordShrinkRequest
	GetHostPolicy() *string
	SetProxied(v bool) *UpdateRecordShrinkRequest
	GetProxied() *bool
	SetRecordId(v int64) *UpdateRecordShrinkRequest
	GetRecordId() *int64
	SetSourceType(v string) *UpdateRecordShrinkRequest
	GetSourceType() *string
	SetTtl(v int32) *UpdateRecordShrinkRequest
	GetTtl() *int32
	SetType(v string) *UpdateRecordShrinkRequest
	GetType() *string
}

type UpdateRecordShrinkRequest struct {
	// The origin authentication information of the CNAME record.
	AuthConfShrink *string `json:"AuthConf,omitempty" xml:"AuthConf,omitempty"`
	// The business scenario of the record for acceleration. Leave the parameter empty if your record is not proxied. Valid values:
	//
	// 	- **video_image**: video and image.
	//
	// 	- **api**: API.
	//
	// 	- **web**: web page.
	//
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The comments of the record.
	//
	// example:
	//
	// This is a remark.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The DNS record information. The format of this field varies based on the record type. For more information, see [Add DNS records](https://www.alibabacloud.com/help/doc-detail/2708761.html).
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
	// 	- **follow_hostname**: match the requested domain name.
	//
	// 	- **follow_origin_domain**: match the origin\\"s domain name.
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
	// The record ID, which can be obtained by calling [ListRecords](https://help.aliyun.com/document_detail/2850265.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The type of the origin for the CNAME record. This parameter is required when you add a CNAME record. Valid values:
	//
	// 	- **OSS*	- : OSS origin.
	//
	// 	- **S3*	- : S3 origin.
	//
	// 	- **LB**: Load Balancer origin.
	//
	// 	- **OP**: origin in an origin pool.
	//
	// 	- **Domain**: common domain name.
	//
	// If you leave the parameter empty or set its value as null, the default is Domain, which is common domain name.
	//
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The TTL of the record. Unit: seconds. The range is 30 to 86,400, or 1. If the value is 1, the TTL of the record is determined by the system.
	//
	// example:
	//
	// 30
	Ttl  *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateRecordShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecordShrinkRequest) GetAuthConfShrink() *string {
	return s.AuthConfShrink
}

func (s *UpdateRecordShrinkRequest) GetBizName() *string {
	return s.BizName
}

func (s *UpdateRecordShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateRecordShrinkRequest) GetDataShrink() *string {
	return s.DataShrink
}

func (s *UpdateRecordShrinkRequest) GetHostPolicy() *string {
	return s.HostPolicy
}

func (s *UpdateRecordShrinkRequest) GetProxied() *bool {
	return s.Proxied
}

func (s *UpdateRecordShrinkRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *UpdateRecordShrinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateRecordShrinkRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *UpdateRecordShrinkRequest) GetType() *string {
	return s.Type
}

func (s *UpdateRecordShrinkRequest) SetAuthConfShrink(v string) *UpdateRecordShrinkRequest {
	s.AuthConfShrink = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetBizName(v string) *UpdateRecordShrinkRequest {
	s.BizName = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetComment(v string) *UpdateRecordShrinkRequest {
	s.Comment = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetDataShrink(v string) *UpdateRecordShrinkRequest {
	s.DataShrink = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetHostPolicy(v string) *UpdateRecordShrinkRequest {
	s.HostPolicy = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetProxied(v bool) *UpdateRecordShrinkRequest {
	s.Proxied = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetRecordId(v int64) *UpdateRecordShrinkRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetSourceType(v string) *UpdateRecordShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetTtl(v int32) *UpdateRecordShrinkRequest {
	s.Ttl = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetType(v string) *UpdateRecordShrinkRequest {
	s.Type = &v
	return s
}

func (s *UpdateRecordShrinkRequest) Validate() error {
	return dara.Validate(s)
}
