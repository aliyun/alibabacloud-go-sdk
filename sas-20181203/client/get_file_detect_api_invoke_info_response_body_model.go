// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileDetectApiInvokeInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetFileDetectApiInvokeInfoResponseBodyData) *GetFileDetectApiInvokeInfoResponseBody
	GetData() *GetFileDetectApiInvokeInfoResponseBodyData
	SetRequestId(v string) *GetFileDetectApiInvokeInfoResponseBody
	GetRequestId() *string
}

type GetFileDetectApiInvokeInfoResponseBody struct {
	// Returns the response body.
	Data *GetFileDetectApiInvokeInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9f368b6e-d60a-43c5-bd6f-c7087f2d****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFileDetectApiInvokeInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileDetectApiInvokeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileDetectApiInvokeInfoResponseBody) GetData() *GetFileDetectApiInvokeInfoResponseBodyData {
	return s.Data
}

func (s *GetFileDetectApiInvokeInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileDetectApiInvokeInfoResponseBody) SetData(v *GetFileDetectApiInvokeInfoResponseBodyData) *GetFileDetectApiInvokeInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetFileDetectApiInvokeInfoResponseBody) SetRequestId(v string) *GetFileDetectApiInvokeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileDetectApiInvokeInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFileDetectApiInvokeInfoResponseBodyData struct {
	// The total number of authorizations.
	//
	// example:
	//
	// 10000
	AuthCount *int64 `json:"AuthCount,omitempty" xml:"AuthCount,omitempty"`
	// The total number of authorizations(excluding trials).
	//
	// example:
	//
	// 20
	AuthCountInSaleVersion *int64 `json:"AuthCountInSaleVersion,omitempty" xml:"AuthCountInSaleVersion,omitempty"`
	// The timestamp of the expiration date of the authorization number.
	//
	// example:
	//
	// 1815753600000
	Expire *int64 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// The frequency of calls.
	//
	// example:
	//
	// 20
	FlowRate *int32 `json:"FlowRate,omitempty" xml:"FlowRate,omitempty"`
	// The number of authorizations used.
	//
	// example:
	//
	// 10
	InvokeCount *int64 `json:"InvokeCount,omitempty" xml:"InvokeCount,omitempty"`
	// The number of authorizations used(excluding trials).
	//
	// example:
	//
	// 20
	InvokeCountInSaleVersion *int64 `json:"InvokeCountInSaleVersion,omitempty" xml:"InvokeCountInSaleVersion,omitempty"`
	// The number of remaining authorizations.
	//
	// example:
	//
	// 900
	RemainAuthCount *int64 `json:"RemainAuthCount,omitempty" xml:"RemainAuthCount,omitempty"`
	// The Authorized Version. Valid values include:
	//
	// 	- **1:*	- trial version
	//
	// 	- **2:*	- Enterprise Edition
	//
	// example:
	//
	// 2
	SaleVersion *int32 `json:"SaleVersion,omitempty" xml:"SaleVersion,omitempty"`
	// The time unit of the frequency limit. Value:
	//
	// 	- **SECONDS**
	//
	// 	- **MINUTES**
	//
	// example:
	//
	// SECONDS
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s GetFileDetectApiInvokeInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFileDetectApiInvokeInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFileDetectApiInvokeInfoResponseBodyData) GetAuthCount() *int64 {
	return s.AuthCount
}

func (s *GetFileDetectApiInvokeInfoResponseBodyData) GetAuthCountInSaleVersion() *int64 {
	return s.AuthCountInSaleVersion
}

func (s *GetFileDetectApiInvokeInfoResponseBodyData) GetExpire() *int64 {
	return s.Expire
}

func (s *GetFileDetectApiInvokeInfoResponseBodyData) GetFlowRate() *int32 {
	return s.FlowRate
}

func (s *GetFileDetectApiInvokeInfoResponseBodyData) GetInvokeCount() *int64 {
	return s.InvokeCount
}

func (s *GetFileDetectApiInvokeInfoResponseBodyData) GetInvokeCountInSaleVersion() *int64 {
	return s.InvokeCountInSaleVersion
}

func (s *GetFileDetectApiInvokeInfoResponseBodyData) GetRemainAuthCount() *int64 {
	return s.RemainAuthCount
}

func (s *GetFileDetectApiInvokeInfoResponseBodyData) GetSaleVersion() *int32 {
	return s.SaleVersion
}

func (s *GetFileDetectApiInvokeInfoResponseBodyData) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *GetFileDetectApiInvokeInfoResponseBodyData) SetAuthCount(v int64) *GetFileDetectApiInvokeInfoResponseBodyData {
	s.AuthCount = &v
	return s
}

func (s *GetFileDetectApiInvokeInfoResponseBodyData) SetAuthCountInSaleVersion(v int64) *GetFileDetectApiInvokeInfoResponseBodyData {
	s.AuthCountInSaleVersion = &v
	return s
}

func (s *GetFileDetectApiInvokeInfoResponseBodyData) SetExpire(v int64) *GetFileDetectApiInvokeInfoResponseBodyData {
	s.Expire = &v
	return s
}

func (s *GetFileDetectApiInvokeInfoResponseBodyData) SetFlowRate(v int32) *GetFileDetectApiInvokeInfoResponseBodyData {
	s.FlowRate = &v
	return s
}

func (s *GetFileDetectApiInvokeInfoResponseBodyData) SetInvokeCount(v int64) *GetFileDetectApiInvokeInfoResponseBodyData {
	s.InvokeCount = &v
	return s
}

func (s *GetFileDetectApiInvokeInfoResponseBodyData) SetInvokeCountInSaleVersion(v int64) *GetFileDetectApiInvokeInfoResponseBodyData {
	s.InvokeCountInSaleVersion = &v
	return s
}

func (s *GetFileDetectApiInvokeInfoResponseBodyData) SetRemainAuthCount(v int64) *GetFileDetectApiInvokeInfoResponseBodyData {
	s.RemainAuthCount = &v
	return s
}

func (s *GetFileDetectApiInvokeInfoResponseBodyData) SetSaleVersion(v int32) *GetFileDetectApiInvokeInfoResponseBodyData {
	s.SaleVersion = &v
	return s
}

func (s *GetFileDetectApiInvokeInfoResponseBodyData) SetTimeUnit(v string) *GetFileDetectApiInvokeInfoResponseBodyData {
	s.TimeUnit = &v
	return s
}

func (s *GetFileDetectApiInvokeInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
