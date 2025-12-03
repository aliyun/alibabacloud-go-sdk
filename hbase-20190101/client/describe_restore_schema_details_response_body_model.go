// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreSchemaDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRestoreSchemaDetailsResponseBody
	GetRequestId() *string
	SetRestoreSchema(v *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) *DescribeRestoreSchemaDetailsResponseBody
	GetRestoreSchema() *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema
}

type DescribeRestoreSchemaDetailsResponseBody struct {
	// example:
	//
	// BC682A80-7677-4294-975C-CFEA425381DE
	RequestId     *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RestoreSchema *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema `json:"RestoreSchema,omitempty" xml:"RestoreSchema,omitempty" type:"Struct"`
}

func (s DescribeRestoreSchemaDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreSchemaDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSchemaDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRestoreSchemaDetailsResponseBody) GetRestoreSchema() *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema {
	return s.RestoreSchema
}

func (s *DescribeRestoreSchemaDetailsResponseBody) SetRequestId(v string) *DescribeRestoreSchemaDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBody) SetRestoreSchema(v *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) *DescribeRestoreSchemaDetailsResponseBody {
	s.RestoreSchema = v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBody) Validate() error {
	if s.RestoreSchema != nil {
		if err := s.RestoreSchema.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRestoreSchemaDetailsResponseBodyRestoreSchema struct {
	// example:
	//
	// 0
	Fail *int32 `json:"Fail,omitempty" xml:"Fail,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize             *int32                                                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RestoreSchemaDetails *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails `json:"RestoreSchemaDetails,omitempty" xml:"RestoreSchemaDetails,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Succeed *int32 `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) GetFail() *int32 {
	return s.Fail
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) GetRestoreSchemaDetails() *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails {
	return s.RestoreSchemaDetails
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) GetSucceed() *int32 {
	return s.Succeed
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) SetFail(v int32) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema {
	s.Fail = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) SetPageNumber(v int32) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) SetPageSize(v int32) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) SetRestoreSchemaDetails(v *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema {
	s.RestoreSchemaDetails = v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) SetSucceed(v int32) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema {
	s.Succeed = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) SetTotal(v int64) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema {
	s.Total = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) Validate() error {
	if s.RestoreSchemaDetails != nil {
		if err := s.RestoreSchemaDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails struct {
	RestoreSchemaDetail []*DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail `json:"RestoreSchemaDetail,omitempty" xml:"RestoreSchemaDetail,omitempty" type:"Repeated"`
}

func (s DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails) GetRestoreSchemaDetail() []*DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	return s.RestoreSchemaDetail
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails) SetRestoreSchemaDetail(v []*DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails {
	s.RestoreSchemaDetail = v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails) Validate() error {
	if s.RestoreSchemaDetail != nil {
		for _, item := range s.RestoreSchemaDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail struct {
	// example:
	//
	// 2020-11-05T06:45:18Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:14Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCEEDED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// default:test1
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) GetMessage() *string {
	return s.Message
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) GetState() *string {
	return s.State
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) GetTable() *string {
	return s.Table
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetEndTime(v string) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetMessage(v string) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.Message = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetStartTime(v string) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetState(v string) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.State = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetTable(v string) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.Table = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) Validate() error {
	return dara.Validate(s)
}
