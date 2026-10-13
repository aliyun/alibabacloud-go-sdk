// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiPathFieldTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListApiPathFieldTemplatesResponseBody
	GetCode() *string
	SetData(v *ListApiPathFieldTemplatesResponseBodyData) *ListApiPathFieldTemplatesResponseBody
	GetData() *ListApiPathFieldTemplatesResponseBodyData
	SetMessage(v string) *ListApiPathFieldTemplatesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListApiPathFieldTemplatesResponseBody
	GetRequestId() *string
}

type ListApiPathFieldTemplatesResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListApiPathFieldTemplatesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListApiPathFieldTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApiPathFieldTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListApiPathFieldTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListApiPathFieldTemplatesResponseBody) GetData() *ListApiPathFieldTemplatesResponseBodyData {
	return s.Data
}

func (s *ListApiPathFieldTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApiPathFieldTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApiPathFieldTemplatesResponseBody) SetCode(v string) *ListApiPathFieldTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListApiPathFieldTemplatesResponseBody) SetData(v *ListApiPathFieldTemplatesResponseBodyData) *ListApiPathFieldTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListApiPathFieldTemplatesResponseBody) SetMessage(v string) *ListApiPathFieldTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListApiPathFieldTemplatesResponseBody) SetRequestId(v string) *ListApiPathFieldTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApiPathFieldTemplatesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApiPathFieldTemplatesResponseBodyData struct {
	// The list of field templates. Each element in the items list is a field template record that contains the templateId, path, and protocol.
	Items []*ListApiPathFieldTemplatesResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListApiPathFieldTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListApiPathFieldTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListApiPathFieldTemplatesResponseBodyData) GetItems() []*ListApiPathFieldTemplatesResponseBodyDataItems {
	return s.Items
}

func (s *ListApiPathFieldTemplatesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApiPathFieldTemplatesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApiPathFieldTemplatesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListApiPathFieldTemplatesResponseBodyData) SetItems(v []*ListApiPathFieldTemplatesResponseBodyDataItems) *ListApiPathFieldTemplatesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListApiPathFieldTemplatesResponseBodyData) SetPageNumber(v int32) *ListApiPathFieldTemplatesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListApiPathFieldTemplatesResponseBodyData) SetPageSize(v int32) *ListApiPathFieldTemplatesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListApiPathFieldTemplatesResponseBodyData) SetTotalSize(v int32) *ListApiPathFieldTemplatesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListApiPathFieldTemplatesResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApiPathFieldTemplatesResponseBodyDataItems struct {
	// The request path corresponding to the field template.
	//
	// example:
	//
	// /v1/chat/completions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The API protocol applicable to the field template.
	//
	// example:
	//
	// openai
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The ID of the field template.
	//
	// example:
	//
	// atpl-ec5df7baf496bfac
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s ListApiPathFieldTemplatesResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListApiPathFieldTemplatesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListApiPathFieldTemplatesResponseBodyDataItems) GetPath() *string {
	return s.Path
}

func (s *ListApiPathFieldTemplatesResponseBodyDataItems) GetProtocol() *string {
	return s.Protocol
}

func (s *ListApiPathFieldTemplatesResponseBodyDataItems) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListApiPathFieldTemplatesResponseBodyDataItems) SetPath(v string) *ListApiPathFieldTemplatesResponseBodyDataItems {
	s.Path = &v
	return s
}

func (s *ListApiPathFieldTemplatesResponseBodyDataItems) SetProtocol(v string) *ListApiPathFieldTemplatesResponseBodyDataItems {
	s.Protocol = &v
	return s
}

func (s *ListApiPathFieldTemplatesResponseBodyDataItems) SetTemplateId(v string) *ListApiPathFieldTemplatesResponseBodyDataItems {
	s.TemplateId = &v
	return s
}

func (s *ListApiPathFieldTemplatesResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
