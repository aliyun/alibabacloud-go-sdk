// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEditableNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListEditableNamespaceResponseBodyData) *ListEditableNamespaceResponseBody
	GetData() *ListEditableNamespaceResponseBodyData
	SetHttpCode(v int32) *ListEditableNamespaceResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *ListEditableNamespaceResponseBody
	GetMessage() *string
	SetReason(v string) *ListEditableNamespaceResponseBody
	GetReason() *string
	SetRequestId(v string) *ListEditableNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEditableNamespaceResponseBody
	GetSuccess() *bool
}

type ListEditableNamespaceResponseBody struct {
	Data      *ListEditableNamespaceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	HttpCode  *int32                                 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	Message   *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Reason    *string                                `json:"reason,omitempty" xml:"reason,omitempty"`
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListEditableNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEditableNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *ListEditableNamespaceResponseBody) GetData() *ListEditableNamespaceResponseBodyData {
	return s.Data
}

func (s *ListEditableNamespaceResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListEditableNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEditableNamespaceResponseBody) GetReason() *string {
	return s.Reason
}

func (s *ListEditableNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEditableNamespaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEditableNamespaceResponseBody) SetData(v *ListEditableNamespaceResponseBodyData) *ListEditableNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *ListEditableNamespaceResponseBody) SetHttpCode(v int32) *ListEditableNamespaceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListEditableNamespaceResponseBody) SetMessage(v string) *ListEditableNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *ListEditableNamespaceResponseBody) SetReason(v string) *ListEditableNamespaceResponseBody {
	s.Reason = &v
	return s
}

func (s *ListEditableNamespaceResponseBody) SetRequestId(v string) *ListEditableNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEditableNamespaceResponseBody) SetSuccess(v bool) *ListEditableNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *ListEditableNamespaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEditableNamespaceResponseBodyData struct {
	EditableNamespaces []*EditableNamespace `json:"editableNamespaces,omitempty" xml:"editableNamespaces,omitempty" type:"Repeated"`
	PageIndex          *string              `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	PageSize           *string              `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total              *string              `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListEditableNamespaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEditableNamespaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEditableNamespaceResponseBodyData) GetEditableNamespaces() []*EditableNamespace {
	return s.EditableNamespaces
}

func (s *ListEditableNamespaceResponseBodyData) GetPageIndex() *string {
	return s.PageIndex
}

func (s *ListEditableNamespaceResponseBodyData) GetPageSize() *string {
	return s.PageSize
}

func (s *ListEditableNamespaceResponseBodyData) GetTotal() *string {
	return s.Total
}

func (s *ListEditableNamespaceResponseBodyData) SetEditableNamespaces(v []*EditableNamespace) *ListEditableNamespaceResponseBodyData {
	s.EditableNamespaces = v
	return s
}

func (s *ListEditableNamespaceResponseBodyData) SetPageIndex(v string) *ListEditableNamespaceResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListEditableNamespaceResponseBodyData) SetPageSize(v string) *ListEditableNamespaceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListEditableNamespaceResponseBodyData) SetTotal(v string) *ListEditableNamespaceResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListEditableNamespaceResponseBodyData) Validate() error {
	if s.EditableNamespaces != nil {
		for _, item := range s.EditableNamespaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
