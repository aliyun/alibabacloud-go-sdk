// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *QueryFaceRecordRequest
	GetCurrentPage() *int64
	SetFaceGroupCode(v string) *QueryFaceRecordRequest
	GetFaceGroupCode() *string
	SetFaceId(v string) *QueryFaceRecordRequest
	GetFaceId() *string
	SetMaxResults(v int32) *QueryFaceRecordRequest
	GetMaxResults() *int32
	SetMerchantUserId(v string) *QueryFaceRecordRequest
	GetMerchantUserId() *string
	SetNextToken(v string) *QueryFaceRecordRequest
	GetNextToken() *string
	SetPageSize(v int32) *QueryFaceRecordRequest
	GetPageSize() *int32
	SetRegistrationType(v string) *QueryFaceRecordRequest
	GetRegistrationType() *string
}

type QueryFaceRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Cu****D08q
	FaceGroupCode *string `json:"FaceGroupCode,omitempty" xml:"FaceGroupCode,omitempty"`
	// example:
	//
	// 5006538
	FaceId *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// CuN6hiD08qr
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// example:
	//
	// AAAAARbaCuN6hiD08qrLdwJ9Fh0OP1yH8z+7FV4KKGUw4X32
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// MANUAL
	RegistrationType *string `json:"RegistrationType,omitempty" xml:"RegistrationType,omitempty"`
}

func (s QueryFaceRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceRecordRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceRecordRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *QueryFaceRecordRequest) GetFaceGroupCode() *string {
	return s.FaceGroupCode
}

func (s *QueryFaceRecordRequest) GetFaceId() *string {
	return s.FaceId
}

func (s *QueryFaceRecordRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryFaceRecordRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *QueryFaceRecordRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryFaceRecordRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryFaceRecordRequest) GetRegistrationType() *string {
	return s.RegistrationType
}

func (s *QueryFaceRecordRequest) SetCurrentPage(v int64) *QueryFaceRecordRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryFaceRecordRequest) SetFaceGroupCode(v string) *QueryFaceRecordRequest {
	s.FaceGroupCode = &v
	return s
}

func (s *QueryFaceRecordRequest) SetFaceId(v string) *QueryFaceRecordRequest {
	s.FaceId = &v
	return s
}

func (s *QueryFaceRecordRequest) SetMaxResults(v int32) *QueryFaceRecordRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryFaceRecordRequest) SetMerchantUserId(v string) *QueryFaceRecordRequest {
	s.MerchantUserId = &v
	return s
}

func (s *QueryFaceRecordRequest) SetNextToken(v string) *QueryFaceRecordRequest {
	s.NextToken = &v
	return s
}

func (s *QueryFaceRecordRequest) SetPageSize(v int32) *QueryFaceRecordRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceRecordRequest) SetRegistrationType(v string) *QueryFaceRecordRequest {
	s.RegistrationType = &v
	return s
}

func (s *QueryFaceRecordRequest) Validate() error {
	return dara.Validate(s)
}
