// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListDatasourcesResponseBody
	GetCode() *int32
	SetData(v *ListDatasourcesResponseBodyData) *ListDatasourcesResponseBody
	GetData() *ListDatasourcesResponseBodyData
	SetMessage(v string) *ListDatasourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDatasourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDatasourcesResponseBody
	GetSuccess() *bool
}

type ListDatasourcesResponseBody struct {
	// example:
	//
	// 200
	Code *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListDatasourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BAC1ADB5-EEB5-5834-93D8-522E067AF8D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDatasourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatasourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasourcesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListDatasourcesResponseBody) GetData() *ListDatasourcesResponseBodyData {
	return s.Data
}

func (s *ListDatasourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDatasourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatasourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDatasourcesResponseBody) SetCode(v int32) *ListDatasourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListDatasourcesResponseBody) SetData(v *ListDatasourcesResponseBodyData) *ListDatasourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListDatasourcesResponseBody) SetMessage(v string) *ListDatasourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListDatasourcesResponseBody) SetRequestId(v string) *ListDatasourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasourcesResponseBody) SetSuccess(v bool) *ListDatasourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListDatasourcesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDatasourcesResponseBodyData struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// eCKqVlS5FKF5EWGGOo8EgQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// -
	Records []*ListDatasourcesResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 30
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDatasourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDatasourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDatasourcesResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDatasourcesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDatasourcesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatasourcesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasourcesResponseBodyData) GetRecords() []*ListDatasourcesResponseBodyDataRecords {
	return s.Records
}

func (s *ListDatasourcesResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListDatasourcesResponseBodyData) SetMaxResults(v int32) *ListDatasourcesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListDatasourcesResponseBodyData) SetNextToken(v string) *ListDatasourcesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListDatasourcesResponseBodyData) SetPageNumber(v int32) *ListDatasourcesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDatasourcesResponseBodyData) SetPageSize(v int32) *ListDatasourcesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDatasourcesResponseBodyData) SetRecords(v []*ListDatasourcesResponseBodyDataRecords) *ListDatasourcesResponseBodyData {
	s.Records = v
	return s
}

func (s *ListDatasourcesResponseBodyData) SetTotal(v int32) *ListDatasourcesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListDatasourcesResponseBodyData) Validate() error {
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

type ListDatasourcesResponseBodyDataRecords struct {
	// example:
	//
	// {"host":"rm-bp1f03mxxxxx.mysql.rds.aliyuncs.com","port":3306,"userName":"test01","database":"test01","other":{"useSSL":"false"}}
	ConnectionParams *string `json:"ConnectionParams,omitempty" xml:"ConnectionParams,omitempty"`
	// example:
	//
	// 145
	DatasourceId *int64 `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
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
	// target
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1827811800555555
	Updater *string `json:"Updater,omitempty" xml:"Updater,omitempty"`
}

func (s ListDatasourcesResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s ListDatasourcesResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListDatasourcesResponseBodyDataRecords) GetConnectionParams() *string {
	return s.ConnectionParams
}

func (s *ListDatasourcesResponseBodyDataRecords) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *ListDatasourcesResponseBodyDataRecords) GetDescription() *string {
	return s.Description
}

func (s *ListDatasourcesResponseBodyDataRecords) GetName() *string {
	return s.Name
}

func (s *ListDatasourcesResponseBodyDataRecords) GetType() *int32 {
	return s.Type
}

func (s *ListDatasourcesResponseBodyDataRecords) GetUpdater() *string {
	return s.Updater
}

func (s *ListDatasourcesResponseBodyDataRecords) SetConnectionParams(v string) *ListDatasourcesResponseBodyDataRecords {
	s.ConnectionParams = &v
	return s
}

func (s *ListDatasourcesResponseBodyDataRecords) SetDatasourceId(v int64) *ListDatasourcesResponseBodyDataRecords {
	s.DatasourceId = &v
	return s
}

func (s *ListDatasourcesResponseBodyDataRecords) SetDescription(v string) *ListDatasourcesResponseBodyDataRecords {
	s.Description = &v
	return s
}

func (s *ListDatasourcesResponseBodyDataRecords) SetName(v string) *ListDatasourcesResponseBodyDataRecords {
	s.Name = &v
	return s
}

func (s *ListDatasourcesResponseBodyDataRecords) SetType(v int32) *ListDatasourcesResponseBodyDataRecords {
	s.Type = &v
	return s
}

func (s *ListDatasourcesResponseBodyDataRecords) SetUpdater(v string) *ListDatasourcesResponseBodyDataRecords {
	s.Updater = &v
	return s
}

func (s *ListDatasourcesResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
