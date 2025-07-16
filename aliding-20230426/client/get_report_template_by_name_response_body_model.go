// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportTemplateByNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultReceivedConvs(v []*GetReportTemplateByNameResponseBodyDefaultReceivedConvs) *GetReportTemplateByNameResponseBody
	GetDefaultReceivedConvs() []*GetReportTemplateByNameResponseBodyDefaultReceivedConvs
	SetDefaultReceivers(v []*GetReportTemplateByNameResponseBodyDefaultReceivers) *GetReportTemplateByNameResponseBody
	GetDefaultReceivers() []*GetReportTemplateByNameResponseBodyDefaultReceivers
	SetFields(v []*GetReportTemplateByNameResponseBodyFields) *GetReportTemplateByNameResponseBody
	GetFields() []*GetReportTemplateByNameResponseBodyFields
	SetId(v string) *GetReportTemplateByNameResponseBody
	GetId() *string
	SetName(v string) *GetReportTemplateByNameResponseBody
	GetName() *string
	SetRequestId(v string) *GetReportTemplateByNameResponseBody
	GetRequestId() *string
	SetUserName(v string) *GetReportTemplateByNameResponseBody
	GetUserName() *string
	SetUserid(v string) *GetReportTemplateByNameResponseBody
	GetUserid() *string
}

type GetReportTemplateByNameResponseBody struct {
	DefaultReceivedConvs []*GetReportTemplateByNameResponseBodyDefaultReceivedConvs `json:"defaultReceivedConvs,omitempty" xml:"defaultReceivedConvs,omitempty" type:"Repeated"`
	DefaultReceivers     []*GetReportTemplateByNameResponseBodyDefaultReceivers     `json:"defaultReceivers,omitempty" xml:"defaultReceivers,omitempty" type:"Repeated"`
	Fields               []*GetReportTemplateByNameResponseBodyFields               `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// example:
	//
	// 11111
	Id   *string `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	UserName  *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// example:
	//
	// 1234
	Userid *string `json:"userid,omitempty" xml:"userid,omitempty"`
}

func (s GetReportTemplateByNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetReportTemplateByNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetReportTemplateByNameResponseBody) GetDefaultReceivedConvs() []*GetReportTemplateByNameResponseBodyDefaultReceivedConvs {
	return s.DefaultReceivedConvs
}

func (s *GetReportTemplateByNameResponseBody) GetDefaultReceivers() []*GetReportTemplateByNameResponseBodyDefaultReceivers {
	return s.DefaultReceivers
}

func (s *GetReportTemplateByNameResponseBody) GetFields() []*GetReportTemplateByNameResponseBodyFields {
	return s.Fields
}

func (s *GetReportTemplateByNameResponseBody) GetId() *string {
	return s.Id
}

func (s *GetReportTemplateByNameResponseBody) GetName() *string {
	return s.Name
}

func (s *GetReportTemplateByNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetReportTemplateByNameResponseBody) GetUserName() *string {
	return s.UserName
}

func (s *GetReportTemplateByNameResponseBody) GetUserid() *string {
	return s.Userid
}

func (s *GetReportTemplateByNameResponseBody) SetDefaultReceivedConvs(v []*GetReportTemplateByNameResponseBodyDefaultReceivedConvs) *GetReportTemplateByNameResponseBody {
	s.DefaultReceivedConvs = v
	return s
}

func (s *GetReportTemplateByNameResponseBody) SetDefaultReceivers(v []*GetReportTemplateByNameResponseBodyDefaultReceivers) *GetReportTemplateByNameResponseBody {
	s.DefaultReceivers = v
	return s
}

func (s *GetReportTemplateByNameResponseBody) SetFields(v []*GetReportTemplateByNameResponseBodyFields) *GetReportTemplateByNameResponseBody {
	s.Fields = v
	return s
}

func (s *GetReportTemplateByNameResponseBody) SetId(v string) *GetReportTemplateByNameResponseBody {
	s.Id = &v
	return s
}

func (s *GetReportTemplateByNameResponseBody) SetName(v string) *GetReportTemplateByNameResponseBody {
	s.Name = &v
	return s
}

