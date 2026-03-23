// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageVisitDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *GetPageVisitDataRequest
	GetAppCode() *string
	SetAppName(v string) *GetPageVisitDataRequest
	GetAppName() *string
	SetEndTime(v string) *GetPageVisitDataRequest
	GetEndTime() *string
	SetPId(v string) *GetPageVisitDataRequest
	GetPId() *string
	SetStartTime(v string) *GetPageVisitDataRequest
	GetStartTime() *string
}

type GetPageVisitDataRequest struct {
	// appCode
	//
	// example:
	//
	// 3c0837d5-e65b-11ec-9985-02420bb080c6
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// appName
	//
	// example:
	//
	// CLOUD_NETWORK
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// endTime
	//
	// example:
	//
	// 2023-07-11
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// pId
	//
	// example:
	//
	// 19048
	PId *string `json:"PId,omitempty" xml:"PId,omitempty"`
	// startTime
	//
	// example:
	//
	// 2022-11-22
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetPageVisitDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPageVisitDataRequest) GoString() string {
	return s.String()
}

func (s *GetPageVisitDataRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetPageVisitDataRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetPageVisitDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetPageVisitDataRequest) GetPId() *string {
	return s.PId
}

func (s *GetPageVisitDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetPageVisitDataRequest) SetAppCode(v string) *GetPageVisitDataRequest {
	s.AppCode = &v
	return s
}

func (s *GetPageVisitDataRequest) SetAppName(v string) *GetPageVisitDataRequest {
	s.AppName = &v
	return s
}

func (s *GetPageVisitDataRequest) SetEndTime(v string) *GetPageVisitDataRequest {
	s.EndTime = &v
	return s
}

func (s *GetPageVisitDataRequest) SetPId(v string) *GetPageVisitDataRequest {
	s.PId = &v
	return s
}

func (s *GetPageVisitDataRequest) SetStartTime(v string) *GetPageVisitDataRequest {
	s.StartTime = &v
	return s
}

func (s *GetPageVisitDataRequest) Validate() error {
	return dara.Validate(s)
}
