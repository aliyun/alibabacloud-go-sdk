// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetResourceOverviewResponseBody
	GetCode() *string
	SetData(v *GetResourceOverviewResponseBodyData) *GetResourceOverviewResponseBody
	GetData() *GetResourceOverviewResponseBodyData
	SetMessage(v string) *GetResourceOverviewResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetResourceOverviewResponseBody
	GetRequestId() *string
}

type GetResourceOverviewResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Resource information.
	Data *GetResourceOverviewResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// DD19A442-93C5-5C97-AFA0-B9C57EBD781B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetResourceOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceOverviewResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetResourceOverviewResponseBody) GetData() *GetResourceOverviewResponseBodyData {
	return s.Data
}

func (s *GetResourceOverviewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetResourceOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceOverviewResponseBody) SetCode(v string) *GetResourceOverviewResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourceOverviewResponseBody) SetData(v *GetResourceOverviewResponseBodyData) *GetResourceOverviewResponseBody {
	s.Data = v
	return s
}

func (s *GetResourceOverviewResponseBody) SetMessage(v string) *GetResourceOverviewResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourceOverviewResponseBody) SetRequestId(v string) *GetResourceOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceOverviewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceOverviewResponseBodyData struct {
	// API information.
	Api *GetResourceOverviewResponseBodyDataApi `json:"api,omitempty" xml:"api,omitempty" type:"Struct"`
	// Gateway information.
	Gateway      *GetResourceOverviewResponseBodyDataGateway        `json:"gateway,omitempty" xml:"gateway,omitempty" type:"Struct"`
	RiskOverview []*GetResourceOverviewResponseBodyDataRiskOverview `json:"riskOverview,omitempty" xml:"riskOverview,omitempty" type:"Repeated"`
}

func (s GetResourceOverviewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetResourceOverviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResourceOverviewResponseBodyData) GetApi() *GetResourceOverviewResponseBodyDataApi {
	return s.Api
}

func (s *GetResourceOverviewResponseBodyData) GetGateway() *GetResourceOverviewResponseBodyDataGateway {
	return s.Gateway
}

func (s *GetResourceOverviewResponseBodyData) GetRiskOverview() []*GetResourceOverviewResponseBodyDataRiskOverview {
	return s.RiskOverview
}

func (s *GetResourceOverviewResponseBodyData) SetApi(v *GetResourceOverviewResponseBodyDataApi) *GetResourceOverviewResponseBodyData {
	s.Api = v
	return s
}

func (s *GetResourceOverviewResponseBodyData) SetGateway(v *GetResourceOverviewResponseBodyDataGateway) *GetResourceOverviewResponseBodyData {
	s.Gateway = v
	return s
}

func (s *GetResourceOverviewResponseBodyData) SetRiskOverview(v []*GetResourceOverviewResponseBodyDataRiskOverview) *GetResourceOverviewResponseBodyData {
	s.RiskOverview = v
	return s
}

