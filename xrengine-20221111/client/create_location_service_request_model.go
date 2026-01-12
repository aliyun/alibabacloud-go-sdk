// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLocationServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateLocationServiceRequest
	GetAppId() *string
	SetImageId(v int64) *CreateLocationServiceRequest
	GetImageId() *int64
	SetJwtToken(v string) *CreateLocationServiceRequest
	GetJwtToken() *string
	SetName(v string) *CreateLocationServiceRequest
	GetName() *string
	SetNote(v string) *CreateLocationServiceRequest
	GetNote() *string
	SetQps(v int64) *CreateLocationServiceRequest
	GetQps() *int64
}

type CreateLocationServiceRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	ImageId  *int64  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	Qps  *int64  `json:"Qps,omitempty" xml:"Qps,omitempty"`
}

func (s CreateLocationServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLocationServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateLocationServiceRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateLocationServiceRequest) GetImageId() *int64 {
	return s.ImageId
}

func (s *CreateLocationServiceRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *CreateLocationServiceRequest) GetName() *string {
	return s.Name
}

func (s *CreateLocationServiceRequest) GetNote() *string {
	return s.Note
}

func (s *CreateLocationServiceRequest) GetQps() *int64 {
	return s.Qps
}

func (s *CreateLocationServiceRequest) SetAppId(v string) *CreateLocationServiceRequest {
	s.AppId = &v
	return s
}

func (s *CreateLocationServiceRequest) SetImageId(v int64) *CreateLocationServiceRequest {
	s.ImageId = &v
	return s
}

func (s *CreateLocationServiceRequest) SetJwtToken(v string) *CreateLocationServiceRequest {
	s.JwtToken = &v
	return s
}

func (s *CreateLocationServiceRequest) SetName(v string) *CreateLocationServiceRequest {
	s.Name = &v
	return s
}

func (s *CreateLocationServiceRequest) SetNote(v string) *CreateLocationServiceRequest {
	s.Note = &v
	return s
}

func (s *CreateLocationServiceRequest) SetQps(v int64) *CreateLocationServiceRequest {
	s.Qps = &v
	return s
}

func (s *CreateLocationServiceRequest) Validate() error {
	return dara.Validate(s)
}
