// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeliveredSupplyItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDeliveredSupplyItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDeliveredSupplyItemsResponse
	GetStatusCode() *int32
	SetBody(v []*QueryDeliveredSupplyItemsResponseBody) *QueryDeliveredSupplyItemsResponse
	GetBody() []*QueryDeliveredSupplyItemsResponseBody
}

type QueryDeliveredSupplyItemsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*QueryDeliveredSupplyItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s QueryDeliveredSupplyItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDeliveredSupplyItemsResponse) GoString() string {
	return s.String()
}

func (s *QueryDeliveredSupplyItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDeliveredSupplyItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDeliveredSupplyItemsResponse) GetBody() []*QueryDeliveredSupplyItemsResponseBody {
	return s.Body
}

func (s *QueryDeliveredSupplyItemsResponse) SetHeaders(v map[string]*string) *QueryDeliveredSupplyItemsResponse {
	s.Headers = v
	return s
}

func (s *QueryDeliveredSupplyItemsResponse) SetStatusCode(v int32) *QueryDeliveredSupplyItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponse) SetBody(v []*QueryDeliveredSupplyItemsResponseBody) *QueryDeliveredSupplyItemsResponse {
	s.Body = v
	return s
}

func (s *QueryDeliveredSupplyItemsResponse) Validate() error {
	return dara.Validate(s)
}

type QueryDeliveredSupplyItemsResponseBody struct {
	AccountId         *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	CommodityTypeCode *string `json:"commodityTypeCode,omitempty" xml:"commodityTypeCode,omitempty"`
	DemandPlanId      *int64  `json:"demandPlanId,omitempty" xml:"demandPlanId,omitempty"`
	PlanTitle         *string `json:"planTitle,omitempty" xml:"planTitle,omitempty"`
	Region            *string `json:"region,omitempty" xml:"region,omitempty"`
	Zone              *string `json:"zone,omitempty" xml:"zone,omitempty"`
	CommodityCode     *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	DemandCount       *int32  `json:"demandCount,omitempty" xml:"demandCount,omitempty"`
	DeliveredAmount   *int32  `json:"deliveredAmount,omitempty" xml:"deliveredAmount,omitempty"`
	OpenCount         *int32  `json:"openCount,omitempty" xml:"openCount,omitempty"`
}

func (s QueryDeliveredSupplyItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDeliveredSupplyItemsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeliveredSupplyItemsResponseBody) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryDeliveredSupplyItemsResponseBody) GetCommodityTypeCode() *string {
	return s.CommodityTypeCode
}

func (s *QueryDeliveredSupplyItemsResponseBody) GetDemandPlanId() *int64 {
	return s.DemandPlanId
}

func (s *QueryDeliveredSupplyItemsResponseBody) GetPlanTitle() *string {
	return s.PlanTitle
}

func (s *QueryDeliveredSupplyItemsResponseBody) GetRegion() *string {
	return s.Region
}

func (s *QueryDeliveredSupplyItemsResponseBody) GetZone() *string {
	return s.Zone
}

func (s *QueryDeliveredSupplyItemsResponseBody) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *QueryDeliveredSupplyItemsResponseBody) GetDemandCount() *int32 {
	return s.DemandCount
}

func (s *QueryDeliveredSupplyItemsResponseBody) GetDeliveredAmount() *int32 {
	return s.DeliveredAmount
}

func (s *QueryDeliveredSupplyItemsResponseBody) GetOpenCount() *int32 {
	return s.OpenCount
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetAccountId(v string) *QueryDeliveredSupplyItemsResponseBody {
	s.AccountId = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetCommodityTypeCode(v string) *QueryDeliveredSupplyItemsResponseBody {
	s.CommodityTypeCode = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetDemandPlanId(v int64) *QueryDeliveredSupplyItemsResponseBody {
	s.DemandPlanId = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetPlanTitle(v string) *QueryDeliveredSupplyItemsResponseBody {
	s.PlanTitle = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetRegion(v string) *QueryDeliveredSupplyItemsResponseBody {
	s.Region = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetZone(v string) *QueryDeliveredSupplyItemsResponseBody {
	s.Zone = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetCommodityCode(v string) *QueryDeliveredSupplyItemsResponseBody {
	s.CommodityCode = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetDemandCount(v int32) *QueryDeliveredSupplyItemsResponseBody {
	s.DemandCount = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetDeliveredAmount(v int32) *QueryDeliveredSupplyItemsResponseBody {
	s.DeliveredAmount = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetOpenCount(v int32) *QueryDeliveredSupplyItemsResponseBody {
	s.OpenCount = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponseBody) Validate() error {
	return dara.Validate(s)
}
