// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceTableMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSourceTableMetaResponseBody
	GetCode() *string
	SetData(v *GetSourceTableMetaResponseBodyData) *GetSourceTableMetaResponseBody
	GetData() *GetSourceTableMetaResponseBodyData
	SetHttpStatusCode(v int32) *GetSourceTableMetaResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSourceTableMetaResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSourceTableMetaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSourceTableMetaResponseBody
	GetSuccess() *bool
}

type GetSourceTableMetaResponseBody struct {
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSourceTableMetaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSourceTableMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSourceTableMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetSourceTableMetaResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSourceTableMetaResponseBody) GetData() *GetSourceTableMetaResponseBodyData {
	return s.Data
}

func (s *GetSourceTableMetaResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSourceTableMetaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSourceTableMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSourceTableMetaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSourceTableMetaResponseBody) SetCode(v string) *GetSourceTableMetaResponseBody {
	s.Code = &v
	return s
}

func (s *GetSourceTableMetaResponseBody) SetData(v *GetSourceTableMetaResponseBodyData) *GetSourceTableMetaResponseBody {
	s.Data = v
	return s
}

func (s *GetSourceTableMetaResponseBody) SetHttpStatusCode(v int32) *GetSourceTableMetaResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSourceTableMetaResponseBody) SetMessage(v string) *GetSourceTableMetaResponseBody {
	s.Message = &v
	return s
}

func (s *GetSourceTableMetaResponseBody) SetRequestId(v string) *GetSourceTableMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSourceTableMetaResponseBody) SetSuccess(v bool) *GetSourceTableMetaResponseBody {
	s.Success = &v
	return s
}

func (s *GetSourceTableMetaResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSourceTableMetaResponseBodyData struct {
	Columns []*GetSourceTableMetaResponseBodyDataColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// example:
	//
	// 300001410.default.sample
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// example:
	//
	// sample
	TableComment *string `json:"TableComment,omitempty" xml:"TableComment,omitempty"`
	// example:
	//
	// sample
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetSourceTableMetaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSourceTableMetaResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSourceTableMetaResponseBodyData) GetColumns() []*GetSourceTableMetaResponseBodyDataColumns {
	return s.Columns
}

func (s *GetSourceTableMetaResponseBodyData) GetGuid() *string {
	return s.Guid
}

func (s *GetSourceTableMetaResponseBodyData) GetTableComment() *string {
	return s.TableComment
}

func (s *GetSourceTableMetaResponseBodyData) GetTableName() *string {
	return s.TableName
}

func (s *GetSourceTableMetaResponseBodyData) SetColumns(v []*GetSourceTableMetaResponseBodyDataColumns) *GetSourceTableMetaResponseBodyData {
	s.Columns = v
	return s
}

func (s *GetSourceTableMetaResponseBodyData) SetGuid(v string) *GetSourceTableMetaResponseBodyData {
	s.Guid = &v
	return s
}

func (s *GetSourceTableMetaResponseBodyData) SetTableComment(v string) *GetSourceTableMetaResponseBodyData {
	s.TableComment = &v
	return s
}

func (s *GetSourceTableMetaResponseBodyData) SetTableName(v string) *GetSourceTableMetaResponseBodyData {
	s.TableName = &v
	return s
}

func (s *GetSourceTableMetaResponseBodyData) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSourceTableMetaResponseBodyDataColumns struct {
	// example:
	//
	// unique id
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// bigint
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// false
	Pk *bool `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// example:
	//
	// false
	Pt *bool `json:"Pt,omitempty" xml:"Pt,omitempty"`
	// example:
	//
	// bigint
	RawDataType *string `json:"RawDataType,omitempty" xml:"RawDataType,omitempty"`
	// example:
	//
	// 1
	SeqNumber *int32 `json:"SeqNumber,omitempty" xml:"SeqNumber,omitempty"`
}

func (s GetSourceTableMetaResponseBodyDataColumns) String() string {
	return dara.Prettify(s)
}

func (s GetSourceTableMetaResponseBodyDataColumns) GoString() string {
	return s.String()
}

func (s *GetSourceTableMetaResponseBodyDataColumns) GetComment() *string {
	return s.Comment
}

func (s *GetSourceTableMetaResponseBodyDataColumns) GetDataType() *string {
	return s.DataType
}

func (s *GetSourceTableMetaResponseBodyDataColumns) GetName() *string {
	return s.Name
}

func (s *GetSourceTableMetaResponseBodyDataColumns) GetPk() *bool {
	return s.Pk
}

func (s *GetSourceTableMetaResponseBodyDataColumns) GetPt() *bool {
	return s.Pt
}

func (s *GetSourceTableMetaResponseBodyDataColumns) GetRawDataType() *string {
	return s.RawDataType
}

func (s *GetSourceTableMetaResponseBodyDataColumns) GetSeqNumber() *int32 {
	return s.SeqNumber
}

func (s *GetSourceTableMetaResponseBodyDataColumns) SetComment(v string) *GetSourceTableMetaResponseBodyDataColumns {
	s.Comment = &v
	return s
}

func (s *GetSourceTableMetaResponseBodyDataColumns) SetDataType(v string) *GetSourceTableMetaResponseBodyDataColumns {
	s.DataType = &v
	return s
}

func (s *GetSourceTableMetaResponseBodyDataColumns) SetName(v string) *GetSourceTableMetaResponseBodyDataColumns {
	s.Name = &v
	return s
}

func (s *GetSourceTableMetaResponseBodyDataColumns) SetPk(v bool) *GetSourceTableMetaResponseBodyDataColumns {
	s.Pk = &v
	return s
}

func (s *GetSourceTableMetaResponseBodyDataColumns) SetPt(v bool) *GetSourceTableMetaResponseBodyDataColumns {
	s.Pt = &v
	return s
}

func (s *GetSourceTableMetaResponseBodyDataColumns) SetRawDataType(v string) *GetSourceTableMetaResponseBodyDataColumns {
	s.RawDataType = &v
	return s
}

func (s *GetSourceTableMetaResponseBodyDataColumns) SetSeqNumber(v int32) *GetSourceTableMetaResponseBodyDataColumns {
	s.SeqNumber = &v
	return s
}

func (s *GetSourceTableMetaResponseBodyDataColumns) Validate() error {
	return dara.Validate(s)
}
