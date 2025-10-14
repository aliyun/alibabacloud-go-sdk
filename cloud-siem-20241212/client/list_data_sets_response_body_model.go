// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataSets(v []*ListDataSetsResponseBodyDataSets) *ListDataSetsResponseBody
	GetDataSets() []*ListDataSetsResponseBodyDataSets
	SetMaxResults(v int32) *ListDataSetsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataSetsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListDataSetsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataSetsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDataSetsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDataSetsResponseBody
	GetTotalCount() *int32
}

type ListDataSetsResponseBody struct {
	DataSets []*ListDataSetsResponseBodyDataSets `json:"DataSets,omitempty" xml:"DataSets,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 157CFBB5-B56F-566F-8991-B3C51799****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 57
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataSetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSetsResponseBody) GetDataSets() []*ListDataSetsResponseBodyDataSets {
	return s.DataSets
}

func (s *ListDataSetsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataSetsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataSetsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataSetsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataSetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSetsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataSetsResponseBody) SetDataSets(v []*ListDataSetsResponseBodyDataSets) *ListDataSetsResponseBody {
	s.DataSets = v
	return s
}

func (s *ListDataSetsResponseBody) SetMaxResults(v int32) *ListDataSetsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataSetsResponseBody) SetNextToken(v string) *ListDataSetsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataSetsResponseBody) SetPageNumber(v int32) *ListDataSetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDataSetsResponseBody) SetPageSize(v int32) *ListDataSetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataSetsResponseBody) SetRequestId(v string) *ListDataSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSetsResponseBody) SetTotalCount(v int32) *ListDataSetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataSetsResponseBody) Validate() error {
	if s.DataSets != nil {
		for _, item := range s.DataSets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataSetsResponseBodyDataSets struct {
	// example:
	//
	// 1713787368000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// lmftest desc
	DataSetDescription *string `json:"DataSetDescription,omitempty" xml:"DataSetDescription,omitempty"`
	// example:
	//
	// ip
	DataSetFieldKeyName *string `json:"DataSetFieldKeyName,omitempty" xml:"DataSetFieldKeyName,omitempty"`
	// example:
	//
	// ["ip","region"]
	DataSetFieldNames *string `json:"DataSetFieldNames,omitempty" xml:"DataSetFieldNames,omitempty"`
	// example:
	//
	// cloudsiem-dataset/1358117679873357_174338773****.csv
	DataSetFileName *string `json:"DataSetFileName,omitempty" xml:"DataSetFileName,omitempty"`
	// example:
	//
	// dataset-t8ha6p7k61rmniqw****
	DataSetId *string `json:"DataSetId,omitempty" xml:"DataSetId,omitempty"`
	// example:
	//
	// lmftest
	DataSetName       *string                                              `json:"DataSetName,omitempty" xml:"DataSetName,omitempty"`
	DataSetReferences []*ListDataSetsResponseBodyDataSetsDataSetReferences `json:"DataSetReferences,omitempty" xml:"DataSetReferences,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	DataSetStatus *int32 `json:"DataSetStatus,omitempty" xml:"DataSetStatus,omitempty"`
	// example:
	//
	// preset
	DataSetType            *string                                                   `json:"DataSetType,omitempty" xml:"DataSetType,omitempty"`
	IpWhitelistRecognizers []*ListDataSetsResponseBodyDataSetsIpWhitelistRecognizers `json:"IpWhitelistRecognizers,omitempty" xml:"IpWhitelistRecognizers,omitempty" type:"Repeated"`
	// example:
	//
	// 1713787368000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListDataSetsResponseBodyDataSets) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetsResponseBodyDataSets) GoString() string {
	return s.String()
}

func (s *ListDataSetsResponseBodyDataSets) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataSetsResponseBodyDataSets) GetDataSetDescription() *string {
	return s.DataSetDescription
}

func (s *ListDataSetsResponseBodyDataSets) GetDataSetFieldKeyName() *string {
	return s.DataSetFieldKeyName
}

func (s *ListDataSetsResponseBodyDataSets) GetDataSetFieldNames() *string {
	return s.DataSetFieldNames
}

func (s *ListDataSetsResponseBodyDataSets) GetDataSetFileName() *string {
	return s.DataSetFileName
}

func (s *ListDataSetsResponseBodyDataSets) GetDataSetId() *string {
	return s.DataSetId
}

