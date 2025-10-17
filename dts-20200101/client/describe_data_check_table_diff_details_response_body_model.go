// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataCheckTableDiffDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DescribeDataCheckTableDiffDetailsResponseBody
	GetDbName() *string
	SetDiffCount(v int64) *DescribeDataCheckTableDiffDetailsResponseBody
	GetDiffCount() *int64
	SetDiffDetails(v []*DescribeDataCheckTableDiffDetailsResponseBodyDiffDetails) *DescribeDataCheckTableDiffDetailsResponseBody
	GetDiffDetails() []*DescribeDataCheckTableDiffDetailsResponseBodyDiffDetails
	SetDynamicMessage(v string) *DescribeDataCheckTableDiffDetailsResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeDataCheckTableDiffDetailsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeDataCheckTableDiffDetailsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeDataCheckTableDiffDetailsResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *DescribeDataCheckTableDiffDetailsResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *DescribeDataCheckTableDiffDetailsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDataCheckTableDiffDetailsResponseBody
	GetSuccess() *bool
	SetTbName(v string) *DescribeDataCheckTableDiffDetailsResponseBody
	GetTbName() *string
}

type DescribeDataCheckTableDiffDetailsResponseBody struct {
	// The name of the source database to which the table that contains inconsistent data belongs.
	//
	// example:
	//
	// db_dtstest
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The number of data rows that contain inconsistent data.
	//
	// example:
	//
	// 1
	DiffCount *int64 `json:"DiffCount,omitempty" xml:"DiffCount,omitempty"`
	// The information about the inconsistent data.
	DiffDetails []*DescribeDataCheckTableDiffDetailsResponseBodyDiffDetails `json:"DiffDetails,omitempty" xml:"DiffDetails,omitempty" type:"Repeated"`
	// The dynamic part in the error message. This parameter is used to replace the \\*\\*%s\\*\\	- variable in the **ErrMessage*	- parameter.
	//
	// > For example, if the value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the value of the **DynamicMessage*	- parameter is **Type**, the specified **Type*	- parameter is invalid.
	//
	// example:
	//
	// Type
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// InvalidParameter
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The value of the parameter tbName is invalid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// dtsog8q1z3tc9t****"
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 621BB4F8-3016-4FAA-8D5A-5D3163CC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The name of the table that contains inconsistent data in the source database.
	//
	// example:
	//
	// test_person
	TbName *string `json:"TbName,omitempty" xml:"TbName,omitempty"`
}

func (s DescribeDataCheckTableDiffDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataCheckTableDiffDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) GetDiffCount() *int64 {
	return s.DiffCount
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) GetDiffDetails() []*DescribeDataCheckTableDiffDetailsResponseBodyDiffDetails {
	return s.DiffDetails
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) GetTbName() *string {
	return s.TbName
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) SetDbName(v string) *DescribeDataCheckTableDiffDetailsResponseBody {
	s.DbName = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) SetDiffCount(v int64) *DescribeDataCheckTableDiffDetailsResponseBody {
	s.DiffCount = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) SetDiffDetails(v []*DescribeDataCheckTableDiffDetailsResponseBodyDiffDetails) *DescribeDataCheckTableDiffDetailsResponseBody {
	s.DiffDetails = v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) SetDynamicMessage(v string) *DescribeDataCheckTableDiffDetailsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) SetErrCode(v string) *DescribeDataCheckTableDiffDetailsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) SetErrMessage(v string) *DescribeDataCheckTableDiffDetailsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) SetHttpStatusCode(v int32) *DescribeDataCheckTableDiffDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) SetInstanceId(v string) *DescribeDataCheckTableDiffDetailsResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) SetRequestId(v string) *DescribeDataCheckTableDiffDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) SetSuccess(v bool) *DescribeDataCheckTableDiffDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) SetTbName(v string) *DescribeDataCheckTableDiffDetailsResponseBody {
	s.TbName = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsResponseBody) Validate() error {
	if s.DiffDetails != nil {
		for _, item := range s.DiffDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataCheckTableDiffDetailsResponseBodyDiffDetails struct {
	// The details of the inconsistent data, whose value is a JSON string. The JSON string contains the following parameters:
	//
	// 	- column: the name of the field.
	//
	// 	- source: the value of the field in the source database.
	//
	// 	- dest: the value of the field in the destination database.
	//
	// 	- isPrimary: indicates whether the field is a primary key.
	//
	// example:
	//
	// [     {         "column": "id",         "source": "9511",         "dest": "9511",         "isPrimary": true     },     {         "column": "state",         "source": "3",         "dest": "2",         "isPrimary": false     },     {         "column": "create_time",         "source": "2023-04-11 14:07:17.0",         "dest": "NULL",         "isPrimary": false     },     {         "column": "update_time",         "source": "2023-04-11 06:07:17.0",         "dest": "2023-04-11 06:02:29.0",         "isPrimary": false     } ]
	Diff *string `json:"Diff,omitempty" xml:"Diff,omitempty"`
	// The time when the data verification was performed.
	//
	// example:
	//
	// 2023-04-23T10:36:05.000+00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The auto-increment primary key that is used to identify the data in a verification result.
	//
	// example:
	//
	// 13058****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeDataCheckTableDiffDetailsResponseBodyDiffDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataCheckTableDiffDetailsResponseBodyDiffDetails) GoString() string {
	return s.String()
}

func (s *DescribeDataCheckTableDiffDetailsResponseBodyDiffDetails) GetDiff() *string {
	return s.Diff
}

func (s *DescribeDataCheckTableDiffDetailsResponseBodyDiffDetails) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeDataCheckTableDiffDetailsResponseBodyDiffDetails) GetId() *int64 {
	return s.Id
}

func (s *DescribeDataCheckTableDiffDetailsResponseBodyDiffDetails) SetDiff(v string) *DescribeDataCheckTableDiffDetailsResponseBodyDiffDetails {
	s.Diff = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsResponseBodyDiffDetails) SetGmtCreated(v string) *DescribeDataCheckTableDiffDetailsResponseBodyDiffDetails {
	s.GmtCreated = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsResponseBodyDiffDetails) SetId(v int64) *DescribeDataCheckTableDiffDetailsResponseBodyDiffDetails {
	s.Id = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsResponseBodyDiffDetails) Validate() error {
	return dara.Validate(s)
}
