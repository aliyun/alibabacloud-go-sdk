// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateAsrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallType(v int64) *CloudCreateAsrRequest
	GetCallType() *int64
	SetEnterpriseId(v int64) *CloudCreateAsrRequest
	GetEnterpriseId() *int64
	SetMainUniqueId(v string) *CloudCreateAsrRequest
	GetMainUniqueId() *string
	SetRecordFile(v string) *CloudCreateAsrRequest
	GetRecordFile() *string
	SetRecordSide(v int64) *CloudCreateAsrRequest
	GetRecordSide() *int64
	SetRecordType(v string) *CloudCreateAsrRequest
	GetRecordType() *string
	SetSupportMp3(v string) *CloudCreateAsrRequest
	GetSupportMp3() *string
	SetUniqueId(v string) *CloudCreateAsrRequest
	GetUniqueId() *string
}

type CloudCreateAsrRequest struct {
	// 通话类型；1:呼入,2:webcall,4:预览外呼,5:预测外呼,6:主叫外呼,9:内部呼叫
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	CallType *int64 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 主通道的唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// sip-2-1489989807.36601
	MainUniqueId *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	// 录音文件名；例: 7000002-20170320140327-1581-1500-record-sip-2-148998.366
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002-20170320140327-1581-1500-record-sip-2-148998.366
	RecordFile *string `json:"RecordFile,omitempty" xml:"RecordFile,omitempty"`
	// 录音通道类型；取值范围：0：混音 1：分线录音时客户侧,2：分线录音时座席侧,3:分线录音时客户侧和座席侧
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	RecordSide *int64 `json:"RecordSide,omitempty" xml:"RecordSide,omitempty"`
	// 语音文件类型,取值：record,voicemail 说明：record是录音， voicemail是留言
	//
	// This parameter is required.
	//
	// example:
	//
	// record
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// 当wav过期时，是否支持使用mp3进行转写；0-不支持（默认）；1-代表支持，如果通话wav过期，则自动使用mp3转写，转写结果中通道类型会被设置为混音类型；混音mp3无法区分两侧，双声道mp3可以区分两侧内容
	//
	// example:
	//
	// 0
	SupportMp3 *string `json:"SupportMp3,omitempty" xml:"SupportMp3,omitempty"`
	// 从通道的唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// sip-2-1489989807.36601
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s CloudCreateAsrRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateAsrRequest) GoString() string {
	return s.String()
}

func (s *CloudCreateAsrRequest) GetCallType() *int64 {
	return s.CallType
}

func (s *CloudCreateAsrRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudCreateAsrRequest) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *CloudCreateAsrRequest) GetRecordFile() *string {
	return s.RecordFile
}

func (s *CloudCreateAsrRequest) GetRecordSide() *int64 {
	return s.RecordSide
}

func (s *CloudCreateAsrRequest) GetRecordType() *string {
	return s.RecordType
}

func (s *CloudCreateAsrRequest) GetSupportMp3() *string {
	return s.SupportMp3
}

func (s *CloudCreateAsrRequest) GetUniqueId() *string {
	return s.UniqueId
}

func (s *CloudCreateAsrRequest) SetCallType(v int64) *CloudCreateAsrRequest {
	s.CallType = &v
	return s
}

func (s *CloudCreateAsrRequest) SetEnterpriseId(v int64) *CloudCreateAsrRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudCreateAsrRequest) SetMainUniqueId(v string) *CloudCreateAsrRequest {
	s.MainUniqueId = &v
	return s
}

func (s *CloudCreateAsrRequest) SetRecordFile(v string) *CloudCreateAsrRequest {
	s.RecordFile = &v
	return s
}

func (s *CloudCreateAsrRequest) SetRecordSide(v int64) *CloudCreateAsrRequest {
	s.RecordSide = &v
	return s
}

func (s *CloudCreateAsrRequest) SetRecordType(v string) *CloudCreateAsrRequest {
	s.RecordType = &v
	return s
}

func (s *CloudCreateAsrRequest) SetSupportMp3(v string) *CloudCreateAsrRequest {
	s.SupportMp3 = &v
	return s
}

func (s *CloudCreateAsrRequest) SetUniqueId(v string) *CloudCreateAsrRequest {
	s.UniqueId = &v
	return s
}

func (s *CloudCreateAsrRequest) Validate() error {
	return dara.Validate(s)
}
