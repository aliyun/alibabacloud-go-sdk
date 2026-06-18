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
	// The origin authentication settings for the CNAME record.
	AuthConfShrink *string `json:"AuthConf,omitempty" xml:"AuthConf,omitempty"`
	// The use case for proxy acceleration. Omit this parameter if proxy acceleration is disabled. Valid values:
	//
	// - **video_image**: Video and images.
	//
	// - **api**: APIs.
	//
	// - **web**: Web pages.
	//
	// example:
	//
	// web
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// A comment for the record.
	//
	// example:
	//
	// This is a remark.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The DNS data for the record. The required content varies based on the record type. For more information, see <props="china">[Documentation](https://help.aliyun.com/document_detail/2708761.html)<props="intl">[Documentation](https://www.alibabacloud.com/help/doc-detail/2708761.html).
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
	// The origin HOST policy. This policy, which applies only to CNAME records, determines the value of the `HOST` header in requests sent to the origin. Valid values:
	//
	// - **follow_hostname**: Follows the host record.
	//
	// - **follow_origin_domain**: Follows the origin domain name.
	//
	// example:
	//
	// follow_origin_domain
	HostPolicy *string `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	HttpPorts  *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// Indicates whether to enable proxy acceleration for the record. Only CNAME and A/AAAA records support proxy acceleration. Valid values:
	//
	// - **true**: Enables proxy acceleration.
	//
	// - **false**: Disables proxy acceleration.
	//
	// example:
	//
	// true
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// The record ID. Call the [ListRecords](https://help.aliyun.com/document_detail/2850265.html) operation to get this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The origin type for the CNAME record. This parameter is required for CNAME records. Valid values:
	//
	// - **OSS**: An OSS origin.
	//
	// - **S3**: An S3 origin.
	//
	// - **LB**: A load balancer origin.
	//
	// - **OP**: An origin address pool origin.
	//
	// - **Domain**: A standard domain name origin.
	//
	// If this parameter is omitted or left empty, the default value is `Domain`.
	//
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The record\\"s time to live (TTL) in seconds. The value must be an integer from **30 to 86400*	- or 1. A value of 1 sets the TTL to automatic.
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
