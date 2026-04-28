// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetRecordUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallType(v int64) *CloudGetRecordUrlRequest
	GetCallType() *int64
	SetDownload(v int64) *CloudGetRecordUrlRequest
	GetDownload() *int64
	SetEnterpriseId(v int64) *CloudGetRecordUrlRequest
	GetEnterpriseId() *int64
	SetRecordFile(v string) *CloudGetRecordUrlRequest
	GetRecordFile() *string
	SetRecordFormat(v int64) *CloudGetRecordUrlRequest
	GetRecordFormat() *int64
	SetRecordSide(v int64) *CloudGetRecordUrlRequest
	GetRecordSide() *int64
	SetRecordType(v string) *CloudGetRecordUrlRequest
	GetRecordType() *string
}

type CloudGetRecordUrlRequest struct {
	// 呼叫类型；说明：开启分线录音时，获取客户侧或座席侧录音需要，recordFormat=1时有效，recordFormat=0时忽略。取值范围：1,2,4,5（数字含义：呼入,webcall,预览外呼,预测外呼）
	//
	// example:
	//
	// 4
	CallType *int64 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 是否下载；１为下载，空或０表示试听
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
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 录音文件名；如7000101-20160815140025-01087120766-01087120766--record-sip-1-1471240825.87
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000101-20160815140025-01087120766-01087120766--record-sip-1-1471240825.87
	RecordFile *string `json:"RecordFile,omitempty" xml:"RecordFile,omitempty"`
	// 录音文件类型；取值说明：0为mp3，1为wav，默认为mp3类型
	//
	// example:
	//
	// 1
	RecordFormat *int64 `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty"`
	// 分线录音录音侧；说明：开启分线录音时，获取客户侧或座席侧录音需要，recordFormat=1时有效，recordFormat=0时忽略。取值范围：1,2 (数字含义：客户侧，座席侧)recordSide不为空时，callType必选
	//
	// example:
	//
	// 2
	RecordSide *int64 `json:"RecordSide,omitempty" xml:"RecordSide,omitempty"`
	// 录音类型，不填可根据 recordFile 解析；record:录音voicemail:留言tsi:彩铃、当开启号码录音状态识别，发起预览外呼，客户号码是手机且客户未接听时返回该地址 rasr:语音机器人客户侧录音 transfer:转接consult:咨询threeway:三方
	//
	// example:
	//
	// record
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
}

func (s CloudGetRecordUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudGetRecordUrlRequest) GoString() string {
	return s.String()
}

func (s *CloudGetRecordUrlRequest) GetCallType() *int64 {
	return s.CallType
}

func (s *CloudGetRecordUrlRequest) GetDownload() *int64 {
	return s.Download
}

func (s *CloudGetRecordUrlRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudGetRecordUrlRequest) GetRecordFile() *string {
	return s.RecordFile
}

func (s *CloudGetRecordUrlRequest) GetRecordFormat() *int64 {
	return s.RecordFormat
}

func (s *CloudGetRecordUrlRequest) GetRecordSide() *int64 {
	return s.RecordSide
}

func (s *CloudGetRecordUrlRequest) GetRecordType() *string {
	return s.RecordType
}

func (s *CloudGetRecordUrlRequest) SetCallType(v int64) *CloudGetRecordUrlRequest {
	s.CallType = &v
	return s
}

func (s *CloudGetRecordUrlRequest) SetDownload(v int64) *CloudGetRecordUrlRequest {
	s.Download = &v
	return s
}

func (s *CloudGetRecordUrlRequest) SetEnterpriseId(v int64) *CloudGetRecordUrlRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudGetRecordUrlRequest) SetRecordFile(v string) *CloudGetRecordUrlRequest {
	s.RecordFile = &v
	return s
}

func (s *CloudGetRecordUrlRequest) SetRecordFormat(v int64) *CloudGetRecordUrlRequest {
	s.RecordFormat = &v
	return s
}

func (s *CloudGetRecordUrlRequest) SetRecordSide(v int64) *CloudGetRecordUrlRequest {
	s.RecordSide = &v
	return s
}

func (s *CloudGetRecordUrlRequest) SetRecordType(v string) *CloudGetRecordUrlRequest {
	s.RecordType = &v
	return s
}

func (s *CloudGetRecordUrlRequest) Validate() error {
	return dara.Validate(s)
}
