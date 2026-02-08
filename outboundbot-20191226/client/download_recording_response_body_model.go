// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadRecordingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DownloadRecordingResponseBody
	GetCode() *string
	SetDownloadParams(v *DownloadRecordingResponseBodyDownloadParams) *DownloadRecordingResponseBody
	GetDownloadParams() *DownloadRecordingResponseBodyDownloadParams
	SetHttpStatusCode(v int32) *DownloadRecordingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DownloadRecordingResponseBody
	GetMessage() *string
	SetRequestId(v string) *DownloadRecordingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DownloadRecordingResponseBody
	GetSuccess() *bool
}

type DownloadRecordingResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Document download URL
	DownloadParams *DownloadRecordingResponseBodyDownloadParams `json:"DownloadParams,omitempty" xml:"DownloadParams,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API prompt message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DownloadRecordingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadRecordingResponseBody) GetCode() *string {
	return s.Code
}

func (s *DownloadRecordingResponseBody) GetDownloadParams() *DownloadRecordingResponseBodyDownloadParams {
	return s.DownloadParams
}

func (s *DownloadRecordingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DownloadRecordingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DownloadRecordingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadRecordingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DownloadRecordingResponseBody) SetCode(v string) *DownloadRecordingResponseBody {
	s.Code = &v
	return s
}

func (s *DownloadRecordingResponseBody) SetDownloadParams(v *DownloadRecordingResponseBodyDownloadParams) *DownloadRecordingResponseBody {
	s.DownloadParams = v
	return s
}

func (s *DownloadRecordingResponseBody) SetHttpStatusCode(v int32) *DownloadRecordingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DownloadRecordingResponseBody) SetMessage(v string) *DownloadRecordingResponseBody {
	s.Message = &v
	return s
}

func (s *DownloadRecordingResponseBody) SetRequestId(v string) *DownloadRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadRecordingResponseBody) SetSuccess(v bool) *DownloadRecordingResponseBody {
	s.Success = &v
	return s
}

func (s *DownloadRecordingResponseBody) Validate() error {
	if s.DownloadParams != nil {
		if err := s.DownloadParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DownloadRecordingResponseBodyDownloadParams struct {
	// Recording file name, usually a UUID
	//
	// example:
	//
	// 281eb174-3865-41c1-9274-7b6813edadab.wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// A URL pointing to the recording file for HTTP download
	//
	// example:
	//
	// http://tiangong-staging.oss-cn-shanghai.aliyuncs.com/record/281eb174-3865-41c1-9274-7b6813edadab.wav?Expires=1578624046&OSSAccessKeyId=LTAI****cqw&Signature=dL2dxWS6VcdZrvG9xOMOBMSP3Fg%3D
	SignatureUrl *string `json:"SignatureUrl,omitempty" xml:"SignatureUrl,omitempty"`
	// List of segmented recordings, including file names and file URLs
	//
	// example:
	//
	// [{"fileName":"10a17c447415424c99491884abe27d8a-1.wav","filePath":"https://ssml-test.oss-cn-shanghai.aliyuncs.com/7253/voiceSlice/10a17c447415424c99491884abe27d8a/10a17c447415424c99491884abe27d8a-1.wav?Expires=1686645470&OSSAccessKeyId=LTA*********kr8M9&Signature=V23OhiV5yIOoouriu6SiWkO9h8E%3D"}]
	VoiceSliceRecordingListJson *string `json:"VoiceSliceRecordingListJson,omitempty" xml:"VoiceSliceRecordingListJson,omitempty"`
}

func (s DownloadRecordingResponseBodyDownloadParams) String() string {
	return dara.Prettify(s)
}

func (s DownloadRecordingResponseBodyDownloadParams) GoString() string {
	return s.String()
}

func (s *DownloadRecordingResponseBodyDownloadParams) GetFileName() *string {
	return s.FileName
}

func (s *DownloadRecordingResponseBodyDownloadParams) GetSignatureUrl() *string {
	return s.SignatureUrl
}

func (s *DownloadRecordingResponseBodyDownloadParams) GetVoiceSliceRecordingListJson() *string {
	return s.VoiceSliceRecordingListJson
}

func (s *DownloadRecordingResponseBodyDownloadParams) SetFileName(v string) *DownloadRecordingResponseBodyDownloadParams {
	s.FileName = &v
	return s
}

func (s *DownloadRecordingResponseBodyDownloadParams) SetSignatureUrl(v string) *DownloadRecordingResponseBodyDownloadParams {
	s.SignatureUrl = &v
	return s
}

func (s *DownloadRecordingResponseBodyDownloadParams) SetVoiceSliceRecordingListJson(v string) *DownloadRecordingResponseBodyDownloadParams {
	s.VoiceSliceRecordingListJson = &v
	return s
}

func (s *DownloadRecordingResponseBodyDownloadParams) Validate() error {
	return dara.Validate(s)
}
