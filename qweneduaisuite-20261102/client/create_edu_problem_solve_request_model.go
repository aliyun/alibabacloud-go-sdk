// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEduProblemSolveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQuery(v *CreateEduProblemSolveRequestQuery) *CreateEduProblemSolveRequest
	GetQuery() *CreateEduProblemSolveRequestQuery
	SetRequestId(v string) *CreateEduProblemSolveRequest
	GetRequestId() *string
}

type CreateEduProblemSolveRequest struct {
	// The query information.
	Query *CreateEduProblemSolveRequestQuery `json:"query,omitempty" xml:"query,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 019F593F-D6EE-5F44-819A-6CDA127D1637
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateEduProblemSolveRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEduProblemSolveRequest) GoString() string {
	return s.String()
}

func (s *CreateEduProblemSolveRequest) GetQuery() *CreateEduProblemSolveRequestQuery {
	return s.Query
}

func (s *CreateEduProblemSolveRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEduProblemSolveRequest) SetQuery(v *CreateEduProblemSolveRequestQuery) *CreateEduProblemSolveRequest {
	s.Query = v
	return s
}

func (s *CreateEduProblemSolveRequest) SetRequestId(v string) *CreateEduProblemSolveRequest {
	s.RequestId = &v
	return s
}

func (s *CreateEduProblemSolveRequest) Validate() error {
	if s.Query != nil {
		if err := s.Query.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEduProblemSolveRequestQuery struct {
	// The query image.
	Image *CreateEduProblemSolveRequestQueryImage `json:"image,omitempty" xml:"image,omitempty" type:"Struct"`
}

func (s CreateEduProblemSolveRequestQuery) String() string {
	return dara.Prettify(s)
}

func (s CreateEduProblemSolveRequestQuery) GoString() string {
	return s.String()
}

func (s *CreateEduProblemSolveRequestQuery) GetImage() *CreateEduProblemSolveRequestQueryImage {
	return s.Image
}

func (s *CreateEduProblemSolveRequestQuery) SetImage(v *CreateEduProblemSolveRequestQueryImage) *CreateEduProblemSolveRequestQuery {
	s.Image = v
	return s
}

func (s *CreateEduProblemSolveRequestQuery) Validate() error {
	if s.Image != nil {
		if err := s.Image.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEduProblemSolveRequestQueryImage struct {
	// The Base64 encoding of the image.
	//
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASABIAAD/4QCMRXhpZgAATU0AK...
	Base64 *string `json:"base64,omitempty" xml:"base64,omitempty"`
	// The image format. Valid values: jpg, jpeg, and png. The value must match the actual file format.
	//
	// example:
	//
	// jpg
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// The image URL.
	//
	// example:
	//
	// https://xxx/xxx.jpg
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateEduProblemSolveRequestQueryImage) String() string {
	return dara.Prettify(s)
}

func (s CreateEduProblemSolveRequestQueryImage) GoString() string {
	return s.String()
}

func (s *CreateEduProblemSolveRequestQueryImage) GetBase64() *string {
	return s.Base64
}

func (s *CreateEduProblemSolveRequestQueryImage) GetFormat() *string {
	return s.Format
}

func (s *CreateEduProblemSolveRequestQueryImage) GetUrl() *string {
	return s.Url
}

func (s *CreateEduProblemSolveRequestQueryImage) SetBase64(v string) *CreateEduProblemSolveRequestQueryImage {
	s.Base64 = &v
	return s
}

func (s *CreateEduProblemSolveRequestQueryImage) SetFormat(v string) *CreateEduProblemSolveRequestQueryImage {
	s.Format = &v
	return s
}

func (s *CreateEduProblemSolveRequestQueryImage) SetUrl(v string) *CreateEduProblemSolveRequestQueryImage {
	s.Url = &v
	return s
}

func (s *CreateEduProblemSolveRequestQueryImage) Validate() error {
	return dara.Validate(s)
}
