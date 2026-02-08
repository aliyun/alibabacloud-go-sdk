// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadScriptRecordingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DownloadScriptRecordingResponseBody
	GetCode() *string
	SetDownloadParams(v *DownloadScriptRecordingResponseBodyDownloadParams) *DownloadScriptRecordingResponseBody
	GetDownloadParams() *DownloadScriptRecordingResponseBodyDownloadParams
	SetHttpStatusCode(v int32) *DownloadScriptRecordingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DownloadScriptRecordingResponseBody
	GetMessage() *string
	SetRequestId(v string) *DownloadScriptRecordingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DownloadScriptRecordingResponseBody
	GetSuccess() *bool
}

type DownloadScriptRecordingResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Parameters
	DownloadParams *DownloadScriptRecordingResponseBodyDownloadParams `json:"DownloadParams,omitempty" xml:"DownloadParams,omitempty" type:"Struct"`
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

func (s DownloadScriptRecordingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadScriptRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadScriptRecordingResponseBody) GetCode() *string {
	return s.Code
}

func (s *DownloadScriptRecordingResponseBody) GetDownloadParams() *DownloadScriptRecordingResponseBodyDownloadParams {
	return s.DownloadParams
}

func (s *DownloadScriptRecordingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DownloadScriptRecordingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DownloadScriptRecordingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadScriptRecordingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DownloadScriptRecordingResponseBody) SetCode(v string) *DownloadScriptRecordingResponseBody {
	s.Code = &v
	return s
}

func (s *DownloadScriptRecordingResponseBody) SetDownloadParams(v *DownloadScriptRecordingResponseBodyDownloadParams) *DownloadScriptRecordingResponseBody {
	s.DownloadParams = v
	return s
}

func (s *DownloadScriptRecordingResponseBody) SetHttpStatusCode(v int32) *DownloadScriptRecordingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DownloadScriptRecordingResponseBody) SetMessage(v string) *DownloadScriptRecordingResponseBody {
	s.Message = &v
	return s
}

func (s *DownloadScriptRecordingResponseBody) SetRequestId(v string) *DownloadScriptRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadScriptRecordingResponseBody) SetSuccess(v bool) *DownloadScriptRecordingResponseBody {
	s.Success = &v
	return s
}

func (s *DownloadScriptRecordingResponseBody) Validate() error {
	if s.DownloadParams != nil {
		if err := s.DownloadParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DownloadScriptRecordingResponseBodyDownloadParams struct {
	// Recording file name
	//
	// example:
	//
	// 281eb174-3865-41c1-9274-7b6813edadab.wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// A URL pointing to the recording file. You can use this URL for playback or download.
	//
	// example:
	//
	// http://tiangong-staging.oss-cn-shanghai.aliyuncs.com/record/281eb174-3865-41c1-9274-7b6813edadab.wav?Expires=1578624046&OSSAccessKeyId=LTAI****cqw&Signature=dL2dxWS6VcdZrvG9xOMOBMSP3Fg%3D
	SignatureUrl *string `json:"SignatureUrl,omitempty" xml:"SignatureUrl,omitempty"`
}

func (s DownloadScriptRecordingResponseBodyDownloadParams) String() string {
	return dara.Prettify(s)
}

func (s DownloadScriptRecordingResponseBodyDownloadParams) GoString() string {
	return s.String()
}

func (s *DownloadScriptRecordingResponseBodyDownloadParams) GetFileName() *string {
	return s.FileName
}

func (s *DownloadScriptRecordingResponseBodyDownloadParams) GetSignatureUrl() *string {
	return s.SignatureUrl
}

func (s *DownloadScriptRecordingResponseBodyDownloadParams) SetFileName(v string) *DownloadScriptRecordingResponseBodyDownloadParams {
	s.FileName = &v
	return s
}

func (s *DownloadScriptRecordingResponseBodyDownloadParams) SetSignatureUrl(v string) *DownloadScriptRecordingResponseBodyDownloadParams {
	s.SignatureUrl = &v
	return s
}

func (s *DownloadScriptRecordingResponseBodyDownloadParams) Validate() error {
	return dara.Validate(s)
}
