// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnnotationMissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAnnotationMissionResponseBody
	GetCode() *string
	SetData(v *CreateAnnotationMissionResponseBodyData) *CreateAnnotationMissionResponseBody
	GetData() *CreateAnnotationMissionResponseBodyData
	SetHttpStatusCode(v int32) *CreateAnnotationMissionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateAnnotationMissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAnnotationMissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAnnotationMissionResponseBody
	GetSuccess() *bool
}

type CreateAnnotationMissionResponseBody struct {
	// The status code of the API.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateAnnotationMissionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Additional information. If the request is successful, \\`success\\` is returned. If the request fails, a specific error code is returned.
	//
	// example:
	//
	// 执行失败：查找元素或操作超时,outTaskId:trademark_regnew_public_cn-uax33hol2uu,flowId:null,nodeUUID:541b6c71-8cce-4ab0-af84-713ca9741821
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. \\`true\\`: The call was successful. \\`false\\`: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAnnotationMissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAnnotationMissionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAnnotationMissionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAnnotationMissionResponseBody) GetData() *CreateAnnotationMissionResponseBodyData {
	return s.Data
}

func (s *CreateAnnotationMissionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateAnnotationMissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAnnotationMissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAnnotationMissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAnnotationMissionResponseBody) SetCode(v string) *CreateAnnotationMissionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAnnotationMissionResponseBody) SetData(v *CreateAnnotationMissionResponseBodyData) *CreateAnnotationMissionResponseBody {
	s.Data = v
	return s
}

func (s *CreateAnnotationMissionResponseBody) SetHttpStatusCode(v int32) *CreateAnnotationMissionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAnnotationMissionResponseBody) SetMessage(v string) *CreateAnnotationMissionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAnnotationMissionResponseBody) SetRequestId(v string) *CreateAnnotationMissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAnnotationMissionResponseBody) SetSuccess(v bool) *CreateAnnotationMissionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAnnotationMissionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAnnotationMissionResponseBodyData struct {
	// Additional information. If the request is successful, \\`success\\` is returned. If the request fails, a specific error code is returned.
	//
	// example:
	//
	// 执行失败：查找元素或操作超时,outTaskId:trademark_regnew_public_cn-uax33hol2uu,flowId:null,nodeUUID:541b6c71-8cce-4ab0-af84-713ca9741821
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAnnotationMissionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAnnotationMissionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAnnotationMissionResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *CreateAnnotationMissionResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAnnotationMissionResponseBodyData) SetMessage(v string) *CreateAnnotationMissionResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreateAnnotationMissionResponseBodyData) SetSuccess(v bool) *CreateAnnotationMissionResponseBodyData {
	s.Success = &v
	return s
}

func (s *CreateAnnotationMissionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
