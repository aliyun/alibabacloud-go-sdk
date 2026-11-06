// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAppInstanceGroupReportListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAppInstanceGroupReportListResponseBody
	GetCode() *string
	SetIndexData(v []*QueryAppInstanceGroupReportListResponseBodyIndexData) *QueryAppInstanceGroupReportListResponseBody
	GetIndexData() []*QueryAppInstanceGroupReportListResponseBodyIndexData
	SetMessage(v string) *QueryAppInstanceGroupReportListResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *QueryAppInstanceGroupReportListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryAppInstanceGroupReportListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryAppInstanceGroupReportListResponseBody
	GetRequestId() *string
}

type QueryAppInstanceGroupReportListResponseBody struct {
	// The business status code. A value of `200` indicates that the request is processed. Other values indicate that the processing failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of app instance group report data.
	IndexData []*QueryAppInstanceGroupReportListResponseBodyIndexData `json:"IndexData,omitempty" xml:"IndexData,omitempty" type:"Repeated"`
	// The description of the request processing result. If the request is successful, `success` is usually returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number corresponding to the current response.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page used in the current response.
	//
	// example:
	//
	// 200
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique identifier of the current request, which can be used to troubleshoot request issues.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAppInstanceGroupReportListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAppInstanceGroupReportListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAppInstanceGroupReportListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAppInstanceGroupReportListResponseBody) GetIndexData() []*QueryAppInstanceGroupReportListResponseBodyIndexData {
	return s.IndexData
}

func (s *QueryAppInstanceGroupReportListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAppInstanceGroupReportListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryAppInstanceGroupReportListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryAppInstanceGroupReportListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAppInstanceGroupReportListResponseBody) SetCode(v string) *QueryAppInstanceGroupReportListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAppInstanceGroupReportListResponseBody) SetIndexData(v []*QueryAppInstanceGroupReportListResponseBodyIndexData) *QueryAppInstanceGroupReportListResponseBody {
	s.IndexData = v
	return s
}

func (s *QueryAppInstanceGroupReportListResponseBody) SetMessage(v string) *QueryAppInstanceGroupReportListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAppInstanceGroupReportListResponseBody) SetPageNumber(v int32) *QueryAppInstanceGroupReportListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryAppInstanceGroupReportListResponseBody) SetPageSize(v int32) *QueryAppInstanceGroupReportListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryAppInstanceGroupReportListResponseBody) SetRequestId(v string) *QueryAppInstanceGroupReportListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAppInstanceGroupReportListResponseBody) Validate() error {
	if s.IndexData != nil {
		for _, item := range s.IndexData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAppInstanceGroupReportListResponseBodyIndexData struct {
	// The ID of the app instance group.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The name of the app instance group.
	//
	// example:
	//
	// R&D workstation group
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The current status of the app instance group.
	//
	// Valid values:
	//
	// - `DEPLOYING`: Deploying.
	//
	// - `DEPLOYED`: Deployed but not published.
	//
	// - `PUBLISHED`: Published and available to provide services based on the current configurations.
	//
	// - `STOPPING`: Stopping.
	//
	// - `STOPPED`: Stopped.
	//
	// - `MAINTAINING`: Maintaining.
	//
	// - `FAILED`: Deployment or running failed.
	//
	// - `MAINTAIN_FAILED`: Maintenance failed.
	//
	// - `DELETING`: Deleting.
	//
	// - `DELETED`: Deleted.
	//
	// example:
	//
	// PUBLISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryAppInstanceGroupReportListResponseBodyIndexData) String() string {
	return dara.Prettify(s)
}

func (s QueryAppInstanceGroupReportListResponseBodyIndexData) GoString() string {
	return s.String()
}

func (s *QueryAppInstanceGroupReportListResponseBodyIndexData) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *QueryAppInstanceGroupReportListResponseBodyIndexData) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *QueryAppInstanceGroupReportListResponseBodyIndexData) GetStatus() *string {
	return s.Status
}

func (s *QueryAppInstanceGroupReportListResponseBodyIndexData) SetAppInstanceGroupId(v string) *QueryAppInstanceGroupReportListResponseBodyIndexData {
	s.AppInstanceGroupId = &v
	return s
}

func (s *QueryAppInstanceGroupReportListResponseBodyIndexData) SetAppInstanceGroupName(v string) *QueryAppInstanceGroupReportListResponseBodyIndexData {
	s.AppInstanceGroupName = &v
	return s
}

func (s *QueryAppInstanceGroupReportListResponseBodyIndexData) SetStatus(v string) *QueryAppInstanceGroupReportListResponseBodyIndexData {
	s.Status = &v
	return s
}

func (s *QueryAppInstanceGroupReportListResponseBodyIndexData) Validate() error {
	return dara.Validate(s)
}
