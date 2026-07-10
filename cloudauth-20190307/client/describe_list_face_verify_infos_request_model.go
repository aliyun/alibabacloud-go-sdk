// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListFaceVerifyInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertifyId(v string) *DescribeListFaceVerifyInfosRequest
	GetCertifyId() *string
	SetGmtEnd(v int64) *DescribeListFaceVerifyInfosRequest
	GetGmtEnd() *int64
	SetGmtStart(v int64) *DescribeListFaceVerifyInfosRequest
	GetGmtStart() *int64
	SetPageNumber(v int32) *DescribeListFaceVerifyInfosRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeListFaceVerifyInfosRequest
	GetPageSize() *int32
	SetSceneId(v int64) *DescribeListFaceVerifyInfosRequest
	GetSceneId() *int64
	SetStatus(v int32) *DescribeListFaceVerifyInfosRequest
	GetStatus() *int32
}

type DescribeListFaceVerifyInfosRequest struct {
	// The authentication ID.
	//
	// example:
	//
	// shs414a8b392a3a338abe0504c75c056
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// The end time of the authentication query.
	//
	// example:
	//
	// 1760716800000
	GmtEnd *int64 `json:"GmtEnd,omitempty" xml:"GmtEnd,omitempty"`
	// The start time of the authentication query.
	//
	// example:
	//
	// 1760112000000
	GmtStart *int64 `json:"GmtStart,omitempty" xml:"GmtStart,omitempty"`
	// The current page number for pagination.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The scene ID.
	//
	// example:
	//
	// 1000009699
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The authentication status. Valid values:
	//
	// - **1**: Authentication passed.
	//
	// - **2**: Authentication failed.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeListFaceVerifyInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeListFaceVerifyInfosRequest) GoString() string {
	return s.String()
}

func (s *DescribeListFaceVerifyInfosRequest) GetCertifyId() *string {
	return s.CertifyId
}

func (s *DescribeListFaceVerifyInfosRequest) GetGmtEnd() *int64 {
	return s.GmtEnd
}

func (s *DescribeListFaceVerifyInfosRequest) GetGmtStart() *int64 {
	return s.GmtStart
}

func (s *DescribeListFaceVerifyInfosRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeListFaceVerifyInfosRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeListFaceVerifyInfosRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DescribeListFaceVerifyInfosRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeListFaceVerifyInfosRequest) SetCertifyId(v string) *DescribeListFaceVerifyInfosRequest {
	s.CertifyId = &v
	return s
}

func (s *DescribeListFaceVerifyInfosRequest) SetGmtEnd(v int64) *DescribeListFaceVerifyInfosRequest {
	s.GmtEnd = &v
	return s
}

func (s *DescribeListFaceVerifyInfosRequest) SetGmtStart(v int64) *DescribeListFaceVerifyInfosRequest {
	s.GmtStart = &v
	return s
}

func (s *DescribeListFaceVerifyInfosRequest) SetPageNumber(v int32) *DescribeListFaceVerifyInfosRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeListFaceVerifyInfosRequest) SetPageSize(v int32) *DescribeListFaceVerifyInfosRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeListFaceVerifyInfosRequest) SetSceneId(v int64) *DescribeListFaceVerifyInfosRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeListFaceVerifyInfosRequest) SetStatus(v int32) *DescribeListFaceVerifyInfosRequest {
	s.Status = &v
	return s
}

func (s *DescribeListFaceVerifyInfosRequest) Validate() error {
	return dara.Validate(s)
}
