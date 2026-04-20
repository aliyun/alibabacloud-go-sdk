// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstructionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstructionsResponseBody
	GetCode() *string
	SetData(v *ListInstructionsResponseBodyData) *ListInstructionsResponseBody
	GetData() *ListInstructionsResponseBodyData
	SetHttpStatusCode(v int32) *ListInstructionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListInstructionsResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListInstructionsResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListInstructionsResponseBody
	GetRequestId() *string
}

type ListInstructionsResponseBody struct {
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListInstructionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// e48e45dd-e47a-4744-a063-f08cbebb1c5a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstructionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstructionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstructionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstructionsResponseBody) GetData() *ListInstructionsResponseBodyData {
	return s.Data
}

func (s *ListInstructionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInstructionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstructionsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListInstructionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstructionsResponseBody) SetCode(v string) *ListInstructionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstructionsResponseBody) SetData(v *ListInstructionsResponseBodyData) *ListInstructionsResponseBody {
	s.Data = v
	return s
}

func (s *ListInstructionsResponseBody) SetHttpStatusCode(v int32) *ListInstructionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstructionsResponseBody) SetMessage(v string) *ListInstructionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstructionsResponseBody) SetParams(v []*string) *ListInstructionsResponseBody {
	s.Params = v
	return s
}

func (s *ListInstructionsResponseBody) SetRequestId(v string) *ListInstructionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstructionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstructionsResponseBodyData struct {
	Instructions []*ListInstructionsResponseBodyDataInstructions `json:"Instructions,omitempty" xml:"Instructions,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstructionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstructionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstructionsResponseBodyData) GetInstructions() []*ListInstructionsResponseBodyDataInstructions {
	return s.Instructions
}

func (s *ListInstructionsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstructionsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstructionsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInstructionsResponseBodyData) SetInstructions(v []*ListInstructionsResponseBodyDataInstructions) *ListInstructionsResponseBodyData {
	s.Instructions = v
	return s
}

func (s *ListInstructionsResponseBodyData) SetPageNumber(v int32) *ListInstructionsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListInstructionsResponseBodyData) SetPageSize(v int32) *ListInstructionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInstructionsResponseBodyData) SetTotalCount(v int32) *ListInstructionsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListInstructionsResponseBodyData) Validate() error {
	if s.Instructions != nil {
		for _, item := range s.Instructions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstructionsResponseBodyDataInstructions struct {
	// example:
	//
	// Transfer00
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {
	//
	// 	"providerId": "xxxxxxxxx",
	//
	// 	"transferMethod": "BYE",
	//
	// 	"transferType": "External",
	//
	// 	"transferTarget": "152******",
	//
	// 	"transferor": "01*****",
	//
	// 	"timeoutSeconds": "10"
	//
	// }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 36e9a4cd-a821-4f78-86fa-a9a4aefeea2e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// TRANSFER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstructionsResponseBodyDataInstructions) String() string {
	return dara.Prettify(s)
}

func (s ListInstructionsResponseBodyDataInstructions) GoString() string {
	return s.String()
}

func (s *ListInstructionsResponseBodyDataInstructions) GetCode() *string {
	return s.Code
}

func (s *ListInstructionsResponseBodyDataInstructions) GetConfig() *string {
	return s.Config
}

func (s *ListInstructionsResponseBodyDataInstructions) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstructionsResponseBodyDataInstructions) GetName() *string {
	return s.Name
}

func (s *ListInstructionsResponseBodyDataInstructions) GetType() *string {
	return s.Type
}

func (s *ListInstructionsResponseBodyDataInstructions) SetCode(v string) *ListInstructionsResponseBodyDataInstructions {
	s.Code = &v
	return s
}

func (s *ListInstructionsResponseBodyDataInstructions) SetConfig(v string) *ListInstructionsResponseBodyDataInstructions {
	s.Config = &v
	return s
}

func (s *ListInstructionsResponseBodyDataInstructions) SetInstanceId(v string) *ListInstructionsResponseBodyDataInstructions {
	s.InstanceId = &v
	return s
}

func (s *ListInstructionsResponseBodyDataInstructions) SetName(v string) *ListInstructionsResponseBodyDataInstructions {
	s.Name = &v
	return s
}

func (s *ListInstructionsResponseBodyDataInstructions) SetType(v string) *ListInstructionsResponseBodyDataInstructions {
	s.Type = &v
	return s
}

func (s *ListInstructionsResponseBodyDataInstructions) Validate() error {
	return dara.Validate(s)
}
