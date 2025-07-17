// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataTrackOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateDataTrackOrderRequest
	GetComment() *string
	SetParam(v *CreateDataTrackOrderRequestParam) *CreateDataTrackOrderRequest
	GetParam() *CreateDataTrackOrderRequestParam
	SetRelatedUserList(v []*string) *CreateDataTrackOrderRequest
	GetRelatedUserList() []*string
	SetTid(v int64) *CreateDataTrackOrderRequest
	GetTid() *int64
}

type CreateDataTrackOrderRequest struct {
	// The purpose or objective of the data tracking ticket. This parameter is used to help reduce unnecessary communication.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The parameters of the ticket.
	//
	// This parameter is required.
	Param *CreateDataTrackOrderRequestParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	// The IDs of the operators that are related to the ticket.
	RelatedUserList []*string `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Repeated"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateDataTrackOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataTrackOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateDataTrackOrderRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDataTrackOrderRequest) GetParam() *CreateDataTrackOrderRequestParam {
	return s.Param
}

func (s *CreateDataTrackOrderRequest) GetRelatedUserList() []*string {
	return s.RelatedUserList
}

func (s *CreateDataTrackOrderRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataTrackOrderRequest) SetComment(v string) *CreateDataTrackOrderRequest {
	s.Comment = &v
	return s
}

func (s *CreateDataTrackOrderRequest) SetParam(v *CreateDataTrackOrderRequestParam) *CreateDataTrackOrderRequest {
	s.Param = v
	return s
}

func (s *CreateDataTrackOrderRequest) SetRelatedUserList(v []*string) *CreateDataTrackOrderRequest {
	s.RelatedUserList = v
	return s
}

func (s *CreateDataTrackOrderRequest) SetTid(v int64) *CreateDataTrackOrderRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataTrackOrderRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDataTrackOrderRequestParam struct {
	// The ID of the database. You can call the [SearchDatabases](https://help.aliyun.com/document_detail/141876.html) operation to query the ID of the database.
	//
	// > You can call this operation to create a data tracking ticket for only physical databases. This operation is not applicable to logical databases.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123***
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The end time of the time range in which you want to track data operations. The time must be in the yyyy-MM-dd HH:mm:ss format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-23 10:00:00
	JobEndTime *string `json:"JobEndTime,omitempty" xml:"JobEndTime,omitempty"`
	// The start time of the time range in which you want to track data operations. The time must be in the yyyy-MM-dd HH:mm:ss format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-23 00:00:00
	JobStartTime *string `json:"JobStartTime,omitempty" xml:"JobStartTime,omitempty"`
	// The names of the tables for which you want to track data operations.
	//
	// This parameter is required.
	TableNames []*string `json:"TableNames,omitempty" xml:"TableNames,omitempty" type:"Repeated"`
	// The types of data operations that you want to track.
	//
	// This parameter is required.
	TrackTypes []*string `json:"TrackTypes,omitempty" xml:"TrackTypes,omitempty" type:"Repeated"`
}

func (s CreateDataTrackOrderRequestParam) String() string {
	return dara.Prettify(s)
}

func (s CreateDataTrackOrderRequestParam) GoString() string {
	return s.String()
}

func (s *CreateDataTrackOrderRequestParam) GetDbId() *string {
	return s.DbId
}

func (s *CreateDataTrackOrderRequestParam) GetJobEndTime() *string {
	return s.JobEndTime
}

func (s *CreateDataTrackOrderRequestParam) GetJobStartTime() *string {
	return s.JobStartTime
}

func (s *CreateDataTrackOrderRequestParam) GetTableNames() []*string {
	return s.TableNames
}

func (s *CreateDataTrackOrderRequestParam) GetTrackTypes() []*string {
	return s.TrackTypes
}

func (s *CreateDataTrackOrderRequestParam) SetDbId(v string) *CreateDataTrackOrderRequestParam {
	s.DbId = &v
	return s
}

func (s *CreateDataTrackOrderRequestParam) SetJobEndTime(v string) *CreateDataTrackOrderRequestParam {
	s.JobEndTime = &v
	return s
}

func (s *CreateDataTrackOrderRequestParam) SetJobStartTime(v string) *CreateDataTrackOrderRequestParam {
	s.JobStartTime = &v
	return s
}

func (s *CreateDataTrackOrderRequestParam) SetTableNames(v []*string) *CreateDataTrackOrderRequestParam {
	s.TableNames = v
	return s
}

func (s *CreateDataTrackOrderRequestParam) SetTrackTypes(v []*string) *CreateDataTrackOrderRequestParam {
	s.TrackTypes = v
	return s
}

func (s *CreateDataTrackOrderRequestParam) Validate() error {
	return dara.Validate(s)
}
