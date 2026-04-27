// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterSaveFlowConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterSaveFlowConfigResponseBodyData) *ModelRouterSaveFlowConfigResponseBody
	GetData() *ModelRouterSaveFlowConfigResponseBodyData
	SetErrCode(v string) *ModelRouterSaveFlowConfigResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterSaveFlowConfigResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterSaveFlowConfigResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterSaveFlowConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterSaveFlowConfigResponseBody
	GetSuccess() *bool
}

type ModelRouterSaveFlowConfigResponseBody struct {
	Data *ModelRouterSaveFlowConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterSaveFlowConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterSaveFlowConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterSaveFlowConfigResponseBody) GetData() *ModelRouterSaveFlowConfigResponseBodyData {
	return s.Data
}

func (s *ModelRouterSaveFlowConfigResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterSaveFlowConfigResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterSaveFlowConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterSaveFlowConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterSaveFlowConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterSaveFlowConfigResponseBody) SetData(v *ModelRouterSaveFlowConfigResponseBodyData) *ModelRouterSaveFlowConfigResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterSaveFlowConfigResponseBody) SetErrCode(v string) *ModelRouterSaveFlowConfigResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterSaveFlowConfigResponseBody) SetErrMessage(v string) *ModelRouterSaveFlowConfigResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterSaveFlowConfigResponseBody) SetHttpStatusCode(v int32) *ModelRouterSaveFlowConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterSaveFlowConfigResponseBody) SetRequestId(v string) *ModelRouterSaveFlowConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterSaveFlowConfigResponseBody) SetSuccess(v bool) *ModelRouterSaveFlowConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterSaveFlowConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterSaveFlowConfigResponseBodyData struct {
	// example:
	//
	// 2026-04-27T18:28:57.987356+08:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2026-04-27T18:28:57.987356+08:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 6
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 607
	ModelId *int32 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 100
	Rpm *int32 `json:"rpm,omitempty" xml:"rpm,omitempty"`
	// example:
	//
	// true
	SmoothFlowEnabled *bool `json:"smoothFlowEnabled,omitempty" xml:"smoothFlowEnabled,omitempty"`
	// example:
	//
	// 10000
	Tpm *int32 `json:"tpm,omitempty" xml:"tpm,omitempty"`
}

func (s ModelRouterSaveFlowConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterSaveFlowConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterSaveFlowConfigResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ModelRouterSaveFlowConfigResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ModelRouterSaveFlowConfigResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ModelRouterSaveFlowConfigResponseBodyData) GetModelId() *int32 {
	return s.ModelId
}

func (s *ModelRouterSaveFlowConfigResponseBodyData) GetRpm() *int32 {
	return s.Rpm
}

func (s *ModelRouterSaveFlowConfigResponseBodyData) GetSmoothFlowEnabled() *bool {
	return s.SmoothFlowEnabled
}

func (s *ModelRouterSaveFlowConfigResponseBodyData) GetTpm() *int32 {
	return s.Tpm
}

func (s *ModelRouterSaveFlowConfigResponseBodyData) SetGmtCreate(v string) *ModelRouterSaveFlowConfigResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ModelRouterSaveFlowConfigResponseBodyData) SetGmtModified(v string) *ModelRouterSaveFlowConfigResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ModelRouterSaveFlowConfigResponseBodyData) SetId(v int64) *ModelRouterSaveFlowConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *ModelRouterSaveFlowConfigResponseBodyData) SetModelId(v int32) *ModelRouterSaveFlowConfigResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *ModelRouterSaveFlowConfigResponseBodyData) SetRpm(v int32) *ModelRouterSaveFlowConfigResponseBodyData {
	s.Rpm = &v
	return s
}

func (s *ModelRouterSaveFlowConfigResponseBodyData) SetSmoothFlowEnabled(v bool) *ModelRouterSaveFlowConfigResponseBodyData {
	s.SmoothFlowEnabled = &v
	return s
}

func (s *ModelRouterSaveFlowConfigResponseBodyData) SetTpm(v int32) *ModelRouterSaveFlowConfigResponseBodyData {
	s.Tpm = &v
	return s
}

func (s *ModelRouterSaveFlowConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
