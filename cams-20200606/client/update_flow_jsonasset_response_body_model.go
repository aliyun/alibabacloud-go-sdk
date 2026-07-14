// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowJSONAssetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateFlowJSONAssetResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateFlowJSONAssetResponseBody
	GetCode() *string
	SetData(v *UpdateFlowJSONAssetResponseBodyData) *UpdateFlowJSONAssetResponseBody
	GetData() *UpdateFlowJSONAssetResponseBodyData
	SetMessage(v string) *UpdateFlowJSONAssetResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateFlowJSONAssetResponseBody
	GetRequestId() *string
}

type UpdateFlowJSONAssetResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. OK indicates success.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *UpdateFlowJSONAssetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFlowJSONAssetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowJSONAssetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlowJSONAssetResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateFlowJSONAssetResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateFlowJSONAssetResponseBody) GetData() *UpdateFlowJSONAssetResponseBodyData {
	return s.Data
}

func (s *UpdateFlowJSONAssetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateFlowJSONAssetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFlowJSONAssetResponseBody) SetAccessDeniedDetail(v string) *UpdateFlowJSONAssetResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateFlowJSONAssetResponseBody) SetCode(v string) *UpdateFlowJSONAssetResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFlowJSONAssetResponseBody) SetData(v *UpdateFlowJSONAssetResponseBodyData) *UpdateFlowJSONAssetResponseBody {
	s.Data = v
	return s
}

func (s *UpdateFlowJSONAssetResponseBody) SetMessage(v string) *UpdateFlowJSONAssetResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFlowJSONAssetResponseBody) SetRequestId(v string) *UpdateFlowJSONAssetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlowJSONAssetResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateFlowJSONAssetResponseBodyData struct {
	// Flow ID。
	//
	// example:
	//
	// 92675332812643****
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s UpdateFlowJSONAssetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowJSONAssetResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateFlowJSONAssetResponseBodyData) GetFlowId() *string {
	return s.FlowId
}

func (s *UpdateFlowJSONAssetResponseBodyData) SetFlowId(v string) *UpdateFlowJSONAssetResponseBodyData {
	s.FlowId = &v
	return s
}

func (s *UpdateFlowJSONAssetResponseBodyData) Validate() error {
	return dara.Validate(s)
}
