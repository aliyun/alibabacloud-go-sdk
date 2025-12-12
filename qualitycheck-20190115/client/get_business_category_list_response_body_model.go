// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBusinessCategoryListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBusinessCategoryListResponseBody
	GetCode() *string
	SetData(v *GetBusinessCategoryListResponseBodyData) *GetBusinessCategoryListResponseBody
	GetData() *GetBusinessCategoryListResponseBodyData
	SetMessage(v string) *GetBusinessCategoryListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBusinessCategoryListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBusinessCategoryListResponseBody
	GetSuccess() *bool
}

type GetBusinessCategoryListResponseBody struct {
	// example:
	//
	// 200
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetBusinessCategoryListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A186A419-FDBE-464C-AED4-7121CAC73BF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBusinessCategoryListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBusinessCategoryListResponseBody) GoString() string {
	return s.String()
}

func (s *GetBusinessCategoryListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBusinessCategoryListResponseBody) GetData() *GetBusinessCategoryListResponseBodyData {
	return s.Data
}

func (s *GetBusinessCategoryListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBusinessCategoryListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBusinessCategoryListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBusinessCategoryListResponseBody) SetCode(v string) *GetBusinessCategoryListResponseBody {
	s.Code = &v
	return s
}

func (s *GetBusinessCategoryListResponseBody) SetData(v *GetBusinessCategoryListResponseBodyData) *GetBusinessCategoryListResponseBody {
	s.Data = v
	return s
}

func (s *GetBusinessCategoryListResponseBody) SetMessage(v string) *GetBusinessCategoryListResponseBody {
	s.Message = &v
	return s
}

func (s *GetBusinessCategoryListResponseBody) SetRequestId(v string) *GetBusinessCategoryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBusinessCategoryListResponseBody) SetSuccess(v bool) *GetBusinessCategoryListResponseBody {
	s.Success = &v
	return s
}

func (s *GetBusinessCategoryListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBusinessCategoryListResponseBodyData struct {
	BusinessCategoryBasicInfo []*GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo `json:"BusinessCategoryBasicInfo,omitempty" xml:"BusinessCategoryBasicInfo,omitempty" type:"Repeated"`
}

func (s GetBusinessCategoryListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBusinessCategoryListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBusinessCategoryListResponseBodyData) GetBusinessCategoryBasicInfo() []*GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo {
	return s.BusinessCategoryBasicInfo
}

func (s *GetBusinessCategoryListResponseBodyData) SetBusinessCategoryBasicInfo(v []*GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) *GetBusinessCategoryListResponseBodyData {
	s.BusinessCategoryBasicInfo = v
	return s
}

func (s *GetBusinessCategoryListResponseBodyData) Validate() error {
	if s.BusinessCategoryBasicInfo != nil {
		for _, item := range s.BusinessCategoryBasicInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo struct {
	// example:
	//
	// 0
	Bid          *int32  `json:"Bid,omitempty" xml:"Bid,omitempty"`
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// example:
	//
	// 0
	ServiceType *int32 `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) GoString() string {
	return s.String()
}

func (s *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) GetBid() *int32 {
	return s.Bid
}

func (s *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) GetBusinessName() *string {
	return s.BusinessName
}

func (s *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) GetServiceType() *int32 {
	return s.ServiceType
}

func (s *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) SetBid(v int32) *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo {
	s.Bid = &v
	return s
}

func (s *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) SetBusinessName(v string) *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo {
	s.BusinessName = &v
	return s
}

func (s *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) SetServiceType(v int32) *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo {
	s.ServiceType = &v
	return s
}

func (s *GetBusinessCategoryListResponseBodyDataBusinessCategoryBasicInfo) Validate() error {
	return dara.Validate(s)
}
