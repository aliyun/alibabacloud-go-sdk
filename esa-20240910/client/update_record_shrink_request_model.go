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
	SetHttpPorts(v string) *UpdateRecordShrinkRequest
	GetHttpPorts() *string
	SetHttpsPorts(v string) *UpdateRecordShrinkRequest
	GetHttpsPorts() *string
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
	// The business scenario for record acceleration. This parameter is not required for records without acceleration enabled. Valid values:
	//
	// - **video_image**: video and image.
	//
	// - **api**: API.
	//
	// - **web**: web page.
	//
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The comment for the record.
	//
	// example:
	//
	// This is a remark.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The DNS information of the record. The content varies depending on the record type. For more information, see <props="china">[documentation](https://help.aliyun.com/document_detail/2708761.html)<props="intl">[documentation](https://www.alibabacloud.com/help/doc-detail/2708761.html).
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
	// The back-to-origin HOST policy. This parameter takes effect when the record type is CNAME. Settings the HOST policy for back-to-origin requests. Valid values:
	//
	// - **follow_hostname**: follows the host record.
	//
	// - **follow_origin_domain**: follows the Origin Domain Name.
	//
	// example:
	//
	// follow_origin_domain
	HostPolicy *string `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	HttpPorts  *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// Specifies whether to enable proxy acceleration for the record. Only CNAME records and A/AAAA records support proxy acceleration. Valid values:
	//
	// - **true**: Enable proxy acceleration.
	//
	// - **false**: Disable proxy acceleration.
	//
	// example:
	//
	// true
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// The ID of the record. You can call [ListRecords](https://help.aliyun.com/document_detail/2850265.html) to obtain the record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The origin server type of the CNAME record. This parameter is required when you add a CNAME record. Valid values:
	//
	// - **OSS**: OSS origin server.
	//
	// - **S3**: S3 origin server.
	//
	// - **LB**: load balancing origin server.
	//
	// - **OP**: IPAM pool origin server.
	//
	// - **Domain**: standard domain name origin server.
	//
	// If this parameter is not specified or is left empty, the default value is Domain, which indicates a standard domain name origin server type.
	//
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time-to-live (TTL) of the record, in seconds. Valid values: **30 to 86400**, or 1. A value of 1 indicates that the TTL of the record is automatically determined.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The DNS type of the record, such as A/AAAA, CNAME, or TXT.
	//
	// example:
	//
	// A/AAAA
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

func (s *UpdateRecordShrinkRequest) GetHttpPorts() *string {
	return s.HttpPorts
}

func (s *UpdateRecordShrinkRequest) GetHttpsPorts() *string {
	return s.HttpsPorts
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

func (s *UpdateRecordShrinkRequest) SetHttpPorts(v string) *UpdateRecordShrinkRequest {
	s.HttpPorts = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetHttpsPorts(v string) *UpdateRecordShrinkRequest {
	s.HttpsPorts = &v
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
