// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCostUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyCostUnitResponseBody
	GetCode() *string
	SetData(v []*ModifyCostUnitResponseBodyData) *ModifyCostUnitResponseBody
	GetData() []*ModifyCostUnitResponseBodyData
	SetMessage(v string) *ModifyCostUnitResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyCostUnitResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyCostUnitResponseBody
	GetSuccess() *bool
}

type ModifyCostUnitResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data []*ModifyCostUnitResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F5B803CF-94D8-43AF-ADB3-D819AAD30E27
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyCostUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostUnitResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCostUnitResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyCostUnitResponseBody) GetData() []*ModifyCostUnitResponseBodyData {
	return s.Data
}

func (s *ModifyCostUnitResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyCostUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCostUnitResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyCostUnitResponseBody) SetCode(v string) *ModifyCostUnitResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyCostUnitResponseBody) SetData(v []*ModifyCostUnitResponseBodyData) *ModifyCostUnitResponseBody {
	s.Data = v
	return s
}

func (s *ModifyCostUnitResponseBody) SetMessage(v string) *ModifyCostUnitResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyCostUnitResponseBody) SetRequestId(v string) *ModifyCostUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCostUnitResponseBody) SetSuccess(v bool) *ModifyCostUnitResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyCostUnitResponseBody) Validate() error {
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

type ModifyCostUnitResponseBodyData struct {
	// Indicates whether the cost center was modified.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The user ID of the cost center owner.
	//
	// example:
	//
	// 823756287
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The ID of the cost center.
	//
	// example:
	//
	// 356349875
	UnitId *int64 `json:"UnitId,omitempty" xml:"UnitId,omitempty"`
}

func (s ModifyCostUnitResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostUnitResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyCostUnitResponseBodyData) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ModifyCostUnitResponseBodyData) GetOwnerUid() *int64 {
	return s.OwnerUid
}

func (s *ModifyCostUnitResponseBodyData) GetUnitId() *int64 {
	return s.UnitId
}

func (s *ModifyCostUnitResponseBodyData) SetIsSuccess(v bool) *ModifyCostUnitResponseBodyData {
	s.IsSuccess = &v
	return s
}

func (s *ModifyCostUnitResponseBodyData) SetOwnerUid(v int64) *ModifyCostUnitResponseBodyData {
	s.OwnerUid = &v
	return s
}

func (s *ModifyCostUnitResponseBodyData) SetUnitId(v int64) *ModifyCostUnitResponseBodyData {
	s.UnitId = &v
	return s
}

func (s *ModifyCostUnitResponseBodyData) Validate() error {
	return dara.Validate(s)
}
