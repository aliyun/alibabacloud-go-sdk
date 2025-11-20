// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketOperateRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecords(v []*ListTicketOperateRecordResponseBodyRecords) *ListTicketOperateRecordResponseBody
	GetRecords() []*ListTicketOperateRecordResponseBodyRecords
	SetRequestId(v string) *ListTicketOperateRecordResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *ListTicketOperateRecordResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *ListTicketOperateRecordResponseBody
	GetVendorType() *string
}

type ListTicketOperateRecordResponseBody struct {
	Records []*ListTicketOperateRecordResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s ListTicketOperateRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTicketOperateRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordResponseBody) GetRecords() []*ListTicketOperateRecordResponseBodyRecords {
	return s.Records
}

func (s *ListTicketOperateRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTicketOperateRecordResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *ListTicketOperateRecordResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *ListTicketOperateRecordResponseBody) SetRecords(v []*ListTicketOperateRecordResponseBodyRecords) *ListTicketOperateRecordResponseBody {
	s.Records = v
	return s
}

func (s *ListTicketOperateRecordResponseBody) SetRequestId(v string) *ListTicketOperateRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTicketOperateRecordResponseBody) SetVendorRequestId(v string) *ListTicketOperateRecordResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *ListTicketOperateRecordResponseBody) SetVendorType(v string) *ListTicketOperateRecordResponseBody {
	s.VendorType = &v
	return s
}

func (s *ListTicketOperateRecordResponseBody) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTicketOperateRecordResponseBodyRecords struct {
	// example:
	//
	// a8iSxxxxgtgiE
	OpenTicketId *string `json:"OpenTicketId,omitempty" xml:"OpenTicketId,omitempty"`
	// example:
	//
	// {"originTakers":[{"nickName":"贤文","unionId":"Dq9hxxxxwiEiE"},{"nickName":"王鸿程","unionId":"4kITooxxxx5wiEiE"}]}
	OperateData *string `json:"OperateData,omitempty" xml:"OperateData,omitempty"`
	// example:
	//
	// 2021-07-09 19:26:09
	OperateTime *string `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	// example:
	//
	// FINISH
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// 结单
	OperationDisplayName *string                                               `json:"OperationDisplayName,omitempty" xml:"OperationDisplayName,omitempty"`
	Operator             *ListTicketOperateRecordResponseBodyRecordsOperator   `json:"Operator,omitempty" xml:"Operator,omitempty" type:"Struct"`
	TicketMemo           *ListTicketOperateRecordResponseBodyRecordsTicketMemo `json:"TicketMemo,omitempty" xml:"TicketMemo,omitempty" type:"Struct"`
}

func (s ListTicketOperateRecordResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s ListTicketOperateRecordResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordResponseBodyRecords) GetOpenTicketId() *string {
	return s.OpenTicketId
}

func (s *ListTicketOperateRecordResponseBodyRecords) GetOperateData() *string {
	return s.OperateData
}

func (s *ListTicketOperateRecordResponseBodyRecords) GetOperateTime() *string {
	return s.OperateTime
}

func (s *ListTicketOperateRecordResponseBodyRecords) GetOperation() *string {
	return s.Operation
}

func (s *ListTicketOperateRecordResponseBodyRecords) GetOperationDisplayName() *string {
	return s.OperationDisplayName
}

func (s *ListTicketOperateRecordResponseBodyRecords) GetOperator() *ListTicketOperateRecordResponseBodyRecordsOperator {
	return s.Operator
}

func (s *ListTicketOperateRecordResponseBodyRecords) GetTicketMemo() *ListTicketOperateRecordResponseBodyRecordsTicketMemo {
	return s.TicketMemo
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetOpenTicketId(v string) *ListTicketOperateRecordResponseBodyRecords {
	s.OpenTicketId = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetOperateData(v string) *ListTicketOperateRecordResponseBodyRecords {
	s.OperateData = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetOperateTime(v string) *ListTicketOperateRecordResponseBodyRecords {
	s.OperateTime = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetOperation(v string) *ListTicketOperateRecordResponseBodyRecords {
	s.Operation = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetOperationDisplayName(v string) *ListTicketOperateRecordResponseBodyRecords {
	s.OperationDisplayName = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetOperator(v *ListTicketOperateRecordResponseBodyRecordsOperator) *ListTicketOperateRecordResponseBodyRecords {
	s.Operator = v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecords) SetTicketMemo(v *ListTicketOperateRecordResponseBodyRecordsTicketMemo) *ListTicketOperateRecordResponseBodyRecords {
	s.TicketMemo = v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecords) Validate() error {
	if s.Operator != nil {
		if err := s.Operator.Validate(); err != nil {
			return err
		}
	}
	if s.TicketMemo != nil {
		if err := s.TicketMemo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTicketOperateRecordResponseBodyRecordsOperator struct {
	// example:
	//
	// 贤文
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// example:
	//
	// 012345
	UnionId *string `json:"UnionId,omitempty" xml:"UnionId,omitempty"`
}

func (s ListTicketOperateRecordResponseBodyRecordsOperator) String() string {
	return dara.Prettify(s)
}

func (s ListTicketOperateRecordResponseBodyRecordsOperator) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordResponseBodyRecordsOperator) GetNickName() *string {
	return s.NickName
}

func (s *ListTicketOperateRecordResponseBodyRecordsOperator) GetUnionId() *string {
	return s.UnionId
}

func (s *ListTicketOperateRecordResponseBodyRecordsOperator) SetNickName(v string) *ListTicketOperateRecordResponseBodyRecordsOperator {
	s.NickName = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecordsOperator) SetUnionId(v string) *ListTicketOperateRecordResponseBodyRecordsOperator {
	s.UnionId = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecordsOperator) Validate() error {
	return dara.Validate(s)
}

type ListTicketOperateRecordResponseBodyRecordsTicketMemo struct {
	Attachments []*ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
	// example:
	//
	// 贤文结束工单
	Memo *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
}

func (s ListTicketOperateRecordResponseBodyRecordsTicketMemo) String() string {
	return dara.Prettify(s)
}

func (s ListTicketOperateRecordResponseBodyRecordsTicketMemo) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordResponseBodyRecordsTicketMemo) GetAttachments() []*ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments {
	return s.Attachments
}

func (s *ListTicketOperateRecordResponseBodyRecordsTicketMemo) GetMemo() *string {
	return s.Memo
}

func (s *ListTicketOperateRecordResponseBodyRecordsTicketMemo) SetAttachments(v []*ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments) *ListTicketOperateRecordResponseBodyRecordsTicketMemo {
	s.Attachments = v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecordsTicketMemo) SetMemo(v string) *ListTicketOperateRecordResponseBodyRecordsTicketMemo {
	s.Memo = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecordsTicketMemo) Validate() error {
	if s.Attachments != nil {
		for _, item := range s.Attachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments struct {
	// example:
	//
	// wahaha.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// ticket/image/447xxxx9/43003/e2xxxec4243e940a1367_1625xxxx99.txt
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments) String() string {
	return dara.Prettify(s)
}

func (s ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments) GetFileName() *string {
	return s.FileName
}

func (s *ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments) GetKey() *string {
	return s.Key
}

func (s *ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments) SetFileName(v string) *ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments) SetKey(v string) *ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments {
	s.Key = &v
	return s
}

func (s *ListTicketOperateRecordResponseBodyRecordsTicketMemoAttachments) Validate() error {
	return dara.Validate(s)
}
