// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileBlobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GetFileBlobsRequest
	GetAccessToken() *string
	SetFilePath(v string) *GetFileBlobsRequest
	GetFilePath() *string
	SetFrom(v int64) *GetFileBlobsRequest
	GetFrom() *int64
	SetOrganizationId(v string) *GetFileBlobsRequest
	GetOrganizationId() *string
	SetRef(v string) *GetFileBlobsRequest
	GetRef() *string
	SetTo(v int64) *GetFileBlobsRequest
	GetTo() *int64
}

type GetFileBlobsRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// src/Test.java
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// example:
	//
	// 10
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// master  / tag1.0 /  ecykhdd
	Ref *string `json:"ref,omitempty" xml:"ref,omitempty"`
	// example:
	//
	// 10
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
}

func (s GetFileBlobsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileBlobsRequest) GoString() string {
	return s.String()
}

func (s *GetFileBlobsRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetFileBlobsRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *GetFileBlobsRequest) GetFrom() *int64 {
	return s.From
}

func (s *GetFileBlobsRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetFileBlobsRequest) GetRef() *string {
	return s.Ref
}

func (s *GetFileBlobsRequest) GetTo() *int64 {
	return s.To
}

func (s *GetFileBlobsRequest) SetAccessToken(v string) *GetFileBlobsRequest {
	s.AccessToken = &v
	return s
}

func (s *GetFileBlobsRequest) SetFilePath(v string) *GetFileBlobsRequest {
	s.FilePath = &v
	return s
}

func (s *GetFileBlobsRequest) SetFrom(v int64) *GetFileBlobsRequest {
	s.From = &v
	return s
}

func (s *GetFileBlobsRequest) SetOrganizationId(v string) *GetFileBlobsRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetFileBlobsRequest) SetRef(v string) *GetFileBlobsRequest {
	s.Ref = &v
	return s
}

func (s *GetFileBlobsRequest) SetTo(v int64) *GetFileBlobsRequest {
	s.To = &v
	return s
}

func (s *GetFileBlobsRequest) Validate() error {
	return dara.Validate(s)
}
