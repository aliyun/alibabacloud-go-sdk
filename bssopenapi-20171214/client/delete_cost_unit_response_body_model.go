// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCostUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCostUnitResponseBody
	GetCode() *string
	SetData(v *DeleteCostUnitResponseBodyData) *DeleteCostUnitResponseBody
	GetData() *DeleteCostUnitResponseBodyData
	SetMessage(v string) *DeleteCostUnitResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCostUnitResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCostUnitResponseBody
	GetSuccess() *bool
}

type DeleteCostUnitResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DeleteCostUnitResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
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
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCostUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCostUnitResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCostUnitResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCostUnitResponseBody) GetData() *DeleteCostUnitResponseBodyData {
	return s.Data
}

func (s *DeleteCostUnitResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCostUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCostUnitResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCostUnitResponseBody) SetCode(v string) *DeleteCostUnitResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCostUnitResponseBody) SetData(v *DeleteCostUnitResponseBodyData) *DeleteCostUnitResponseBody {
	s.Data = v
	return s
}

func (s *DeleteCostUnitResponseBody) SetMessage(v string) *DeleteCostUnitResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCostUnitResponseBody) SetRequestId(v string) *DeleteCostUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCostUnitResponseBody) SetSuccess(v bool) *DeleteCostUnitResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCostUnitResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteCostUnitResponseBodyData struct {
	// Indicates whether the call is complete.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The user ID of the cost center owner.
	//
	// example:
	//
	// 12431
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The ID of the cost center.
	//
	// example:
	//
	// 123412343
	UnitId *int64 `json:"UnitId,omitempty" xml:"UnitId,omitempty"`
}

func (s DeleteCostUnitResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteCostUnitResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteCostUnitResponseBodyData) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteCostUnitResponseBodyData) GetOwnerUid() *int64 {
	return s.OwnerUid
}

func (s *DeleteCostUnitResponseBodyData) GetUnitId() *int64 {
	return s.UnitId
}

func (s *DeleteCostUnitResponseBodyData) SetIsSuccess(v bool) *DeleteCostUnitResponseBodyData {
	s.IsSuccess = &v
	return s
}

func (s *DeleteCostUnitResponseBodyData) SetOwnerUid(v int64) *DeleteCostUnitResponseBodyData {
	s.OwnerUid = &v
	return s
}

func (s *DeleteCostUnitResponseBodyData) SetUnitId(v int64) *DeleteCostUnitResponseBodyData {
	s.UnitId = &v
	return s
}

func (s *DeleteCostUnitResponseBodyData) Validate() error {
	return dara.Validate(s)
}
