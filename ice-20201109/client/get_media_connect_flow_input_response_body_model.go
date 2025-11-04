// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaConnectFlowInputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *GetMediaConnectFlowInputResponseBodyContent) *GetMediaConnectFlowInputResponseBody
	GetContent() *GetMediaConnectFlowInputResponseBodyContent
	SetDescription(v string) *GetMediaConnectFlowInputResponseBody
	GetDescription() *string
	SetRequestId(v string) *GetMediaConnectFlowInputResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *GetMediaConnectFlowInputResponseBody
	GetRetCode() *int32
}

type GetMediaConnectFlowInputResponseBody struct {
	// The response body.
	Content *GetMediaConnectFlowInputResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The call description.
	//
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D4C231DF-103A-55FF-8D09-E699552457DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned error code. A value of 0 indicates the call is successful.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s GetMediaConnectFlowInputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConnectFlowInputResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaConnectFlowInputResponseBody) GetContent() *GetMediaConnectFlowInputResponseBodyContent {
	return s.Content
}

func (s *GetMediaConnectFlowInputResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetMediaConnectFlowInputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaConnectFlowInputResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *GetMediaConnectFlowInputResponseBody) SetContent(v *GetMediaConnectFlowInputResponseBodyContent) *GetMediaConnectFlowInputResponseBody {
	s.Content = v
	return s
}

