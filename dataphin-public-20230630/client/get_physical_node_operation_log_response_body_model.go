// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalNodeOperationLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPhysicalNodeOperationLogResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetPhysicalNodeOperationLogResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPhysicalNodeOperationLogResponseBody
	GetMessage() *string
	SetOperationLogList(v []*GetPhysicalNodeOperationLogResponseBodyOperationLogList) *GetPhysicalNodeOperationLogResponseBody
	GetOperationLogList() []*GetPhysicalNodeOperationLogResponseBodyOperationLogList
	SetRequestId(v string) *GetPhysicalNodeOperationLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPhysicalNodeOperationLogResponseBody
	GetSuccess() *bool
}

type GetPhysicalNodeOperationLogResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message          *string                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	OperationLogList []*GetPhysicalNodeOperationLogResponseBodyOperationLogList `json:"OperationLogList,omitempty" xml:"OperationLogList,omitempty" type:"Repeated"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPhysicalNodeOperationLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeOperationLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeOperationLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPhysicalNodeOperationLogResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPhysicalNodeOperationLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPhysicalNodeOperationLogResponseBody) GetOperationLogList() []*GetPhysicalNodeOperationLogResponseBodyOperationLogList {
	return s.OperationLogList
}

func (s *GetPhysicalNodeOperationLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPhysicalNodeOperationLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPhysicalNodeOperationLogResponseBody) SetCode(v string) *GetPhysicalNodeOperationLogResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBody) SetHttpStatusCode(v int32) *GetPhysicalNodeOperationLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBody) SetMessage(v string) *GetPhysicalNodeOperationLogResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBody) SetOperationLogList(v []*GetPhysicalNodeOperationLogResponseBodyOperationLogList) *GetPhysicalNodeOperationLogResponseBody {
	s.OperationLogList = v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBody) SetRequestId(v string) *GetPhysicalNodeOperationLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBody) SetSuccess(v bool) *GetPhysicalNodeOperationLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBody) Validate() error {
	if s.OperationLogList != nil {
		for _, item := range s.OperationLogList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPhysicalNodeOperationLogResponseBodyOperationLogList struct {
	// example:
	//
	// xx
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// example:
	//
	// 2024-05-30 16:47:13
	OperationTime *string `json:"OperationTime,omitempty" xml:"OperationTime,omitempty"`
	// example:
	//
	// PAUSE_TASK
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// example:
	//
	// 132222
	Operator     *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
}

func (s GetPhysicalNodeOperationLogResponseBodyOperationLogList) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeOperationLogResponseBodyOperationLogList) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeOperationLogResponseBodyOperationLogList) GetContext() *string {
	return s.Context
}

func (s *GetPhysicalNodeOperationLogResponseBodyOperationLogList) GetOperationTime() *string {
	return s.OperationTime
}

func (s *GetPhysicalNodeOperationLogResponseBodyOperationLogList) GetOperationType() *string {
	return s.OperationType
}

func (s *GetPhysicalNodeOperationLogResponseBodyOperationLogList) GetOperator() *string {
	return s.Operator
}

func (s *GetPhysicalNodeOperationLogResponseBodyOperationLogList) GetOperatorName() *string {
	return s.OperatorName
}

func (s *GetPhysicalNodeOperationLogResponseBodyOperationLogList) SetContext(v string) *GetPhysicalNodeOperationLogResponseBodyOperationLogList {
	s.Context = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBodyOperationLogList) SetOperationTime(v string) *GetPhysicalNodeOperationLogResponseBodyOperationLogList {
	s.OperationTime = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBodyOperationLogList) SetOperationType(v string) *GetPhysicalNodeOperationLogResponseBodyOperationLogList {
	s.OperationType = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBodyOperationLogList) SetOperator(v string) *GetPhysicalNodeOperationLogResponseBodyOperationLogList {
	s.Operator = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBodyOperationLogList) SetOperatorName(v string) *GetPhysicalNodeOperationLogResponseBodyOperationLogList {
	s.OperatorName = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBodyOperationLogList) Validate() error {
	return dara.Validate(s)
}
