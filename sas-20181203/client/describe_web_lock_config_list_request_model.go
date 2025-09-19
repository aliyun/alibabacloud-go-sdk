// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockConfigListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DescribeWebLockConfigListRequest
	GetId() *int64
	SetLang(v string) *DescribeWebLockConfigListRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeWebLockConfigListRequest
	GetSourceIp() *string
	SetUuid(v string) *DescribeWebLockConfigListRequest
	GetUuid() *string
}

type DescribeWebLockConfigListRequest struct {
	// The configuration ID of the protected directory.
	//
	// example:
	//
	// 1404656
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Valid values:
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
	// 1.2.3.4
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The UUID of the server.
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// inet-1234567****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeWebLockConfigListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockConfigListRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebLockConfigListRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeWebLockConfigListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeWebLockConfigListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeWebLockConfigListRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeWebLockConfigListRequest) SetId(v int64) *DescribeWebLockConfigListRequest {
	s.Id = &v
	return s
}

func (s *DescribeWebLockConfigListRequest) SetLang(v string) *DescribeWebLockConfigListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWebLockConfigListRequest) SetSourceIp(v string) *DescribeWebLockConfigListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebLockConfigListRequest) SetUuid(v string) *DescribeWebLockConfigListRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeWebLockConfigListRequest) Validate() error {
	return dara.Validate(s)
}