func (s *ListDataSetsResponseBodyDataSets) GetDataSetName() *string {
	return s.DataSetName
}

func (s *ListDataSetsResponseBodyDataSets) GetDataSetReferences() []*ListDataSetsResponseBodyDataSetsDataSetReferences {
	return s.DataSetReferences
}

func (s *ListDataSetsResponseBodyDataSets) GetDataSetStatus() *int32 {
	return s.DataSetStatus
}

func (s *ListDataSetsResponseBodyDataSets) GetDataSetType() *string {
	return s.DataSetType
}

func (s *ListDataSetsResponseBodyDataSets) GetIpWhitelistRecognizers() []*ListDataSetsResponseBodyDataSetsIpWhitelistRecognizers {
	return s.IpWhitelistRecognizers
}

func (s *ListDataSetsResponseBodyDataSets) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListDataSetsResponseBodyDataSets) SetCreateTime(v int64) *ListDataSetsResponseBodyDataSets {
	s.CreateTime = &v
	return s
}

func (s *ListDataSetsResponseBodyDataSets) SetDataSetDescription(v string) *ListDataSetsResponseBodyDataSets {
	s.DataSetDescription = &v
	return s
}

func (s *ListDataSetsResponseBodyDataSets) SetDataSetFieldKeyName(v string) *ListDataSetsResponseBodyDataSets {
	s.DataSetFieldKeyName = &v
	return s
}

func (s *ListDataSetsResponseBodyDataSets) SetDataSetFieldNames(v string) *ListDataSetsResponseBodyDataSets {
	s.DataSetFieldNames = &v
	return s
}

func (s *ListDataSetsResponseBodyDataSets) SetDataSetFileName(v string) *ListDataSetsResponseBodyDataSets {
	s.DataSetFileName = &v
	return s
}

func (s *ListDataSetsResponseBodyDataSets) SetDataSetId(v string) *ListDataSetsResponseBodyDataSets {
	s.DataSetId = &v
	return s
}

func (s *ListDataSetsResponseBodyDataSets) SetDataSetName(v string) *ListDataSetsResponseBodyDataSets {
	s.DataSetName = &v
	return s
}

func (s *ListDataSetsResponseBodyDataSets) SetDataSetReferences(v []*ListDataSetsResponseBodyDataSetsDataSetReferences) *ListDataSetsResponseBodyDataSets {
	s.DataSetReferences = v
	return s
}

func (s *ListDataSetsResponseBodyDataSets) SetDataSetStatus(v int32) *ListDataSetsResponseBodyDataSets {
	s.DataSetStatus = &v
	return s
}

func (s *ListDataSetsResponseBodyDataSets) SetDataSetType(v string) *ListDataSetsResponseBodyDataSets {
	s.DataSetType = &v
	return s
}

func (s *ListDataSetsResponseBodyDataSets) SetIpWhitelistRecognizers(v []*ListDataSetsResponseBodyDataSetsIpWhitelistRecognizers) *ListDataSetsResponseBodyDataSets {
	s.IpWhitelistRecognizers = v
	return s
}

func (s *ListDataSetsResponseBodyDataSets) SetUpdateTime(v int64) *ListDataSetsResponseBodyDataSets {
	s.UpdateTime = &v
	return s
}

