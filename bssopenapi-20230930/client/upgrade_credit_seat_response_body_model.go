// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeCreditSeatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpgradeCreditSeatResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpgradeCreditSeatResponseBody
	GetCode() *string
	SetData(v *UpgradeCreditSeatResponseBodyData) *UpgradeCreditSeatResponseBody
	GetData() *UpgradeCreditSeatResponseBodyData
	SetMessage(v string) *UpgradeCreditSeatResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpgradeCreditSeatResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradeCreditSeatResponseBody
	GetSuccess() *bool
}

type UpgradeCreditSeatResponseBody struct {
	AccessDeniedDetail *string                            `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *UpgradeCreditSeatResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradeCreditSeatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeCreditSeatResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeCreditSeatResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpgradeCreditSeatResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpgradeCreditSeatResponseBody) GetData() *UpgradeCreditSeatResponseBodyData {
	return s.Data
}

func (s *UpgradeCreditSeatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradeCreditSeatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeCreditSeatResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeCreditSeatResponseBody) SetAccessDeniedDetail(v string) *UpgradeCreditSeatResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpgradeCreditSeatResponseBody) SetCode(v string) *UpgradeCreditSeatResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeCreditSeatResponseBody) SetData(v *UpgradeCreditSeatResponseBodyData) *UpgradeCreditSeatResponseBody {
	s.Data = v
	return s
}

func (s *UpgradeCreditSeatResponseBody) SetMessage(v string) *UpgradeCreditSeatResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeCreditSeatResponseBody) SetRequestId(v string) *UpgradeCreditSeatResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeCreditSeatResponseBody) SetSuccess(v bool) *UpgradeCreditSeatResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeCreditSeatResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpgradeCreditSeatResponseBodyData struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OrderId    *int64    `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s UpgradeCreditSeatResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpgradeCreditSeatResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpgradeCreditSeatResponseBodyData) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *UpgradeCreditSeatResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *UpgradeCreditSeatResponseBodyData) SetInstanceId(v []*string) *UpgradeCreditSeatResponseBodyData {
	s.InstanceId = v
	return s
}

func (s *UpgradeCreditSeatResponseBodyData) SetOrderId(v int64) *UpgradeCreditSeatResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *UpgradeCreditSeatResponseBodyData) Validate() error {
	return dara.Validate(s)
}
