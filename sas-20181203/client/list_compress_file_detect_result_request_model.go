// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCompressFileDetectResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListCompressFileDetectResultRequest
	GetCurrentPage() *int32
	SetHashKey(v string) *ListCompressFileDetectResultRequest
	GetHashKey() *string
	SetPageSize(v int32) *ListCompressFileDetectResultRequest
	GetPageSize() *int32
	SetSourceIp(v string) *ListCompressFileDetectResultRequest
	GetSourceIp() *string
}

type ListCompressFileDetectResultRequest struct {
	// The page number. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The identifier of the file. Only MD5 hash values are supported.
	//
	// example:
	//
	// 0a212417e65c26ff133cfff28f6c****
	HashKey *string `json:"HashKey,omitempty" xml:"HashKey,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 27.9.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ListCompressFileDetectResultRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCompressFileDetectResultRequest) GoString() string {
	return s.String()
}

func (s *ListCompressFileDetectResultRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCompressFileDetectResultRequest) GetHashKey() *string {
	return s.HashKey
}

func (s *ListCompressFileDetectResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCompressFileDetectResultRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ListCompressFileDetectResultRequest) SetCurrentPage(v int32) *ListCompressFileDetectResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCompressFileDetectResultRequest) SetHashKey(v string) *ListCompressFileDetectResultRequest {
	s.HashKey = &v
	return s
}

func (s *ListCompressFileDetectResultRequest) SetPageSize(v int32) *ListCompressFileDetectResultRequest {
	s.PageSize = &v
	return s
}

func (s *ListCompressFileDetectResultRequest) SetSourceIp(v string) *ListCompressFileDetectResultRequest {
	s.SourceIp = &v
	return s
}

func (s *ListCompressFileDetectResultRequest) Validate() error {
	return dara.Validate(s)
}
