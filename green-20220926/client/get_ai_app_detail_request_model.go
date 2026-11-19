// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetAiAppDetailRequest
	GetAppId() *string
	SetEndTime(v string) *GetAiAppDetailRequest
	GetEndTime() *string
	SetRegionId(v string) *GetAiAppDetailRequest
	GetRegionId() *string
	SetStartTime(v string) *GetAiAppDetailRequest
	GetStartTime() *string
}

type GetAiAppDetailRequest struct {
	// The AI application ID. This parameter is required. The AppId is derived from AI applications on integrated external AI platforms, such as Model Studio and Dify. Before calling this API, you must call InitAiAppScan to initiate the activation and then call ConfirmAiAppScan to confirm the activation of the AI application detection service. To obtain a valid AppId, you can call ListAiAppByPage to query the list of integrated applications.
	//
	// This parameter is required.
	//
	// example:
	//
	// id-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end time.
	//
	// example:
	//
	// 2026-01-02 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2026-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAiAppDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppDetailRequest) GoString() string {
	return s.String()
}

func (s *GetAiAppDetailRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetAiAppDetailRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetAiAppDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAiAppDetailRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetAiAppDetailRequest) SetAppId(v string) *GetAiAppDetailRequest {
	s.AppId = &v
	return s
}

func (s *GetAiAppDetailRequest) SetEndTime(v string) *GetAiAppDetailRequest {
	s.EndTime = &v
	return s
}

func (s *GetAiAppDetailRequest) SetRegionId(v string) *GetAiAppDetailRequest {
	s.RegionId = &v
	return s
}

func (s *GetAiAppDetailRequest) SetStartTime(v string) *GetAiAppDetailRequest {
	s.StartTime = &v
	return s
}

func (s *GetAiAppDetailRequest) Validate() error {
	return dara.Validate(s)
}
