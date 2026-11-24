// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEduHomeworkCorrectSubmitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDewarp(v bool) *CreateEduHomeworkCorrectSubmitRequest
	GetDewarp() *bool
	SetImageBase64(v string) *CreateEduHomeworkCorrectSubmitRequest
	GetImageBase64() *string
	SetImageUrl(v string) *CreateEduHomeworkCorrectSubmitRequest
	GetImageUrl() *string
	SetRequestId(v string) *CreateEduHomeworkCorrectSubmitRequest
	GetRequestId() *string
}

type CreateEduHomeworkCorrectSubmitRequest struct {
	// Specifies whether to correct the image distortion.
	//
	// example:
	//
	// true
	Dewarp *bool `json:"dewarp,omitempty" xml:"dewarp,omitempty"`
	// The Base64-encoded image.
	//
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASABIAAD/4QCMRXhpZgAATU0AK...
	ImageBase64 *string `json:"imageBase64,omitempty" xml:"imageBase64,omitempty"`
	// The URL of the image.
	//
	// example:
	//
	// https://xxx/xxx.jpg
	ImageUrl *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019F593F-D6EE-5F44-819A-6CDA127D1637
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateEduHomeworkCorrectSubmitRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEduHomeworkCorrectSubmitRequest) GoString() string {
	return s.String()
}

func (s *CreateEduHomeworkCorrectSubmitRequest) GetDewarp() *bool {
	return s.Dewarp
}

func (s *CreateEduHomeworkCorrectSubmitRequest) GetImageBase64() *string {
	return s.ImageBase64
}

func (s *CreateEduHomeworkCorrectSubmitRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateEduHomeworkCorrectSubmitRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEduHomeworkCorrectSubmitRequest) SetDewarp(v bool) *CreateEduHomeworkCorrectSubmitRequest {
	s.Dewarp = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitRequest) SetImageBase64(v string) *CreateEduHomeworkCorrectSubmitRequest {
	s.ImageBase64 = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitRequest) SetImageUrl(v string) *CreateEduHomeworkCorrectSubmitRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitRequest) SetRequestId(v string) *CreateEduHomeworkCorrectSubmitRequest {
	s.RequestId = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitRequest) Validate() error {
	return dara.Validate(s)
}
