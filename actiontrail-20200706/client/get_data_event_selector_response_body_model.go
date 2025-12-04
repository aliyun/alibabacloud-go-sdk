// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataEventSelectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataEventSelectors(v string) *GetDataEventSelectorResponseBody
	GetDataEventSelectors() *string
	SetIsTrailAllRegion(v bool) *GetDataEventSelectorResponseBody
	GetIsTrailAllRegion() *bool
	SetRequestId(v string) *GetDataEventSelectorResponseBody
	GetRequestId() *string
	SetSlsDeliveryConfigs(v []*GetDataEventSelectorResponseBodySlsDeliveryConfigs) *GetDataEventSelectorResponseBody
	GetSlsDeliveryConfigs() []*GetDataEventSelectorResponseBodySlsDeliveryConfigs
	SetTrailArn(v string) *GetDataEventSelectorResponseBody
	GetTrailArn() *string
}

type GetDataEventSelectorResponseBody struct {
	// example:
	//
	// [{"EventName":{"Equals":["GetObject","CopyObject","AppendObject"]},"ReadWriteType":"All","ServiceName":"Oss"}]
	DataEventSelectors *string `json:"DataEventSelectors,omitempty" xml:"DataEventSelectors,omitempty"`
	// example:
	//
	// true
	IsTrailAllRegion *bool `json:"IsTrailAllRegion,omitempty" xml:"IsTrailAllRegion,omitempty"`
	// example:
	//
	// 90771C32-635B-529C-950C-75A9607D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// This parameter is required.
	SlsDeliveryConfigs []*GetDataEventSelectorResponseBodySlsDeliveryConfigs `json:"SlsDeliveryConfigs,omitempty" xml:"SlsDeliveryConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// acs:actiontrail:cn-shanghai:159498693826****:trail/trail-name
	TrailArn *string `json:"TrailArn,omitempty" xml:"TrailArn,omitempty"`
}

func (s GetDataEventSelectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataEventSelectorResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataEventSelectorResponseBody) GetDataEventSelectors() *string {
	return s.DataEventSelectors
}

func (s *GetDataEventSelectorResponseBody) GetIsTrailAllRegion() *bool {
	return s.IsTrailAllRegion
}

func (s *GetDataEventSelectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataEventSelectorResponseBody) GetSlsDeliveryConfigs() []*GetDataEventSelectorResponseBodySlsDeliveryConfigs {
	return s.SlsDeliveryConfigs
}

func (s *GetDataEventSelectorResponseBody) GetTrailArn() *string {
	return s.TrailArn
}

func (s *GetDataEventSelectorResponseBody) SetDataEventSelectors(v string) *GetDataEventSelectorResponseBody {
	s.DataEventSelectors = &v
	return s
}

func (s *GetDataEventSelectorResponseBody) SetIsTrailAllRegion(v bool) *GetDataEventSelectorResponseBody {
	s.IsTrailAllRegion = &v
	return s
}

func (s *GetDataEventSelectorResponseBody) SetRequestId(v string) *GetDataEventSelectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataEventSelectorResponseBody) SetSlsDeliveryConfigs(v []*GetDataEventSelectorResponseBodySlsDeliveryConfigs) *GetDataEventSelectorResponseBody {
	s.SlsDeliveryConfigs = v
	return s
}

func (s *GetDataEventSelectorResponseBody) SetTrailArn(v string) *GetDataEventSelectorResponseBody {
	s.TrailArn = &v
	return s
}

func (s *GetDataEventSelectorResponseBody) Validate() error {
	if s.SlsDeliveryConfigs != nil {
		for _, item := range s.SlsDeliveryConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataEventSelectorResponseBodySlsDeliveryConfigs struct {
	// example:
	//
	// 2024-12-18T03:25:36Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// LogServiceException
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// acs:log:cn-shanghai:159498693826****:project/actiontrail-log-159498693826****-cn-shanghai
	RegionSlsProjectArn *string `json:"RegionSlsProjectArn,omitempty" xml:"RegionSlsProjectArn,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// cn-shanghai
	TrailRegion *string `json:"TrailRegion,omitempty" xml:"TrailRegion,omitempty"`
}

func (s GetDataEventSelectorResponseBodySlsDeliveryConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetDataEventSelectorResponseBodySlsDeliveryConfigs) GoString() string {
	return s.String()
}

func (s *GetDataEventSelectorResponseBodySlsDeliveryConfigs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetDataEventSelectorResponseBodySlsDeliveryConfigs) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataEventSelectorResponseBodySlsDeliveryConfigs) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataEventSelectorResponseBodySlsDeliveryConfigs) GetRegionSlsProjectArn() *string {
	return s.RegionSlsProjectArn
}

func (s *GetDataEventSelectorResponseBodySlsDeliveryConfigs) GetStatus() *string {
	return s.Status
}

func (s *GetDataEventSelectorResponseBodySlsDeliveryConfigs) GetTrailRegion() *string {
	return s.TrailRegion
}

func (s *GetDataEventSelectorResponseBodySlsDeliveryConfigs) SetCreateTime(v string) *GetDataEventSelectorResponseBodySlsDeliveryConfigs {
	s.CreateTime = &v
	return s
}

func (s *GetDataEventSelectorResponseBodySlsDeliveryConfigs) SetErrorCode(v string) *GetDataEventSelectorResponseBodySlsDeliveryConfigs {
	s.ErrorCode = &v
	return s
}

func (s *GetDataEventSelectorResponseBodySlsDeliveryConfigs) SetErrorMessage(v string) *GetDataEventSelectorResponseBodySlsDeliveryConfigs {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataEventSelectorResponseBodySlsDeliveryConfigs) SetRegionSlsProjectArn(v string) *GetDataEventSelectorResponseBodySlsDeliveryConfigs {
	s.RegionSlsProjectArn = &v
	return s
}

func (s *GetDataEventSelectorResponseBodySlsDeliveryConfigs) SetStatus(v string) *GetDataEventSelectorResponseBodySlsDeliveryConfigs {
	s.Status = &v
	return s
}

func (s *GetDataEventSelectorResponseBodySlsDeliveryConfigs) SetTrailRegion(v string) *GetDataEventSelectorResponseBodySlsDeliveryConfigs {
	s.TrailRegion = &v
	return s
}

func (s *GetDataEventSelectorResponseBodySlsDeliveryConfigs) Validate() error {
	return dara.Validate(s)
}
