// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineArtifactUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetPipelineArtifactUrlResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetPipelineArtifactUrlResponseBody
	GetErrorMessage() *string
	SetFileUrl(v string) *GetPipelineArtifactUrlResponseBody
	GetFileUrl() *string
	SetRequestId(v string) *GetPipelineArtifactUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPipelineArtifactUrlResponseBody
	GetSuccess() *bool
}

type GetPipelineArtifactUrlResponseBody struct {
	// example:
	//
	// ”“
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// http://aliyun.com/
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPipelineArtifactUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineArtifactUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineArtifactUrlResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetPipelineArtifactUrlResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetPipelineArtifactUrlResponseBody) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetPipelineArtifactUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPipelineArtifactUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPipelineArtifactUrlResponseBody) SetErrorCode(v string) *GetPipelineArtifactUrlResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineArtifactUrlResponseBody) SetErrorMessage(v string) *GetPipelineArtifactUrlResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineArtifactUrlResponseBody) SetFileUrl(v string) *GetPipelineArtifactUrlResponseBody {
	s.FileUrl = &v
	return s
}

func (s *GetPipelineArtifactUrlResponseBody) SetRequestId(v string) *GetPipelineArtifactUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineArtifactUrlResponseBody) SetSuccess(v bool) *GetPipelineArtifactUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetPipelineArtifactUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
