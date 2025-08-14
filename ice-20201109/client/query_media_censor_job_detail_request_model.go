// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaCensorJobDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *QueryMediaCensorJobDetailRequest
	GetJobId() *string
	SetMaximumPageSize(v int64) *QueryMediaCensorJobDetailRequest
	GetMaximumPageSize() *int64
	SetNextPageToken(v string) *QueryMediaCensorJobDetailRequest
	GetNextPageToken() *string
	SetOwnerAccount(v string) *QueryMediaCensorJobDetailRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QueryMediaCensorJobDetailRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryMediaCensorJobDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryMediaCensorJobDetailRequest
	GetResourceOwnerId() *int64
}

type QueryMediaCensorJobDetailRequest struct {
	// The ID of the content moderation job. You can obtain the job ID from the response parameters of the [SubmitMediaCensorJob](https://help.aliyun.com/document_detail/444848.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2288c6ca184c0e47098a5b665e2a12****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The number of entries per page.
	//
	// 	- Default value: **30**.
	//
	// 	- Valid values: **1 to 300**.
	//
	// example:
	//
	// 30
	MaximumPageSize *int64 `json:"MaximumPageSize,omitempty" xml:"MaximumPageSize,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// ae0fd49c0840e14daf0d66a75b83****
	NextPageToken        *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryMediaCensorJobDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailRequest) GetJobId() *string {
	return s.JobId
}

func (s *QueryMediaCensorJobDetailRequest) GetMaximumPageSize() *int64 {
	return s.MaximumPageSize
}

func (s *QueryMediaCensorJobDetailRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *QueryMediaCensorJobDetailRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QueryMediaCensorJobDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryMediaCensorJobDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryMediaCensorJobDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryMediaCensorJobDetailRequest) SetJobId(v string) *QueryMediaCensorJobDetailRequest {
	s.JobId = &v
	return s
}

func (s *QueryMediaCensorJobDetailRequest) SetMaximumPageSize(v int64) *QueryMediaCensorJobDetailRequest {
	s.MaximumPageSize = &v
	return s
}

func (s *QueryMediaCensorJobDetailRequest) SetNextPageToken(v string) *QueryMediaCensorJobDetailRequest {
	s.NextPageToken = &v
	return s
}

func (s *QueryMediaCensorJobDetailRequest) SetOwnerAccount(v string) *QueryMediaCensorJobDetailRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QueryMediaCensorJobDetailRequest) SetOwnerId(v int64) *QueryMediaCensorJobDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryMediaCensorJobDetailRequest) SetResourceOwnerAccount(v string) *QueryMediaCensorJobDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryMediaCensorJobDetailRequest) SetResourceOwnerId(v int64) *QueryMediaCensorJobDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryMediaCensorJobDetailRequest) Validate() error {
	return dara.Validate(s)
}
