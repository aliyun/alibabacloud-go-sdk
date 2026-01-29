// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCostUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCostUnitResponseBody
	GetCode() *string
	SetData(v *CreateCostUnitResponseBodyData) *CreateCostUnitResponseBody
	GetData() *CreateCostUnitResponseBodyData
	SetMessage(v string) *CreateCostUnitResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCostUnitResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCostUnitResponseBody
	GetSuccess() *bool
}

type CreateCostUnitResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *CreateCostUnitResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// Successful！
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F5B803CF-94D8-43AF-ADB3-D819AAD30E27
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCostUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCostUnitResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCostUnitResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCostUnitResponseBody) GetData() *CreateCostUnitResponseBodyData {
	return s.Data
}

func (s *CreateCostUnitResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCostUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCostUnitResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCostUnitResponseBody) SetCode(v string) *CreateCostUnitResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCostUnitResponseBody) SetData(v *CreateCostUnitResponseBodyData) *CreateCostUnitResponseBody {
	s.Data = v
	return s
}

func (s *CreateCostUnitResponseBody) SetMessage(v string) *CreateCostUnitResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCostUnitResponseBody) SetRequestId(v string) *CreateCostUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCostUnitResponseBody) SetSuccess(v bool) *CreateCostUnitResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCostUnitResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCostUnitResponseBodyData struct {
	// The list of cost center entities.
	CostUnitDtoList []*CreateCostUnitResponseBodyDataCostUnitDtoList `json:"CostUnitDtoList,omitempty" xml:"CostUnitDtoList,omitempty" type:"Repeated"`
}

func (s CreateCostUnitResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCostUnitResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCostUnitResponseBodyData) GetCostUnitDtoList() []*CreateCostUnitResponseBodyDataCostUnitDtoList {
	return s.CostUnitDtoList
}

func (s *CreateCostUnitResponseBodyData) SetCostUnitDtoList(v []*CreateCostUnitResponseBodyDataCostUnitDtoList) *CreateCostUnitResponseBodyData {
	s.CostUnitDtoList = v
	return s
}

func (s *CreateCostUnitResponseBodyData) Validate() error {
	if s.CostUnitDtoList != nil {
		for _, item := range s.CostUnitDtoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCostUnitResponseBodyDataCostUnitDtoList struct {
	// The user ID of the owner of the cost center.
	//
	// example:
	//
	// 26387563
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The ID of the parent cost center. A value of -1 indicates the root cost center.
	//
	// example:
	//
	// -1
	ParentUnitId *int64 `json:"ParentUnitId,omitempty" xml:"ParentUnitId,omitempty"`
	// The ID of the cost center.
	//
	// example:
	//
	// 84327659328
	UnitId *int64 `json:"UnitId,omitempty" xml:"UnitId,omitempty"`
	// The name of the cost center.
	//
	// example:
	//
	// test
	UnitName *string `json:"UnitName,omitempty" xml:"UnitName,omitempty"`
}

func (s CreateCostUnitResponseBodyDataCostUnitDtoList) String() string {
	return dara.Prettify(s)
}

func (s CreateCostUnitResponseBodyDataCostUnitDtoList) GoString() string {
	return s.String()
}

func (s *CreateCostUnitResponseBodyDataCostUnitDtoList) GetOwnerUid() *int64 {
	return s.OwnerUid
}

func (s *CreateCostUnitResponseBodyDataCostUnitDtoList) GetParentUnitId() *int64 {
	return s.ParentUnitId
}

func (s *CreateCostUnitResponseBodyDataCostUnitDtoList) GetUnitId() *int64 {
	return s.UnitId
}

func (s *CreateCostUnitResponseBodyDataCostUnitDtoList) GetUnitName() *string {
	return s.UnitName
}

func (s *CreateCostUnitResponseBodyDataCostUnitDtoList) SetOwnerUid(v int64) *CreateCostUnitResponseBodyDataCostUnitDtoList {
	s.OwnerUid = &v
	return s
}

func (s *CreateCostUnitResponseBodyDataCostUnitDtoList) SetParentUnitId(v int64) *CreateCostUnitResponseBodyDataCostUnitDtoList {
	s.ParentUnitId = &v
	return s
}

func (s *CreateCostUnitResponseBodyDataCostUnitDtoList) SetUnitId(v int64) *CreateCostUnitResponseBodyDataCostUnitDtoList {
	s.UnitId = &v
	return s
}

func (s *CreateCostUnitResponseBodyDataCostUnitDtoList) SetUnitName(v string) *CreateCostUnitResponseBodyDataCostUnitDtoList {
	s.UnitName = &v
	return s
}

func (s *CreateCostUnitResponseBodyDataCostUnitDtoList) Validate() error {
	return dara.Validate(s)
}
