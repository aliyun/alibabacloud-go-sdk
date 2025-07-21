// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLabelsResponseBody
	GetCode() *string
	SetData(v []*ListLabelsResponseBodyData) *ListLabelsResponseBody
	GetData() []*ListLabelsResponseBodyData
	SetMessage(v string) *ListLabelsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListLabelsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListLabelsResponseBody
	GetRequestId() *string
}

type ListLabelsResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListLabelsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLabelsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLabelsResponseBody) GetData() []*ListLabelsResponseBodyData {
	return s.Data
}

func (s *ListLabelsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLabelsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLabelsResponseBody) SetCode(v string) *ListLabelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListLabelsResponseBody) SetData(v []*ListLabelsResponseBodyData) *ListLabelsResponseBody {
	s.Data = v
	return s
}

func (s *ListLabelsResponseBody) SetMessage(v string) *ListLabelsResponseBody {
	s.Message = &v
	return s
}

func (s *ListLabelsResponseBody) SetNextToken(v string) *ListLabelsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListLabelsResponseBody) SetRequestId(v string) *ListLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLabelsResponseBodyData struct {
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LabelId     *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListLabelsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLabelsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLabelsResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *ListLabelsResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListLabelsResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListLabelsResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListLabelsResponseBodyData) GetLabelId() *string {
	return s.LabelId
}

func (s *ListLabelsResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *ListLabelsResponseBodyData) SetContent(v string) *ListLabelsResponseBodyData {
	s.Content = &v
	return s
}

func (s *ListLabelsResponseBodyData) SetGmtCreate(v int64) *ListLabelsResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListLabelsResponseBodyData) SetGmtModified(v int64) *ListLabelsResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListLabelsResponseBodyData) SetId(v int64) *ListLabelsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListLabelsResponseBodyData) SetLabelId(v string) *ListLabelsResponseBodyData {
	s.LabelId = &v
	return s
}

func (s *ListLabelsResponseBodyData) SetTenantId(v string) *ListLabelsResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *ListLabelsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