func (s *GetMediaConnectFlowInputResponseBody) SetDescription(v string) *GetMediaConnectFlowInputResponseBody {
	s.Description = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBody) SetRequestId(v string) *GetMediaConnectFlowInputResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBody) SetRetCode(v int32) *GetMediaConnectFlowInputResponseBody {
	s.RetCode = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaConnectFlowInputResponseBodyContent struct {
	BackupCidrs         *string `json:"BackupCidrs,omitempty" xml:"BackupCidrs,omitempty"`
	BackupCreateTime    *string `json:"BackupCreateTime,omitempty" xml:"BackupCreateTime,omitempty"`
	BackupInputName     *string `json:"BackupInputName,omitempty" xml:"BackupInputName,omitempty"`
	BackupInputStatus   *string `json:"BackupInputStatus,omitempty" xml:"BackupInputStatus,omitempty"`
	BackupInputUrl      *string `json:"BackupInputUrl,omitempty" xml:"BackupInputUrl,omitempty"`
	BackupMaxBitrate    *int32  `json:"BackupMaxBitrate,omitempty" xml:"BackupMaxBitrate,omitempty"`
	BackupSrtLatency    *int32  `json:"BackupSrtLatency,omitempty" xml:"BackupSrtLatency,omitempty"`
	BackupSrtPassphrase *string `json:"BackupSrtPassphrase,omitempty" xml:"BackupSrtPassphrase,omitempty"`
	BackupSrtPbkeyLen   *int32  `json:"BackupSrtPbkeyLen,omitempty" xml:"BackupSrtPbkeyLen,omitempty"`
	// The IP address whitelist in CIDR format. CIDR blocks are separated with commas (,).
	//
	// example:
	//
	// 10.211.0.0/17
	Cidrs *string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty"`
	// The time when the flow was created.
	//
	// example:
	//
	// 2024-07-18T01:29:24Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The source name.
	//
	// example:
	//
	// AliTestInput
	InputName *string `json:"InputName,omitempty" xml:"InputName,omitempty"`
	// The source type.
	//
	// Valid values:
	//
	// 	- RTMP-PUSH
	//
	// 	- SRT-Caller
	//
	// 	- RTMP-PULL
	//
	// 	- SRT-Listener
	//
	// 	- Flow
	//
	// example:
	//
	// RTMP-PUSH
	InputProtocol *string `json:"InputProtocol,omitempty" xml:"InputProtocol,omitempty"`
	InputStatus   *string `json:"InputStatus,omitempty" xml:"InputStatus,omitempty"`
	// The source URL.
	//
	// example:
	//
	// rtmp://1.2.3.4:1935/live/AliTestInput_8666ec062190f00e263012666319a5be
	InputUrl *string `json:"InputUrl,omitempty" xml:"InputUrl,omitempty"`
	// The maximum bitrate. Unit: bit/s.
	//
	// example:
	//
	// 2000000
	MaxBitrate *int32 `json:"MaxBitrate,omitempty" xml:"MaxBitrate,omitempty"`
	// The ID of the source flow. This parameter is returned when the source type is Flow.
	//
	// example:
	//
	// 05c3adf4-aa0e-421d-a991-48ceae3e642e
	PairFlowId *string `json:"PairFlowId,omitempty" xml:"PairFlowId,omitempty"`
	// The output of the source flow. This parameter is returned when the source type is Flow.
	//
	// example:
	//
	// AliTestOutput
	PairOutputName *string `json:"PairOutputName,omitempty" xml:"PairOutputName,omitempty"`
	// The latency for the SRT stream. Unit: milliseconds. This parameter is returned when the source type is SRT-Listener or SRT-Caller.
	//
	// example:
	//
	// 1000
	SrtLatency *int32 `json:"SrtLatency,omitempty" xml:"SrtLatency,omitempty"`
	// The SRT key. This parameter is returned when the source type is SRT-Listener or SRT-Caller.
	//
	// example:
	//
	// FICUBPX4Q77DYHRF
	SrtPassphrase *string `json:"SrtPassphrase,omitempty" xml:"SrtPassphrase,omitempty"`
	// The encryption key length. This parameter is returned when the source type is SRT-Listener or SRT-Caller.
	//
	// Valid values:
	//
	// 	- 0
	//
	// 	- 16
	//
	// 	- 24
	//
	// 	- 32
	//
	// example:
	//
	// 32
	SrtPbkeyLen *int32 `json:"SrtPbkeyLen,omitempty" xml:"SrtPbkeyLen,omitempty"`
}

func (s GetMediaConnectFlowInputResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConnectFlowInputResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetBackupCidrs() *string {
	return s.BackupCidrs
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetBackupCreateTime() *string {
	return s.BackupCreateTime
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetBackupInputName() *string {
	return s.BackupInputName
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetBackupInputStatus() *string {
	return s.BackupInputStatus
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetBackupInputUrl() *string {
	return s.BackupInputUrl
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetBackupMaxBitrate() *int32 {
	return s.BackupMaxBitrate
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetBackupSrtLatency() *int32 {
	return s.BackupSrtLatency
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetBackupSrtPassphrase() *string {
	return s.BackupSrtPassphrase
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetBackupSrtPbkeyLen() *int32 {
	return s.BackupSrtPbkeyLen
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetCidrs() *string {
	return s.Cidrs
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetInputName() *string {
	return s.InputName
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetInputProtocol() *string {
	return s.InputProtocol
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetInputStatus() *string {
	return s.InputStatus
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetInputUrl() *string {
	return s.InputUrl
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetMaxBitrate() *int32 {
	return s.MaxBitrate
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetPairFlowId() *string {
	return s.PairFlowId
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetPairOutputName() *string {
	return s.PairOutputName
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetSrtLatency() *int32 {
	return s.SrtLatency
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetSrtPassphrase() *string {
	return s.SrtPassphrase
}

func (s *GetMediaConnectFlowInputResponseBodyContent) GetSrtPbkeyLen() *int32 {
	return s.SrtPbkeyLen
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetBackupCidrs(v string) *GetMediaConnectFlowInputResponseBodyContent {
	s.BackupCidrs = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetBackupCreateTime(v string) *GetMediaConnectFlowInputResponseBodyContent {
	s.BackupCreateTime = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetBackupInputName(v string) *GetMediaConnectFlowInputResponseBodyContent {
	s.BackupInputName = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetBackupInputStatus(v string) *GetMediaConnectFlowInputResponseBodyContent {
	s.BackupInputStatus = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetBackupInputUrl(v string) *GetMediaConnectFlowInputResponseBodyContent {
	s.BackupInputUrl = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetBackupMaxBitrate(v int32) *GetMediaConnectFlowInputResponseBodyContent {
	s.BackupMaxBitrate = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetBackupSrtLatency(v int32) *GetMediaConnectFlowInputResponseBodyContent {
	s.BackupSrtLatency = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetBackupSrtPassphrase(v string) *GetMediaConnectFlowInputResponseBodyContent {
	s.BackupSrtPassphrase = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetBackupSrtPbkeyLen(v int32) *GetMediaConnectFlowInputResponseBodyContent {
	s.BackupSrtPbkeyLen = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetCidrs(v string) *GetMediaConnectFlowInputResponseBodyContent {
	s.Cidrs = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetCreateTime(v string) *GetMediaConnectFlowInputResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetInputName(v string) *GetMediaConnectFlowInputResponseBodyContent {
	s.InputName = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetInputProtocol(v string) *GetMediaConnectFlowInputResponseBodyContent {
	s.InputProtocol = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetInputStatus(v string) *GetMediaConnectFlowInputResponseBodyContent {
	s.InputStatus = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetInputUrl(v string) *GetMediaConnectFlowInputResponseBodyContent {
	s.InputUrl = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetMaxBitrate(v int32) *GetMediaConnectFlowInputResponseBodyContent {
	s.MaxBitrate = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetPairFlowId(v string) *GetMediaConnectFlowInputResponseBodyContent {
	s.PairFlowId = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetPairOutputName(v string) *GetMediaConnectFlowInputResponseBodyContent {
	s.PairOutputName = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetSrtLatency(v int32) *GetMediaConnectFlowInputResponseBodyContent {
	s.SrtLatency = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetSrtPassphrase(v string) *GetMediaConnectFlowInputResponseBodyContent {
	s.SrtPassphrase = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) SetSrtPbkeyLen(v int32) *GetMediaConnectFlowInputResponseBodyContent {
	s.SrtPbkeyLen = &v
	return s
}

func (s *GetMediaConnectFlowInputResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
