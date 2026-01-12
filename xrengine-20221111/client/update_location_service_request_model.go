// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLocationServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *UpdateLocationServiceRequest
	GetId() *int64
	SetJwtToken(v string) *UpdateLocationServiceRequest
	GetJwtToken() *string
	SetNote(v string) *UpdateLocationServiceRequest
	GetNote() *string
	SetQps(v int32) *UpdateLocationServiceRequest
	GetQps() *int32
	SetSvcName(v string) *UpdateLocationServiceRequest
	GetSvcName() *string
	SetSvcState(v string) *UpdateLocationServiceRequest
	GetSvcState() *string
}

type UpdateLocationServiceRequest struct {
	// This parameter is required.
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Note     *string `json:"Note,omitempty" xml:"Note,omitempty"`
	Qps      *int32  `json:"Qps,omitempty" xml:"Qps,omitempty"`
	SvcName  *string `json:"SvcName,omitempty" xml:"SvcName,omitempty"`
	SvcState *string `json:"SvcState,omitempty" xml:"SvcState,omitempty"`
}

func (s UpdateLocationServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLocationServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateLocationServiceRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateLocationServiceRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *UpdateLocationServiceRequest) GetNote() *string {
	return s.Note
}

func (s *UpdateLocationServiceRequest) GetQps() *int32 {
	return s.Qps
}

func (s *UpdateLocationServiceRequest) GetSvcName() *string {
	return s.SvcName
}

func (s *UpdateLocationServiceRequest) GetSvcState() *string {
	return s.SvcState
}

func (s *UpdateLocationServiceRequest) SetId(v int64) *UpdateLocationServiceRequest {
	s.Id = &v
	return s
}

func (s *UpdateLocationServiceRequest) SetJwtToken(v string) *UpdateLocationServiceRequest {
	s.JwtToken = &v
	return s
}

func (s *UpdateLocationServiceRequest) SetNote(v string) *UpdateLocationServiceRequest {
	s.Note = &v
	return s
}

func (s *UpdateLocationServiceRequest) SetQps(v int32) *UpdateLocationServiceRequest {
	s.Qps = &v
	return s
}

func (s *UpdateLocationServiceRequest) SetSvcName(v string) *UpdateLocationServiceRequest {
	s.SvcName = &v
	return s
}

func (s *UpdateLocationServiceRequest) SetSvcState(v string) *UpdateLocationServiceRequest {
	s.SvcState = &v
	return s
}

func (s *UpdateLocationServiceRequest) Validate() error {
	return dara.Validate(s)
}
