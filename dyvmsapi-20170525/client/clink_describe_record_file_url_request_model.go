// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeRecordFileUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDownload(v int64) *ClinkDescribeRecordFileUrlRequest
	GetDownload() *int64
	SetEnterpriseId(v int64) *ClinkDescribeRecordFileUrlRequest
	GetEnterpriseId() *int64
	SetMainUniqueId(v string) *ClinkDescribeRecordFileUrlRequest
	GetMainUniqueId() *string
	SetOwnerId(v int64) *ClinkDescribeRecordFileUrlRequest
	GetOwnerId() *int64
	SetRecordSide(v int64) *ClinkDescribeRecordFileUrlRequest
	GetRecordSide() *int64
	SetRecordType(v string) *ClinkDescribeRecordFileUrlRequest
	GetRecordType() *string
	SetResourceOwnerAccount(v string) *ClinkDescribeRecordFileUrlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkDescribeRecordFileUrlRequest
	GetResourceOwnerId() *int64
	SetTimeout(v int64) *ClinkDescribeRecordFileUrlRequest
	GetTimeout() *int64
}

type ClinkDescribeRecordFileUrlRequest struct {
	// ０表示试听，１为下载，默认为1
	//
	// example:
	//
	// 1
	Download *int64 `json:"Download,omitempty" xml:"Download,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 通话记录唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// uniq_1-162046xxxx.12
	MainUniqueId *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 不传递获取mp3格式录音，传递时获取wav格式录音。1：双轨录音客户侧，2：双轨录音座席侧，3：两侧合成录音
	//
	// example:
	//
	// 1
	RecordSide *int64 `json:"RecordSide,omitempty" xml:"RecordSide,omitempty"`
	// 录音类型。 "record": 通话录音，"voicemail": 留言。默认值为 "record"
	//
	// example:
	//
	// record
	RecordType           *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 获取录音地址超时时长，单位为秒，默认为一小时，范围在一到二十四小时。
	//
	// example:
	//
	// 67
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s ClinkDescribeRecordFileUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeRecordFileUrlRequest) GoString() string {
	return s.String()
}

func (s *ClinkDescribeRecordFileUrlRequest) GetDownload() *int64 {
	return s.Download
}

func (s *ClinkDescribeRecordFileUrlRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkDescribeRecordFileUrlRequest) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *ClinkDescribeRecordFileUrlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkDescribeRecordFileUrlRequest) GetRecordSide() *int64 {
	return s.RecordSide
}

func (s *ClinkDescribeRecordFileUrlRequest) GetRecordType() *string {
	return s.RecordType
}

func (s *ClinkDescribeRecordFileUrlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkDescribeRecordFileUrlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkDescribeRecordFileUrlRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *ClinkDescribeRecordFileUrlRequest) SetDownload(v int64) *ClinkDescribeRecordFileUrlRequest {
	s.Download = &v
	return s
}

func (s *ClinkDescribeRecordFileUrlRequest) SetEnterpriseId(v int64) *ClinkDescribeRecordFileUrlRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkDescribeRecordFileUrlRequest) SetMainUniqueId(v string) *ClinkDescribeRecordFileUrlRequest {
	s.MainUniqueId = &v
	return s
}

func (s *ClinkDescribeRecordFileUrlRequest) SetOwnerId(v int64) *ClinkDescribeRecordFileUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkDescribeRecordFileUrlRequest) SetRecordSide(v int64) *ClinkDescribeRecordFileUrlRequest {
	s.RecordSide = &v
	return s
}

func (s *ClinkDescribeRecordFileUrlRequest) SetRecordType(v string) *ClinkDescribeRecordFileUrlRequest {
	s.RecordType = &v
	return s
}

func (s *ClinkDescribeRecordFileUrlRequest) SetResourceOwnerAccount(v string) *ClinkDescribeRecordFileUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkDescribeRecordFileUrlRequest) SetResourceOwnerId(v int64) *ClinkDescribeRecordFileUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkDescribeRecordFileUrlRequest) SetTimeout(v int64) *ClinkDescribeRecordFileUrlRequest {
	s.Timeout = &v
	return s
}

func (s *ClinkDescribeRecordFileUrlRequest) Validate() error {
	return dara.Validate(s)
}
