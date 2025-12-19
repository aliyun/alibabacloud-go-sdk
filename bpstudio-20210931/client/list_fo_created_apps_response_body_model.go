// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFoCreatedAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFoCreatedAppsResponseBody
	GetCode() *string
	SetData(v []*ListFoCreatedAppsResponseBodyData) *ListFoCreatedAppsResponseBody
	GetData() []*ListFoCreatedAppsResponseBodyData
	SetMessage(v string) *ListFoCreatedAppsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListFoCreatedAppsResponseBody
	GetRequestId() *string
}

type ListFoCreatedAppsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The templates.
	Data []*ListFoCreatedAppsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned error message.
	//
	// example:
	//
	// Cannot find region according to your domain.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40F63F07-3AB6-53B3-8825-0580C130E3EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFoCreatedAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFoCreatedAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFoCreatedAppsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFoCreatedAppsResponseBody) GetData() []*ListFoCreatedAppsResponseBodyData {
	return s.Data
}

func (s *ListFoCreatedAppsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFoCreatedAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFoCreatedAppsResponseBody) SetCode(v string) *ListFoCreatedAppsResponseBody {
	s.Code = &v
	return s
}

func (s *ListFoCreatedAppsResponseBody) SetData(v []*ListFoCreatedAppsResponseBodyData) *ListFoCreatedAppsResponseBody {
	s.Data = v
	return s
}

func (s *ListFoCreatedAppsResponseBody) SetMessage(v string) *ListFoCreatedAppsResponseBody {
	s.Message = &v
	return s
}

func (s *ListFoCreatedAppsResponseBody) SetRequestId(v string) *ListFoCreatedAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFoCreatedAppsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFoCreatedAppsResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// JIX9NEZUALGS46UI
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The URL of the error report.
	//
	// example:
	//
	// https://api.aliyun.com/troubleshoot?q=ServiceUnavailable&product=BPStudio&requestId=4CDA03A3-C652-1408-8ABD-7E652A7CBFB6
	ReportUrl *string `json:"ReportUrl,omitempty" xml:"ReportUrl,omitempty"`
	// The state of the application.
	//
	// example:
	//
	// Deployed_Success
	//
	// Destroyed_Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The title.
	//
	// example:
	//
	// 容灾计划1
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListFoCreatedAppsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFoCreatedAppsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFoCreatedAppsResponseBodyData) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListFoCreatedAppsResponseBodyData) GetReportUrl() *string {
	return s.ReportUrl
}

func (s *ListFoCreatedAppsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListFoCreatedAppsResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *ListFoCreatedAppsResponseBodyData) SetApplicationId(v string) *ListFoCreatedAppsResponseBodyData {
	s.ApplicationId = &v
	return s
}

func (s *ListFoCreatedAppsResponseBodyData) SetReportUrl(v string) *ListFoCreatedAppsResponseBodyData {
	s.ReportUrl = &v
	return s
}

func (s *ListFoCreatedAppsResponseBodyData) SetStatus(v string) *ListFoCreatedAppsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListFoCreatedAppsResponseBodyData) SetTitle(v string) *ListFoCreatedAppsResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListFoCreatedAppsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