func (s *GetResourceOverviewResponseBodyData) Validate() error {
	if s.Api != nil {
		if err := s.Api.Validate(); err != nil {
			return err
		}
	}
	if s.Gateway != nil {
		if err := s.Gateway.Validate(); err != nil {
			return err
		}
	}
	if s.RiskOverview != nil {
		for _, item := range s.RiskOverview {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResourceOverviewResponseBodyDataApi struct {
	// Number of published APIs.
	//
	// example:
	//
	// 1
	PublishedCount *int64 `json:"publishedCount,omitempty" xml:"publishedCount,omitempty"`
	// Number of APIs.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetResourceOverviewResponseBodyDataApi) String() string {
	return dara.Prettify(s)
}

func (s GetResourceOverviewResponseBodyDataApi) GoString() string {
	return s.String()
}

func (s *GetResourceOverviewResponseBodyDataApi) GetPublishedCount() *int64 {
	return s.PublishedCount
}

func (s *GetResourceOverviewResponseBodyDataApi) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetResourceOverviewResponseBodyDataApi) SetPublishedCount(v int64) *GetResourceOverviewResponseBodyDataApi {
	s.PublishedCount = &v
	return s
}

func (s *GetResourceOverviewResponseBodyDataApi) SetTotalCount(v int64) *GetResourceOverviewResponseBodyDataApi {
	s.TotalCount = &v
	return s
}

func (s *GetResourceOverviewResponseBodyDataApi) Validate() error {
	return dara.Validate(s)
}

type GetResourceOverviewResponseBodyDataGateway struct {
	// Number of running gateways.
	//
	// example:
	//
	// 1
	RunningCount *int64 `json:"runningCount,omitempty" xml:"runningCount,omitempty"`
	// Number of gateway instances.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetResourceOverviewResponseBodyDataGateway) String() string {
	return dara.Prettify(s)
}

func (s GetResourceOverviewResponseBodyDataGateway) GoString() string {
	return s.String()
}

func (s *GetResourceOverviewResponseBodyDataGateway) GetRunningCount() *int64 {
	return s.RunningCount
}

func (s *GetResourceOverviewResponseBodyDataGateway) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetResourceOverviewResponseBodyDataGateway) SetRunningCount(v int64) *GetResourceOverviewResponseBodyDataGateway {
	s.RunningCount = &v
	return s
}

func (s *GetResourceOverviewResponseBodyDataGateway) SetTotalCount(v int64) *GetResourceOverviewResponseBodyDataGateway {
	s.TotalCount = &v
	return s
}

func (s *GetResourceOverviewResponseBodyDataGateway) Validate() error {
	return dara.Validate(s)
}

type GetResourceOverviewResponseBodyDataRiskOverview struct {
	// example:
	//
	// 1
	Count       *string                                                       `json:"count,omitempty" xml:"count,omitempty"`
	RiskDetails []*GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails `json:"riskDetails,omitempty" xml:"riskDetails,omitempty" type:"Repeated"`
	// example:
	//
	// LOW
	RiskLevel *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
}

func (s GetResourceOverviewResponseBodyDataRiskOverview) String() string {
	return dara.Prettify(s)
}

func (s GetResourceOverviewResponseBodyDataRiskOverview) GoString() string {
	return s.String()
}

func (s *GetResourceOverviewResponseBodyDataRiskOverview) GetCount() *string {
	return s.Count
}

func (s *GetResourceOverviewResponseBodyDataRiskOverview) GetRiskDetails() []*GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails {
	return s.RiskDetails
}

func (s *GetResourceOverviewResponseBodyDataRiskOverview) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *GetResourceOverviewResponseBodyDataRiskOverview) SetCount(v string) *GetResourceOverviewResponseBodyDataRiskOverview {
	s.Count = &v
	return s
}

func (s *GetResourceOverviewResponseBodyDataRiskOverview) SetRiskDetails(v []*GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails) *GetResourceOverviewResponseBodyDataRiskOverview {
	s.RiskDetails = v
	return s
}

func (s *GetResourceOverviewResponseBodyDataRiskOverview) SetRiskLevel(v string) *GetResourceOverviewResponseBodyDataRiskOverview {
	s.RiskLevel = &v
	return s
}

func (s *GetResourceOverviewResponseBodyDataRiskOverview) Validate() error {
	if s.RiskDetails != nil {
		for _, item := range s.RiskDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails struct {
	// example:
	//
	// gw-xxxxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// test-gateway
	GatewayName *string `json:"gatewayName,omitempty" xml:"gatewayName,omitempty"`
	// example:
	//
	// LOW
	RiskLevel *string   `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	RiskNames []*string `json:"riskNames,omitempty" xml:"riskNames,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Score *string `json:"score,omitempty" xml:"score,omitempty"`
}

func (s GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails) String() string {
	return dara.Prettify(s)
}

func (s GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails) GoString() string {
	return s.String()
}

func (s *GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails) GetGatewayName() *string {
	return s.GatewayName
}

func (s *GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails) GetRiskNames() []*string {
	return s.RiskNames
}

func (s *GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails) GetScore() *string {
	return s.Score
}

func (s *GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails) SetGatewayId(v string) *GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails {
	s.GatewayId = &v
	return s
}

func (s *GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails) SetGatewayName(v string) *GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails {
	s.GatewayName = &v
	return s
}

func (s *GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails) SetRiskLevel(v string) *GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails {
	s.RiskLevel = &v
	return s
}

func (s *GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails) SetRiskNames(v []*string) *GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails {
	s.RiskNames = v
	return s
}

func (s *GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails) SetScore(v string) *GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails {
	s.Score = &v
	return s
}

func (s *GetResourceOverviewResponseBodyDataRiskOverviewRiskDetails) Validate() error {
	return dara.Validate(s)
}