func (s *GetReportTemplateByNameResponseBody) SetRequestId(v string) *GetReportTemplateByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetReportTemplateByNameResponseBody) SetUserName(v string) *GetReportTemplateByNameResponseBody {
	s.UserName = &v
	return s
}

func (s *GetReportTemplateByNameResponseBody) SetUserid(v string) *GetReportTemplateByNameResponseBody {
	s.Userid = &v
	return s
}

func (s *GetReportTemplateByNameResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetReportTemplateByNameResponseBodyDefaultReceivedConvs struct {
	// example:
	//
	// cid12334##2341
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetReportTemplateByNameResponseBodyDefaultReceivedConvs) String() string {
	return dara.Prettify(s)
}

func (s GetReportTemplateByNameResponseBodyDefaultReceivedConvs) GoString() string {
	return s.String()
}

func (s *GetReportTemplateByNameResponseBodyDefaultReceivedConvs) GetConversationId() *string {
	return s.ConversationId
}

func (s *GetReportTemplateByNameResponseBodyDefaultReceivedConvs) GetTitle() *string {
	return s.Title
}

func (s *GetReportTemplateByNameResponseBodyDefaultReceivedConvs) SetConversationId(v string) *GetReportTemplateByNameResponseBodyDefaultReceivedConvs {
	s.ConversationId = &v
	return s
}

func (s *GetReportTemplateByNameResponseBodyDefaultReceivedConvs) SetTitle(v string) *GetReportTemplateByNameResponseBodyDefaultReceivedConvs {
	s.Title = &v
	return s
}

func (s *GetReportTemplateByNameResponseBodyDefaultReceivedConvs) Validate() error {
	return dara.Validate(s)
}

type GetReportTemplateByNameResponseBodyDefaultReceivers struct {
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// example:
	//
	// 1234
	Userid *string `json:"Userid,omitempty" xml:"Userid,omitempty"`
}

func (s GetReportTemplateByNameResponseBodyDefaultReceivers) String() string {
	return dara.Prettify(s)
}

func (s GetReportTemplateByNameResponseBodyDefaultReceivers) GoString() string {
	return s.String()
}

func (s *GetReportTemplateByNameResponseBodyDefaultReceivers) GetUserName() *string {
	return s.UserName
}

func (s *GetReportTemplateByNameResponseBodyDefaultReceivers) GetUserid() *string {
	return s.Userid
}

func (s *GetReportTemplateByNameResponseBodyDefaultReceivers) SetUserName(v string) *GetReportTemplateByNameResponseBodyDefaultReceivers {
	s.UserName = &v
	return s
}

func (s *GetReportTemplateByNameResponseBodyDefaultReceivers) SetUserid(v string) *GetReportTemplateByNameResponseBodyDefaultReceivers {
	s.Userid = &v
	return s
}

func (s *GetReportTemplateByNameResponseBodyDefaultReceivers) Validate() error {
	return dara.Validate(s)
}

type GetReportTemplateByNameResponseBodyFields struct {
	// example:
	//
	// key1
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// example:
	//
	// 0
	Sort *int64 `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetReportTemplateByNameResponseBodyFields) String() string {
	return dara.Prettify(s)
}

func (s GetReportTemplateByNameResponseBodyFields) GoString() string {
	return s.String()
}

func (s *GetReportTemplateByNameResponseBodyFields) GetFieldName() *string {
	return s.FieldName
}

func (s *GetReportTemplateByNameResponseBodyFields) GetSort() *int64 {
	return s.Sort
}

func (s *GetReportTemplateByNameResponseBodyFields) GetType() *int64 {
	return s.Type
}

func (s *GetReportTemplateByNameResponseBodyFields) SetFieldName(v string) *GetReportTemplateByNameResponseBodyFields {
	s.FieldName = &v
	return s
}

func (s *GetReportTemplateByNameResponseBodyFields) SetSort(v int64) *GetReportTemplateByNameResponseBodyFields {
	s.Sort = &v
	return s
}

func (s *GetReportTemplateByNameResponseBodyFields) SetType(v int64) *GetReportTemplateByNameResponseBodyFields {
	s.Type = &v
	return s
}

func (s *GetReportTemplateByNameResponseBodyFields) Validate() error {
	return dara.Validate(s)
}
