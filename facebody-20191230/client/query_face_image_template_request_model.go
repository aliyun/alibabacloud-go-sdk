// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceImageTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int64) *QueryFaceImageTemplateRequest
	GetPageNo() *int64
	SetPageSize(v int64) *QueryFaceImageTemplateRequest
	GetPageSize() *int64
	SetTemplateId(v string) *QueryFaceImageTemplateRequest
	GetTemplateId() *string
}

type QueryFaceImageTemplateRequest struct {
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 6cd509ea-54fa-4730-8e9d-c94cadcda048
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryFaceImageTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceImageTemplateRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceImageTemplateRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QueryFaceImageTemplateRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryFaceImageTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *QueryFaceImageTemplateRequest) SetPageNo(v int64) *QueryFaceImageTemplateRequest {
	s.PageNo = &v
	return s
}

func (s *QueryFaceImageTemplateRequest) SetPageSize(v int64) *QueryFaceImageTemplateRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceImageTemplateRequest) SetTemplateId(v string) *QueryFaceImageTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *QueryFaceImageTemplateRequest) Validate() error {
	return dara.Validate(s)
}
