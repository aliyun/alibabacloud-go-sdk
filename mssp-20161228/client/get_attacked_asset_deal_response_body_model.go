// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackedAssetDealResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAttackedAssetDealResponseBody
	GetCode() *string
	SetData(v *GetAttackedAssetDealResponseBodyData) *GetAttackedAssetDealResponseBody
	GetData() *GetAttackedAssetDealResponseBodyData
	SetHttpStatusCode(v int32) *GetAttackedAssetDealResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAttackedAssetDealResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAttackedAssetDealResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAttackedAssetDealResponseBody
	GetSuccess() *bool
}

type GetAttackedAssetDealResponseBody struct {
	// Interface return code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data query result.
	Data *GetAttackedAssetDealResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1E74F11C-B4A8-5774-962C-02003BA8504E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the query was successful.<br />
	//
	// **Enum values:**
	//
	// 	- true: Success.
	//
	// 	- false: Failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAttackedAssetDealResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAttackedAssetDealResponseBody) GoString() string {
	return s.String()
}

func (s *GetAttackedAssetDealResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAttackedAssetDealResponseBody) GetData() *GetAttackedAssetDealResponseBodyData {
	return s.Data
}

func (s *GetAttackedAssetDealResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAttackedAssetDealResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAttackedAssetDealResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAttackedAssetDealResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAttackedAssetDealResponseBody) SetCode(v string) *GetAttackedAssetDealResponseBody {
	s.Code = &v
	return s
}

func (s *GetAttackedAssetDealResponseBody) SetData(v *GetAttackedAssetDealResponseBodyData) *GetAttackedAssetDealResponseBody {
	s.Data = v
	return s
}

func (s *GetAttackedAssetDealResponseBody) SetHttpStatusCode(v int32) *GetAttackedAssetDealResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAttackedAssetDealResponseBody) SetMessage(v string) *GetAttackedAssetDealResponseBody {
	s.Message = &v
	return s
}

func (s *GetAttackedAssetDealResponseBody) SetRequestId(v string) *GetAttackedAssetDealResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAttackedAssetDealResponseBody) SetSuccess(v bool) *GetAttackedAssetDealResponseBody {
	s.Success = &v
	return s
}

func (s *GetAttackedAssetDealResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAttackedAssetDealResponseBodyData struct {
	// Collection of attacked asset convergence trends.
	EcsTrendList []*GetAttackedAssetDealResponseBodyDataEcsTrendList `json:"EcsTrendList,omitempty" xml:"EcsTrendList,omitempty" type:"Repeated"`
}

func (s GetAttackedAssetDealResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAttackedAssetDealResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAttackedAssetDealResponseBodyData) GetEcsTrendList() []*GetAttackedAssetDealResponseBodyDataEcsTrendList {
	return s.EcsTrendList
}

func (s *GetAttackedAssetDealResponseBodyData) SetEcsTrendList(v []*GetAttackedAssetDealResponseBodyDataEcsTrendList) *GetAttackedAssetDealResponseBodyData {
	s.EcsTrendList = v
	return s
}

func (s *GetAttackedAssetDealResponseBodyData) Validate() error {
	if s.EcsTrendList != nil {
		for _, item := range s.EcsTrendList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAttackedAssetDealResponseBodyDataEcsTrendList struct {
	// Date point.
	//
	// example:
	//
	// 202312或20231205
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// Number of processed items.
	//
	// example:
	//
	// 2
	DealCount *int64 `json:"DealCount,omitempty" xml:"DealCount,omitempty"`
	// Number of discovered items.
	//
	// example:
	//
	// 暂时无值，有疑问请联系管理员
	FindCount *int64 `json:"FindCount,omitempty" xml:"FindCount,omitempty"`
}

func (s GetAttackedAssetDealResponseBodyDataEcsTrendList) String() string {
	return dara.Prettify(s)
}

func (s GetAttackedAssetDealResponseBodyDataEcsTrendList) GoString() string {
	return s.String()
}

func (s *GetAttackedAssetDealResponseBodyDataEcsTrendList) GetDate() *string {
	return s.Date
}

func (s *GetAttackedAssetDealResponseBodyDataEcsTrendList) GetDealCount() *int64 {
	return s.DealCount
}

func (s *GetAttackedAssetDealResponseBodyDataEcsTrendList) GetFindCount() *int64 {
	return s.FindCount
}

func (s *GetAttackedAssetDealResponseBodyDataEcsTrendList) SetDate(v string) *GetAttackedAssetDealResponseBodyDataEcsTrendList {
	s.Date = &v
	return s
}

func (s *GetAttackedAssetDealResponseBodyDataEcsTrendList) SetDealCount(v int64) *GetAttackedAssetDealResponseBodyDataEcsTrendList {
	s.DealCount = &v
	return s
}

func (s *GetAttackedAssetDealResponseBodyDataEcsTrendList) SetFindCount(v int64) *GetAttackedAssetDealResponseBodyDataEcsTrendList {
	s.FindCount = &v
	return s
}

func (s *GetAttackedAssetDealResponseBodyDataEcsTrendList) Validate() error {
	return dara.Validate(s)
}
