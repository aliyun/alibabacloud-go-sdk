// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPurchaseControlRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPurchaseControlRecordResponseBody
	GetCode() *string
	SetData(v []*GetPurchaseControlRecordResponseBodyData) *GetPurchaseControlRecordResponseBody
	GetData() []*GetPurchaseControlRecordResponseBodyData
	SetPageNo(v int32) *GetPurchaseControlRecordResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *GetPurchaseControlRecordResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetPurchaseControlRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPurchaseControlRecordResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *GetPurchaseControlRecordResponseBody
	GetTotal() *int32
}

type GetPurchaseControlRecordResponseBody struct {
	// example:
	//
	// 200
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetPurchaseControlRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// RequestId
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetPurchaseControlRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPurchaseControlRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetPurchaseControlRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPurchaseControlRecordResponseBody) GetData() []*GetPurchaseControlRecordResponseBodyData {
	return s.Data
}

func (s *GetPurchaseControlRecordResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetPurchaseControlRecordResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetPurchaseControlRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPurchaseControlRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPurchaseControlRecordResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetPurchaseControlRecordResponseBody) SetCode(v string) *GetPurchaseControlRecordResponseBody {
	s.Code = &v
	return s
}

func (s *GetPurchaseControlRecordResponseBody) SetData(v []*GetPurchaseControlRecordResponseBodyData) *GetPurchaseControlRecordResponseBody {
	s.Data = v
	return s
}

func (s *GetPurchaseControlRecordResponseBody) SetPageNo(v int32) *GetPurchaseControlRecordResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetPurchaseControlRecordResponseBody) SetPageSize(v int32) *GetPurchaseControlRecordResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetPurchaseControlRecordResponseBody) SetRequestId(v string) *GetPurchaseControlRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPurchaseControlRecordResponseBody) SetSuccess(v bool) *GetPurchaseControlRecordResponseBody {
	s.Success = &v
	return s
}

func (s *GetPurchaseControlRecordResponseBody) SetTotal(v int32) *GetPurchaseControlRecordResponseBody {
	s.Total = &v
	return s
}

func (s *GetPurchaseControlRecordResponseBody) Validate() error {
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

type GetPurchaseControlRecordResponseBodyData struct {
	// example:
	//
	// ban
	ChangedType *string `json:"ChangedType,omitempty" xml:"ChangedType,omitempty"`
	// example:
	//
	// ACPN
	OperationPath *string `json:"OperationPath,omitempty" xml:"OperationPath,omitempty"`
	// example:
	//
	// 2023-12-15 10:34:36
	OperationTime *string `json:"OperationTime,omitempty" xml:"OperationTime,omitempty"`
	// example:
	//
	// 292828565558721922
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s GetPurchaseControlRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPurchaseControlRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPurchaseControlRecordResponseBodyData) GetChangedType() *string {
	return s.ChangedType
}

func (s *GetPurchaseControlRecordResponseBodyData) GetOperationPath() *string {
	return s.OperationPath
}

func (s *GetPurchaseControlRecordResponseBodyData) GetOperationTime() *string {
	return s.OperationTime
}

func (s *GetPurchaseControlRecordResponseBodyData) GetOperator() *string {
	return s.Operator
}

func (s *GetPurchaseControlRecordResponseBodyData) SetChangedType(v string) *GetPurchaseControlRecordResponseBodyData {
	s.ChangedType = &v
	return s
}

func (s *GetPurchaseControlRecordResponseBodyData) SetOperationPath(v string) *GetPurchaseControlRecordResponseBodyData {
	s.OperationPath = &v
	return s
}

func (s *GetPurchaseControlRecordResponseBodyData) SetOperationTime(v string) *GetPurchaseControlRecordResponseBodyData {
	s.OperationTime = &v
	return s
}

func (s *GetPurchaseControlRecordResponseBodyData) SetOperator(v string) *GetPurchaseControlRecordResponseBodyData {
	s.Operator = &v
	return s
}

func (s *GetPurchaseControlRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
