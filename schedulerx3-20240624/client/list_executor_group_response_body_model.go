// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutorGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListExecutorGroupResponseBody
	GetCode() *int32
	SetData(v *ListExecutorGroupResponseBodyData) *ListExecutorGroupResponseBody
	GetData() *ListExecutorGroupResponseBodyData
	SetMaxResults(v int32) *ListExecutorGroupResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListExecutorGroupResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListExecutorGroupResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListExecutorGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListExecutorGroupResponseBody
	GetSuccess() *bool
}

type ListExecutorGroupResponseBody struct {
	// example:
	//
	// 200
	Code *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListExecutorGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// eCKqVlS5FKF5EWGGOo8EgQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5EF879D0-3B43-5AD1-9BF7-52418F9C5E73
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListExecutorGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListExecutorGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListExecutorGroupResponseBody) GetData() *ListExecutorGroupResponseBodyData {
	return s.Data
}

func (s *ListExecutorGroupResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExecutorGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListExecutorGroupResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExecutorGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExecutorGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListExecutorGroupResponseBody) SetCode(v int32) *ListExecutorGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListExecutorGroupResponseBody) SetData(v *ListExecutorGroupResponseBodyData) *ListExecutorGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListExecutorGroupResponseBody) SetMaxResults(v int32) *ListExecutorGroupResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListExecutorGroupResponseBody) SetMessage(v string) *ListExecutorGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListExecutorGroupResponseBody) SetNextToken(v string) *ListExecutorGroupResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExecutorGroupResponseBody) SetRequestId(v string) *ListExecutorGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExecutorGroupResponseBody) SetSuccess(v bool) *ListExecutorGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ListExecutorGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListExecutorGroupResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// -
	Records []*ListExecutorGroupResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListExecutorGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExecutorGroupResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListExecutorGroupResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExecutorGroupResponseBodyData) GetRecords() []*ListExecutorGroupResponseBodyDataRecords {
	return s.Records
}

func (s *ListExecutorGroupResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListExecutorGroupResponseBodyData) SetPageNumber(v int32) *ListExecutorGroupResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListExecutorGroupResponseBodyData) SetPageSize(v int32) *ListExecutorGroupResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListExecutorGroupResponseBodyData) SetRecords(v []*ListExecutorGroupResponseBodyDataRecords) *ListExecutorGroupResponseBodyData {
	s.Records = v
	return s
}

func (s *ListExecutorGroupResponseBodyData) SetTotal(v int32) *ListExecutorGroupResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListExecutorGroupResponseBodyData) Validate() error {
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

type ListExecutorGroupResponseBodyDataRecords struct {
	// example:
	//
	// f/HElmWgOgmb0mlbRRkNuQ==
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// api_key
	AuthType *string                                             `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	CiteList []*ListExecutorGroupResponseBodyDataRecordsCiteList `json:"CiteList,omitempty" xml:"CiteList,omitempty" type:"Repeated"`
	// example:
	//
	// my first workflow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// job01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// public
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// example:
	//
	// openai
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 12
	WorkerId *int64 `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
	// example:
	//
	// k8s_service
	WorkerType *string `json:"WorkerType,omitempty" xml:"WorkerType,omitempty"`
	// example:
	//
	// [{"address":"http://47.111.188.191:18789"}]
	Workers *string `json:"Workers,omitempty" xml:"Workers,omitempty"`
}

func (s ListExecutorGroupResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorGroupResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetApiKey() *string {
	return s.ApiKey
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetAuthType() *string {
	return s.AuthType
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetCiteList() []*ListExecutorGroupResponseBodyDataRecordsCiteList {
	return s.CiteList
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetDescription() *string {
	return s.Description
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetName() *string {
	return s.Name
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetNetwork() *string {
	return s.Network
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetProtocol() *string {
	return s.Protocol
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetWorkerId() *int64 {
	return s.WorkerId
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetWorkerType() *string {
	return s.WorkerType
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetWorkers() *string {
	return s.Workers
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetApiKey(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.ApiKey = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetAuthType(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.AuthType = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetCiteList(v []*ListExecutorGroupResponseBodyDataRecordsCiteList) *ListExecutorGroupResponseBodyDataRecords {
	s.CiteList = v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetDescription(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.Description = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetName(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.Name = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetNetwork(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.Network = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetProtocol(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.Protocol = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetWorkerId(v int64) *ListExecutorGroupResponseBodyDataRecords {
	s.WorkerId = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetWorkerType(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.WorkerType = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetWorkers(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.Workers = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) Validate() error {
	if s.CiteList != nil {
		for _, item := range s.CiteList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListExecutorGroupResponseBodyDataRecordsCiteList struct {
	// example:
	//
	// Tornado
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// group1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ListExecutorGroupResponseBodyDataRecordsCiteList) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorGroupResponseBodyDataRecordsCiteList) GoString() string {
	return s.String()
}

func (s *ListExecutorGroupResponseBodyDataRecordsCiteList) GetAppName() *string {
	return s.AppName
}

func (s *ListExecutorGroupResponseBodyDataRecordsCiteList) GetDescription() *string {
	return s.Description
}

func (s *ListExecutorGroupResponseBodyDataRecordsCiteList) SetAppName(v string) *ListExecutorGroupResponseBodyDataRecordsCiteList {
	s.AppName = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecordsCiteList) SetDescription(v string) *ListExecutorGroupResponseBodyDataRecordsCiteList {
	s.Description = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecordsCiteList) Validate() error {
	return dara.Validate(s)
}
