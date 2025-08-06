// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaExportJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMediaExportJobsRequest
	GetAppId() *string
	SetEndTime(v string) *ListMediaExportJobsRequest
	GetEndTime() *string
	SetMediaType(v string) *ListMediaExportJobsRequest
	GetMediaType() *string
	SetOwnerId(v int64) *ListMediaExportJobsRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *ListMediaExportJobsRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListMediaExportJobsRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *ListMediaExportJobsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListMediaExportJobsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *ListMediaExportJobsRequest
	GetStartTime() *string
	SetStatus(v string) *ListMediaExportJobsRequest
	GetStatus() *string
}

type ListMediaExportJobsRequest struct {
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MediaType            *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMediaExportJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMediaExportJobsRequest) GoString() string {
	return s.String()
}

func (s *ListMediaExportJobsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMediaExportJobsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListMediaExportJobsRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *ListMediaExportJobsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListMediaExportJobsRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListMediaExportJobsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListMediaExportJobsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListMediaExportJobsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListMediaExportJobsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListMediaExportJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListMediaExportJobsRequest) SetAppId(v string) *ListMediaExportJobsRequest {
	s.AppId = &v
	return s
}

func (s *ListMediaExportJobsRequest) SetEndTime(v string) *ListMediaExportJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListMediaExportJobsRequest) SetMediaType(v string) *ListMediaExportJobsRequest {
	s.MediaType = &v
	return s
}

func (s *ListMediaExportJobsRequest) SetOwnerId(v int64) *ListMediaExportJobsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListMediaExportJobsRequest) SetPageNo(v int64) *ListMediaExportJobsRequest {
	s.PageNo = &v
	return s
}

func (s *ListMediaExportJobsRequest) SetPageSize(v int64) *ListMediaExportJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMediaExportJobsRequest) SetResourceOwnerAccount(v string) *ListMediaExportJobsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListMediaExportJobsRequest) SetResourceOwnerId(v int64) *ListMediaExportJobsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListMediaExportJobsRequest) SetStartTime(v string) *ListMediaExportJobsRequest {
	s.StartTime = &v
	return s
}

func (s *ListMediaExportJobsRequest) SetStatus(v string) *ListMediaExportJobsRequest {
	s.Status = &v
	return s
}

func (s *ListMediaExportJobsRequest) Validate() error {
	return dara.Validate(s)
}
