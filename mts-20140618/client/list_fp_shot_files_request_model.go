// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFpShotFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListFpShotFilesRequest
	GetEndTime() *string
	SetFpDBId(v string) *ListFpShotFilesRequest
	GetFpDBId() *string
	SetNextPageToken(v string) *ListFpShotFilesRequest
	GetNextPageToken() *string
	SetOwnerAccount(v string) *ListFpShotFilesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListFpShotFilesRequest
	GetOwnerId() *int64
	SetPageSize(v int32) *ListFpShotFilesRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListFpShotFilesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListFpShotFilesRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *ListFpShotFilesRequest
	GetStartTime() *string
}

type ListFpShotFilesRequest struct {
	// The end of the time range to query. The media files to be returned must be stored before the specified end time. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// > This parameter is available only in the China (Beijing), China (Hangzhou), and China (Shanghai) regions.
	//
	// example:
	//
	// 2022-09-08T23:32:56Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the media fingerprint library whose files you want to query. You can obtain the library ID from the response parameters of the [CreateFpShotDB](https://help.aliyun.com/document_detail/170149.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2288c6ca184c0e47098a5b665e2a12****
	FpDBId *string `json:"FpDBId,omitempty" xml:"FpDBId,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// ae0fd49c0840e14daf0d66a75b83****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. The media files to be returned must be stored after the specified start time. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// > This parameter is available only in the China (Beijing), China (Hangzhou), and China (Shanghai) regions.
	//
	// example:
	//
	// 2022-09-01T00:00:28Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListFpShotFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFpShotFilesRequest) GoString() string {
	return s.String()
}

func (s *ListFpShotFilesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListFpShotFilesRequest) GetFpDBId() *string {
	return s.FpDBId
}

func (s *ListFpShotFilesRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListFpShotFilesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListFpShotFilesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListFpShotFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFpShotFilesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListFpShotFilesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListFpShotFilesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListFpShotFilesRequest) SetEndTime(v string) *ListFpShotFilesRequest {
	s.EndTime = &v
	return s
}

func (s *ListFpShotFilesRequest) SetFpDBId(v string) *ListFpShotFilesRequest {
	s.FpDBId = &v
	return s
}

func (s *ListFpShotFilesRequest) SetNextPageToken(v string) *ListFpShotFilesRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListFpShotFilesRequest) SetOwnerAccount(v string) *ListFpShotFilesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListFpShotFilesRequest) SetOwnerId(v int64) *ListFpShotFilesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListFpShotFilesRequest) SetPageSize(v int32) *ListFpShotFilesRequest {
	s.PageSize = &v
	return s
}

func (s *ListFpShotFilesRequest) SetResourceOwnerAccount(v string) *ListFpShotFilesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListFpShotFilesRequest) SetResourceOwnerId(v int64) *ListFpShotFilesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListFpShotFilesRequest) SetStartTime(v string) *ListFpShotFilesRequest {
	s.StartTime = &v
	return s
}

func (s *ListFpShotFilesRequest) Validate() error {
	return dara.Validate(s)
}
