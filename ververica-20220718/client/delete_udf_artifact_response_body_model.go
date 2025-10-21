// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUdfArtifactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteUdfArtifactResult) *DeleteUdfArtifactResponseBody
	GetData() *DeleteUdfArtifactResult
	SetErrorCode(v string) *DeleteUdfArtifactResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteUdfArtifactResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *DeleteUdfArtifactResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *DeleteUdfArtifactResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteUdfArtifactResponseBody
	GetSuccess() *bool
}

type DeleteUdfArtifactResponseBody struct {
	Data *DeleteUdfArtifactResult `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-ABCF-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteUdfArtifactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUdfArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUdfArtifactResponseBody) GetData() *DeleteUdfArtifactResult {
	return s.Data
}

func (s *DeleteUdfArtifactResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteUdfArtifactResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteUdfArtifactResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *DeleteUdfArtifactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUdfArtifactResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUdfArtifactResponseBody) SetData(v *DeleteUdfArtifactResult) *DeleteUdfArtifactResponseBody {
	s.Data = v
	return s
}

func (s *DeleteUdfArtifactResponseBody) SetErrorCode(v string) *DeleteUdfArtifactResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteUdfArtifactResponseBody) SetErrorMessage(v string) *DeleteUdfArtifactResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteUdfArtifactResponseBody) SetHttpCode(v int32) *DeleteUdfArtifactResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteUdfArtifactResponseBody) SetRequestId(v string) *DeleteUdfArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUdfArtifactResponseBody) SetSuccess(v bool) *DeleteUdfArtifactResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUdfArtifactResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
