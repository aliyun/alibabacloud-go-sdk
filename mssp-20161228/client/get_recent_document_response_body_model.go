// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecentDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRecentDocumentResponseBody
	GetCode() *string
	SetData(v []*GetRecentDocumentResponseBodyData) *GetRecentDocumentResponseBody
	GetData() []*GetRecentDocumentResponseBodyData
	SetHttpStatusCode(v int32) *GetRecentDocumentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetRecentDocumentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRecentDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRecentDocumentResponseBody
	GetSuccess() *bool
}

type GetRecentDocumentResponseBody struct {
	// Interface response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data []*GetRecentDocumentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4916FA8D-F294-518D-B373-8B59D63CAB19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRecentDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRecentDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecentDocumentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRecentDocumentResponseBody) GetData() []*GetRecentDocumentResponseBodyData {
	return s.Data
}

func (s *GetRecentDocumentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetRecentDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRecentDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRecentDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRecentDocumentResponseBody) SetCode(v string) *GetRecentDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *GetRecentDocumentResponseBody) SetData(v []*GetRecentDocumentResponseBodyData) *GetRecentDocumentResponseBody {
	s.Data = v
	return s
}

func (s *GetRecentDocumentResponseBody) SetHttpStatusCode(v int32) *GetRecentDocumentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRecentDocumentResponseBody) SetMessage(v string) *GetRecentDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *GetRecentDocumentResponseBody) SetRequestId(v string) *GetRecentDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecentDocumentResponseBody) SetSuccess(v bool) *GetRecentDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *GetRecentDocumentResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRecentDocumentResponseBodyData struct {
	// Primary key ID.
	//
	// example:
	//
	// 360491
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Document name
	//
	// example:
	//
	// 文档名称测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Upload time.
	//
	// example:
	//
	// 2023-03-20 14:30:38
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s GetRecentDocumentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRecentDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRecentDocumentResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetRecentDocumentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetRecentDocumentResponseBodyData) GetUploadTime() *string {
	return s.UploadTime
}

func (s *GetRecentDocumentResponseBodyData) SetId(v int64) *GetRecentDocumentResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetRecentDocumentResponseBodyData) SetName(v string) *GetRecentDocumentResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetRecentDocumentResponseBodyData) SetUploadTime(v string) *GetRecentDocumentResponseBodyData {
	s.UploadTime = &v
	return s
}

func (s *GetRecentDocumentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
