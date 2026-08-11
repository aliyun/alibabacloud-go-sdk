// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppDetailStatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetAiAppDetailStatRequest
	GetAppId() *string
	SetEndTime(v string) *GetAiAppDetailStatRequest
	GetEndTime() *string
	SetRegionId(v string) *GetAiAppDetailStatRequest
	GetRegionId() *string
	SetStartTime(v string) *GetAiAppDetailStatRequest
	GetStartTime() *string
}

type GetAiAppDetailStatRequest struct {
	// The ID of the AI application. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// id-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end time of the query.
	//
	// example:
	//
	// 2025-07-09 10:30:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the region where the application resides.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of the query.
	//
	// example:
	//
	// 2024-09-10 14:48:01
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAiAppDetailStatRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppDetailStatRequest) GoString() string {
	return s.String()
}

func (s *GetAiAppDetailStatRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetAiAppDetailStatRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetAiAppDetailStatRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAiAppDetailStatRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetAiAppDetailStatRequest) SetAppId(v string) *GetAiAppDetailStatRequest {
	s.AppId = &v
	return s
}

func (s *GetAiAppDetailStatRequest) SetEndTime(v string) *GetAiAppDetailStatRequest {
	s.EndTime = &v
	return s
}

func (s *GetAiAppDetailStatRequest) SetRegionId(v string) *GetAiAppDetailStatRequest {
	s.RegionId = &v
	return s
}

func (s *GetAiAppDetailStatRequest) SetStartTime(v string) *GetAiAppDetailStatRequest {
	s.StartTime = &v
	return s
}

func (s *GetAiAppDetailStatRequest) Validate() error {
	return dara.Validate(s)
}
