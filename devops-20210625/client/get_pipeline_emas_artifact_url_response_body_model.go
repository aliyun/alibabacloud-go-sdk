// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineEmasArtifactUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetPipelineEmasArtifactUrlResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetPipelineEmasArtifactUrlResponseBody
	GetErrorMessage() *string
	SetFileUrl(v string) *GetPipelineEmasArtifactUrlResponseBody
	GetFileUrl() *string
	SetRequestId(v string) *GetPipelineEmasArtifactUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPipelineEmasArtifactUrlResponseBody
	GetSuccess() *bool
}

type GetPipelineEmasArtifactUrlResponseBody struct {
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
	// http://aliyun.com
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

func (s GetPipelineEmasArtifactUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineEmasArtifactUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineEmasArtifactUrlResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetPipelineEmasArtifactUrlResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetPipelineEmasArtifactUrlResponseBody) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetPipelineEmasArtifactUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPipelineEmasArtifactUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPipelineEmasArtifactUrlResponseBody) SetErrorCode(v string) *GetPipelineEmasArtifactUrlResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineEmasArtifactUrlResponseBody) SetErrorMessage(v string) *GetPipelineEmasArtifactUrlResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineEmasArtifactUrlResponseBody) SetFileUrl(v string) *GetPipelineEmasArtifactUrlResponseBody {
	s.FileUrl = &v
	return s
}

func (s *GetPipelineEmasArtifactUrlResponseBody) SetRequestId(v string) *GetPipelineEmasArtifactUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineEmasArtifactUrlResponseBody) SetSuccess(v bool) *GetPipelineEmasArtifactUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetPipelineEmasArtifactUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