func (s *ListDataSetsResponseBodyDataSets) Validate() error {
	if s.DataSetReferences != nil {
		for _, item := range s.DataSetReferences {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IpWhitelistRecognizers != nil {
		for _, item := range s.IpWhitelistRecognizers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataSetsResponseBodyDataSetsDataSetReferences struct {
	// example:
	//
	// dataset-nhcrmjpx1zsorlaq****
	DataSetId *string `json:"DataSetId,omitempty" xml:"DataSetId,omitempty"`
	// example:
	//
	// 456232
	DataSetReferenceId *string `json:"DataSetReferenceId,omitempty" xml:"DataSetReferenceId,omitempty"`
	// example:
	//
	// playbook_name
	DataSetReferenceName *string `json:"DataSetReferenceName,omitempty" xml:"DataSetReferenceName,omitempty"`
	// example:
	//
	// playbook
	DataSetReferenceType *string `json:"DataSetReferenceType,omitempty" xml:"DataSetReferenceType,omitempty"`
}

func (s ListDataSetsResponseBodyDataSetsDataSetReferences) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetsResponseBodyDataSetsDataSetReferences) GoString() string {
	return s.String()
}

func (s *ListDataSetsResponseBodyDataSetsDataSetReferences) GetDataSetId() *string {
	return s.DataSetId
}

func (s *ListDataSetsResponseBodyDataSetsDataSetReferences) GetDataSetReferenceId() *string {
	return s.DataSetReferenceId
}

func (s *ListDataSetsResponseBodyDataSetsDataSetReferences) GetDataSetReferenceName() *string {
	return s.DataSetReferenceName
}

func (s *ListDataSetsResponseBodyDataSetsDataSetReferences) GetDataSetReferenceType() *string {
	return s.DataSetReferenceType
}

func (s *ListDataSetsResponseBodyDataSetsDataSetReferences) SetDataSetId(v string) *ListDataSetsResponseBodyDataSetsDataSetReferences {
	s.DataSetId = &v
	return s
}

func (s *ListDataSetsResponseBodyDataSetsDataSetReferences) SetDataSetReferenceId(v string) *ListDataSetsResponseBodyDataSetsDataSetReferences {
	s.DataSetReferenceId = &v
	return s
}

func (s *ListDataSetsResponseBodyDataSetsDataSetReferences) SetDataSetReferenceName(v string) *ListDataSetsResponseBodyDataSetsDataSetReferences {
	s.DataSetReferenceName = &v
	return s
}

func (s *ListDataSetsResponseBodyDataSetsDataSetReferences) SetDataSetReferenceType(v string) *ListDataSetsResponseBodyDataSetsDataSetReferences {
	s.DataSetReferenceType = &v
	return s
}

func (s *ListDataSetsResponseBodyDataSetsDataSetReferences) Validate() error {
	return dara.Validate(s)
}

type ListDataSetsResponseBodyDataSetsIpWhitelistRecognizers struct {
	// example:
	//
	// enabled
	AutoRecognizeStatus *string `json:"AutoRecognizeStatus,omitempty" xml:"AutoRecognizeStatus,omitempty"`
	// example:
	//
	// waf_back_source_ip
	IpWhitelistRecognizerType *string `json:"IpWhitelistRecognizerType,omitempty" xml:"IpWhitelistRecognizerType,omitempty"`
	// example:
	//
	// current_account
	RecognizeScope *string `json:"RecognizeScope,omitempty" xml:"RecognizeScope,omitempty"`
	// example:
	//
	// 1713787368000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListDataSetsResponseBodyDataSetsIpWhitelistRecognizers) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetsResponseBodyDataSetsIpWhitelistRecognizers) GoString() string {
	return s.String()
}

func (s *ListDataSetsResponseBodyDataSetsIpWhitelistRecognizers) GetAutoRecognizeStatus() *string {
	return s.AutoRecognizeStatus
}

func (s *ListDataSetsResponseBodyDataSetsIpWhitelistRecognizers) GetIpWhitelistRecognizerType() *string {
	return s.IpWhitelistRecognizerType
}

func (s *ListDataSetsResponseBodyDataSetsIpWhitelistRecognizers) GetRecognizeScope() *string {
	return s.RecognizeScope
}

func (s *ListDataSetsResponseBodyDataSetsIpWhitelistRecognizers) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListDataSetsResponseBodyDataSetsIpWhitelistRecognizers) SetAutoRecognizeStatus(v string) *ListDataSetsResponseBodyDataSetsIpWhitelistRecognizers {
	s.AutoRecognizeStatus = &v
	return s
}

func (s *ListDataSetsResponseBodyDataSetsIpWhitelistRecognizers) SetIpWhitelistRecognizerType(v string) *ListDataSetsResponseBodyDataSetsIpWhitelistRecognizers {
	s.IpWhitelistRecognizerType = &v
	return s
}

func (s *ListDataSetsResponseBodyDataSetsIpWhitelistRecognizers) SetRecognizeScope(v string) *ListDataSetsResponseBodyDataSetsIpWhitelistRecognizers {
	s.RecognizeScope = &v
	return s
}

func (s *ListDataSetsResponseBodyDataSetsIpWhitelistRecognizers) SetUpdateTime(v int64) *ListDataSetsResponseBodyDataSetsIpWhitelistRecognizers {
	s.UpdateTime = &v
	return s
}

func (s *ListDataSetsResponseBodyDataSetsIpWhitelistRecognizers) Validate() error {
	return dara.Validate(s)
}
