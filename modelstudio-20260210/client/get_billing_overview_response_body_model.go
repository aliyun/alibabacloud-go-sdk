// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBillingOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBillingOverviewResponseBody
	GetCode() *string
	SetData(v *GetBillingOverviewResponseBodyData) *GetBillingOverviewResponseBody
	GetData() *GetBillingOverviewResponseBodyData
	SetMessage(v string) *GetBillingOverviewResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBillingOverviewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBillingOverviewResponseBody
	GetSuccess() *bool
}

type GetBillingOverviewResponseBody struct {
	// The request result code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The bill overview data.
	Data *GetBillingOverviewResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request result message.
	//
	// example:
	//
	// null
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BB521414-5D38-5E66-AA66-963B2B4200E2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetBillingOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBillingOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetBillingOverviewResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBillingOverviewResponseBody) GetData() *GetBillingOverviewResponseBodyData {
	return s.Data
}

func (s *GetBillingOverviewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBillingOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBillingOverviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBillingOverviewResponseBody) SetCode(v string) *GetBillingOverviewResponseBody {
	s.Code = &v
	return s
}

func (s *GetBillingOverviewResponseBody) SetData(v *GetBillingOverviewResponseBodyData) *GetBillingOverviewResponseBody {
	s.Data = v
	return s
}

func (s *GetBillingOverviewResponseBody) SetMessage(v string) *GetBillingOverviewResponseBody {
	s.Message = &v
	return s
}

func (s *GetBillingOverviewResponseBody) SetRequestId(v string) *GetBillingOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBillingOverviewResponseBody) SetSuccess(v bool) *GetBillingOverviewResponseBody {
	s.Success = &v
	return s
}

func (s *GetBillingOverviewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBillingOverviewResponseBodyData struct {
	// The currency of the amount.
	//
	// example:
	//
	// USD
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// The top N groups sorted by amount in descending order.
	Groups []*GetBillingOverviewResponseBodyDataGroups `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
	// The total pretax amount.
	//
	// example:
	//
	// 28729.32
	PretaxAmount *string `json:"pretaxAmount,omitempty" xml:"pretaxAmount,omitempty"`
	// The total tax amount.
	//
	// example:
	//
	// 2499.28
	TaxAmount *string `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	// The total amount.
	//
	// example:
	//
	// 31228.60
	TotalAmount *string `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
}

func (s GetBillingOverviewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBillingOverviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBillingOverviewResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *GetBillingOverviewResponseBodyData) GetGroups() []*GetBillingOverviewResponseBodyDataGroups {
	return s.Groups
}

func (s *GetBillingOverviewResponseBodyData) GetPretaxAmount() *string {
	return s.PretaxAmount
}

func (s *GetBillingOverviewResponseBodyData) GetTaxAmount() *string {
	return s.TaxAmount
}

func (s *GetBillingOverviewResponseBodyData) GetTotalAmount() *string {
	return s.TotalAmount
}

func (s *GetBillingOverviewResponseBodyData) SetCurrency(v string) *GetBillingOverviewResponseBodyData {
	s.Currency = &v
	return s
}

func (s *GetBillingOverviewResponseBodyData) SetGroups(v []*GetBillingOverviewResponseBodyDataGroups) *GetBillingOverviewResponseBodyData {
	s.Groups = v
	return s
}

func (s *GetBillingOverviewResponseBodyData) SetPretaxAmount(v string) *GetBillingOverviewResponseBodyData {
	s.PretaxAmount = &v
	return s
}

func (s *GetBillingOverviewResponseBodyData) SetTaxAmount(v string) *GetBillingOverviewResponseBodyData {
	s.TaxAmount = &v
	return s
}

func (s *GetBillingOverviewResponseBodyData) SetTotalAmount(v string) *GetBillingOverviewResponseBodyData {
	s.TotalAmount = &v
	return s
}

func (s *GetBillingOverviewResponseBodyData) Validate() error {
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBillingOverviewResponseBodyDataGroups struct {
	// The amount of the current group.
	//
	// example:
	//
	// 3000
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// The list of commodity codes associated with the current group.
	ArticleCodes []*string `json:"articleCodes,omitempty" xml:"articleCodes,omitempty" type:"Repeated"`
	// The grouping dimension value. A null value is returned as DIMENSION_FILTER_NULL_VALUE.
	//
	// example:
	//
	// inference
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The display name of the group, which is affected by the locale parameter. A null value is displayed as -.
	//
	// example:
	//
	// Model invocation
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ratio of the group amount to the total amount of the top N groups.
	//
	// example:
	//
	// 0.10
	Percentage *string `json:"percentage,omitempty" xml:"percentage,omitempty"`
}

func (s GetBillingOverviewResponseBodyDataGroups) String() string {
	return dara.Prettify(s)
}

func (s GetBillingOverviewResponseBodyDataGroups) GoString() string {
	return s.String()
}

func (s *GetBillingOverviewResponseBodyDataGroups) GetAmount() *string {
	return s.Amount
}

func (s *GetBillingOverviewResponseBodyDataGroups) GetArticleCodes() []*string {
	return s.ArticleCodes
}

func (s *GetBillingOverviewResponseBodyDataGroups) GetKey() *string {
	return s.Key
}

func (s *GetBillingOverviewResponseBodyDataGroups) GetName() *string {
	return s.Name
}

func (s *GetBillingOverviewResponseBodyDataGroups) GetPercentage() *string {
	return s.Percentage
}

func (s *GetBillingOverviewResponseBodyDataGroups) SetAmount(v string) *GetBillingOverviewResponseBodyDataGroups {
	s.Amount = &v
	return s
}

func (s *GetBillingOverviewResponseBodyDataGroups) SetArticleCodes(v []*string) *GetBillingOverviewResponseBodyDataGroups {
	s.ArticleCodes = v
	return s
}

func (s *GetBillingOverviewResponseBodyDataGroups) SetKey(v string) *GetBillingOverviewResponseBodyDataGroups {
	s.Key = &v
	return s
}

func (s *GetBillingOverviewResponseBodyDataGroups) SetName(v string) *GetBillingOverviewResponseBodyDataGroups {
	s.Name = &v
	return s
}

func (s *GetBillingOverviewResponseBodyDataGroups) SetPercentage(v string) *GetBillingOverviewResponseBodyDataGroups {
	s.Percentage = &v
	return s
}

func (s *GetBillingOverviewResponseBodyDataGroups) Validate() error {
	return dara.Validate(s)
}
