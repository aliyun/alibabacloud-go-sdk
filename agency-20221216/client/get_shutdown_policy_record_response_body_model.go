// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetShutdownPolicyRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetShutdownPolicyRecordResponseBody
	GetCode() *string
	SetData(v []*GetShutdownPolicyRecordResponseBodyData) *GetShutdownPolicyRecordResponseBody
	GetData() []*GetShutdownPolicyRecordResponseBodyData
	SetPageNo(v int32) *GetShutdownPolicyRecordResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *GetShutdownPolicyRecordResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetShutdownPolicyRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetShutdownPolicyRecordResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *GetShutdownPolicyRecordResponseBody
	GetTotal() *int32
}

type GetShutdownPolicyRecordResponseBody struct {
	// example:
	//
	// 200
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetShutdownPolicyRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 210e876f16704666020714468dab35
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

func (s GetShutdownPolicyRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetShutdownPolicyRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetShutdownPolicyRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetShutdownPolicyRecordResponseBody) GetData() []*GetShutdownPolicyRecordResponseBodyData {
	return s.Data
}

func (s *GetShutdownPolicyRecordResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetShutdownPolicyRecordResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetShutdownPolicyRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetShutdownPolicyRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetShutdownPolicyRecordResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetShutdownPolicyRecordResponseBody) SetCode(v string) *GetShutdownPolicyRecordResponseBody {
	s.Code = &v
	return s
}

func (s *GetShutdownPolicyRecordResponseBody) SetData(v []*GetShutdownPolicyRecordResponseBodyData) *GetShutdownPolicyRecordResponseBody {
	s.Data = v
	return s
}

func (s *GetShutdownPolicyRecordResponseBody) SetPageNo(v int32) *GetShutdownPolicyRecordResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetShutdownPolicyRecordResponseBody) SetPageSize(v int32) *GetShutdownPolicyRecordResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetShutdownPolicyRecordResponseBody) SetRequestId(v string) *GetShutdownPolicyRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetShutdownPolicyRecordResponseBody) SetSuccess(v bool) *GetShutdownPolicyRecordResponseBody {
	s.Success = &v
	return s
}

func (s *GetShutdownPolicyRecordResponseBody) SetTotal(v int32) *GetShutdownPolicyRecordResponseBody {
	s.Total = &v
	return s
}

func (s *GetShutdownPolicyRecordResponseBody) Validate() error {
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

type GetShutdownPolicyRecordResponseBodyData struct {
	// example:
	//
	// noStop
	CurrentPolicy *string `json:"CurrentPolicy,omitempty" xml:"CurrentPolicy,omitempty"`
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
	// 11111111111
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// immediatelyStop
	PreviousPolicy *string `json:"PreviousPolicy,omitempty" xml:"PreviousPolicy,omitempty"`
}

func (s GetShutdownPolicyRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetShutdownPolicyRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetShutdownPolicyRecordResponseBodyData) GetCurrentPolicy() *string {
	return s.CurrentPolicy
}

func (s *GetShutdownPolicyRecordResponseBodyData) GetOperationPath() *string {
	return s.OperationPath
}

func (s *GetShutdownPolicyRecordResponseBodyData) GetOperationTime() *string {
	return s.OperationTime
}

func (s *GetShutdownPolicyRecordResponseBodyData) GetOperator() *string {
	return s.Operator
}

func (s *GetShutdownPolicyRecordResponseBodyData) GetPreviousPolicy() *string {
	return s.PreviousPolicy
}

func (s *GetShutdownPolicyRecordResponseBodyData) SetCurrentPolicy(v string) *GetShutdownPolicyRecordResponseBodyData {
	s.CurrentPolicy = &v
	return s
}

func (s *GetShutdownPolicyRecordResponseBodyData) SetOperationPath(v string) *GetShutdownPolicyRecordResponseBodyData {
	s.OperationPath = &v
	return s
}

func (s *GetShutdownPolicyRecordResponseBodyData) SetOperationTime(v string) *GetShutdownPolicyRecordResponseBodyData {
	s.OperationTime = &v
	return s
}

func (s *GetShutdownPolicyRecordResponseBodyData) SetOperator(v string) *GetShutdownPolicyRecordResponseBodyData {
	s.Operator = &v
	return s
}

func (s *GetShutdownPolicyRecordResponseBodyData) SetPreviousPolicy(v string) *GetShutdownPolicyRecordResponseBodyData {
	s.PreviousPolicy = &v
	return s
}

func (s *GetShutdownPolicyRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
