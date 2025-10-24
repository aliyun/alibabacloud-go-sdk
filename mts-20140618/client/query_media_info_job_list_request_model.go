// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaInfoJobListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaInfoJobIds(v string) *QueryMediaInfoJobListRequest
	GetMediaInfoJobIds() *string
	SetOwnerAccount(v string) *QueryMediaInfoJobListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QueryMediaInfoJobListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryMediaInfoJobListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryMediaInfoJobListRequest
	GetResourceOwnerId() *int64
}

type QueryMediaInfoJobListRequest struct {
	// The IDs of the media information analysis jobs.
	//
	// 	- You can query up to 10 jobs at a time. Separate multiple IDs with commas (,).
	//
	// 	- You can obtain the details from the response parameters of the [SubmitMediaInfoJob](https://help.aliyun.com/document_detail/602827.html) operation.
	//
	// >  If you do not specify the JobIds parameter, the **InvalidParameter*	- error code is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23ca1d184c0e4341e5b665e2a12****
	MediaInfoJobIds      *string `json:"MediaInfoJobIds,omitempty" xml:"MediaInfoJobIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryMediaInfoJobListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaInfoJobListRequest) GoString() string {
	return s.String()
}

func (s *QueryMediaInfoJobListRequest) GetMediaInfoJobIds() *string {
	return s.MediaInfoJobIds
}

func (s *QueryMediaInfoJobListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QueryMediaInfoJobListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryMediaInfoJobListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryMediaInfoJobListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryMediaInfoJobListRequest) SetMediaInfoJobIds(v string) *QueryMediaInfoJobListRequest {
	s.MediaInfoJobIds = &v
	return s
}

func (s *QueryMediaInfoJobListRequest) SetOwnerAccount(v string) *QueryMediaInfoJobListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QueryMediaInfoJobListRequest) SetOwnerId(v int64) *QueryMediaInfoJobListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryMediaInfoJobListRequest) SetResourceOwnerAccount(v string) *QueryMediaInfoJobListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryMediaInfoJobListRequest) SetResourceOwnerId(v int64) *QueryMediaInfoJobListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryMediaInfoJobListRequest) Validate() error {
	return dara.Validate(s)
}
