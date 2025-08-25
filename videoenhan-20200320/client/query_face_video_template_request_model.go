// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceVideoTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int64) *QueryFaceVideoTemplateRequest
	GetPageNo() *int64
	SetPageSize(v int64) *QueryFaceVideoTemplateRequest
	GetPageSize() *int64
	SetTemplateId(v string) *QueryFaceVideoTemplateRequest
	GetTemplateId() *string
}

type QueryFaceVideoTemplateRequest struct {
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 3bf2418c-7adf-4002-a9d6-2f7cf1889c0d
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryFaceVideoTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceVideoTemplateRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceVideoTemplateRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QueryFaceVideoTemplateRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryFaceVideoTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *QueryFaceVideoTemplateRequest) SetPageNo(v int64) *QueryFaceVideoTemplateRequest {
	s.PageNo = &v
	return s
}

func (s *QueryFaceVideoTemplateRequest) SetPageSize(v int64) *QueryFaceVideoTemplateRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceVideoTemplateRequest) SetTemplateId(v string) *QueryFaceVideoTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *QueryFaceVideoTemplateRequest) Validate() error {
	return dara.Validate(s)
}
