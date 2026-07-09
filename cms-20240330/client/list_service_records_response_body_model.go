// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServiceRecordsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceRecordsResponseBody
	GetNextToken() *string
	SetRecords(v []*ListServiceRecordsResponseBodyRecords) *ListServiceRecordsResponseBody
	GetRecords() []*ListServiceRecordsResponseBodyRecords
	SetRequestId(v string) *ListServiceRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListServiceRecordsResponseBody
	GetTotalCount() *int32
}

type ListServiceRecordsResponseBody struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// 2-ba4d-4b9f-aa24-dcb067a30f1c
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The list of ticket operation records.
	Records []*ListServiceRecordsResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 0CEC5375-XXXX-XXXX-XXXX-9A629907C1F0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListServiceRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceRecordsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceRecordsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceRecordsResponseBody) GetRecords() []*ListServiceRecordsResponseBodyRecords {
	return s.Records
}

func (s *ListServiceRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceRecordsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListServiceRecordsResponseBody) SetMaxResults(v int32) *ListServiceRecordsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceRecordsResponseBody) SetNextToken(v string) *ListServiceRecordsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceRecordsResponseBody) SetRecords(v []*ListServiceRecordsResponseBodyRecords) *ListServiceRecordsResponseBody {
	s.Records = v
	return s
}

func (s *ListServiceRecordsResponseBody) SetRequestId(v string) *ListServiceRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceRecordsResponseBody) SetTotalCount(v int32) *ListServiceRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceRecordsResponseBody) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceRecordsResponseBodyRecords struct {
	// The entry content in JSON string format. The format varies depending on the recordType.
	//
	// example:
	//
	// {
	//
	//   "project": "proj-xtrace-16c988dcfe21fcb73c5e6f234927d998-cn-hangzhou",
	//
	//   "storeName": "app-biz-log",
	//
	//   "regionId": "cn-hangzhou",
	//
	//   "bindType": "logstore",
	//
	//   "traceIdRelateField": ""
	//
	// }
	RecordContent *string `json:"recordContent,omitempty" xml:"recordContent,omitempty"`
	// The type of the linked entry. Currently supported:
	//
	// logCorrelation, which indicates application log association.
	//
	// example:
	//
	// logCorrelation
	RecordType *string `json:"recordType,omitempty" xml:"recordType,omitempty"`
	// The unique identifier of the service.
	//
	// example:
	//
	// gaddp9ap8q@f8ca37734da3eda787dbb
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The workspace.
	//
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListServiceRecordsResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s ListServiceRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListServiceRecordsResponseBodyRecords) GetRecordContent() *string {
	return s.RecordContent
}

func (s *ListServiceRecordsResponseBodyRecords) GetRecordType() *string {
	return s.RecordType
}

func (s *ListServiceRecordsResponseBodyRecords) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListServiceRecordsResponseBodyRecords) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListServiceRecordsResponseBodyRecords) SetRecordContent(v string) *ListServiceRecordsResponseBodyRecords {
	s.RecordContent = &v
	return s
}

func (s *ListServiceRecordsResponseBodyRecords) SetRecordType(v string) *ListServiceRecordsResponseBodyRecords {
	s.RecordType = &v
	return s
}

func (s *ListServiceRecordsResponseBodyRecords) SetServiceId(v string) *ListServiceRecordsResponseBodyRecords {
	s.ServiceId = &v
	return s
}

func (s *ListServiceRecordsResponseBodyRecords) SetWorkspace(v string) *ListServiceRecordsResponseBodyRecords {
	s.Workspace = &v
	return s
}

func (s *ListServiceRecordsResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}
