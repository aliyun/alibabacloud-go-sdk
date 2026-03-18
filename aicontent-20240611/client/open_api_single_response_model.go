// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenApiSingleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OpenApiSingleResponseData) *OpenApiSingleResponse
	GetData() *OpenApiSingleResponseData
	SetErrCode(v string) *OpenApiSingleResponse
	GetErrCode() *string
	SetErrMessage(v string) *OpenApiSingleResponse
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *OpenApiSingleResponse
	GetHttpStatusCode() *int32
	SetRequestId(v string) *OpenApiSingleResponse
	GetRequestId() *string
	SetSuccess(v bool) *OpenApiSingleResponse
	GetSuccess() *bool
}

type OpenApiSingleResponse struct {
	// example:
	//
	// []
	Data       *OpenApiSingleResponseData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode    *string                    `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string                    `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OpenApiSingleResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenApiSingleResponse) GoString() string {
	return s.String()
}

func (s *OpenApiSingleResponse) GetData() *OpenApiSingleResponseData {
	return s.Data
}

func (s *OpenApiSingleResponse) GetErrCode() *string {
	return s.ErrCode
}

func (s *OpenApiSingleResponse) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *OpenApiSingleResponse) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *OpenApiSingleResponse) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenApiSingleResponse) GetSuccess() *bool {
	return s.Success
}

func (s *OpenApiSingleResponse) SetData(v *OpenApiSingleResponseData) *OpenApiSingleResponse {
	s.Data = v
	return s
}

func (s *OpenApiSingleResponse) SetErrCode(v string) *OpenApiSingleResponse {
	s.ErrCode = &v
	return s
}

func (s *OpenApiSingleResponse) SetErrMessage(v string) *OpenApiSingleResponse {
	s.ErrMessage = &v
	return s
}

func (s *OpenApiSingleResponse) SetHttpStatusCode(v int32) *OpenApiSingleResponse {
	s.HttpStatusCode = &v
	return s
}

func (s *OpenApiSingleResponse) SetRequestId(v string) *OpenApiSingleResponse {
	s.RequestId = &v
	return s
}

func (s *OpenApiSingleResponse) SetSuccess(v bool) *OpenApiSingleResponse {
	s.Success = &v
	return s
}

func (s *OpenApiSingleResponse) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OpenApiSingleResponseData struct {
	// example:
	//
	// FINISHED
	ModelTrainStatus *string `json:"modelTrainStatus,omitempty" xml:"modelTrainStatus,omitempty"`
}

func (s OpenApiSingleResponseData) String() string {
	return dara.Prettify(s)
}

func (s OpenApiSingleResponseData) GoString() string {
	return s.String()
}

func (s *OpenApiSingleResponseData) GetModelTrainStatus() *string {
	return s.ModelTrainStatus
}

func (s *OpenApiSingleResponseData) SetModelTrainStatus(v string) *OpenApiSingleResponseData {
	s.ModelTrainStatus = &v
	return s
}

func (s *OpenApiSingleResponseData) Validate() error {
	return dara.Validate(s)
}
