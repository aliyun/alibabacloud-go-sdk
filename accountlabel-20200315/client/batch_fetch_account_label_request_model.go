// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchFetchAccountLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstant(v bool) *BatchFetchAccountLabelRequest
	GetInstant() *bool
	SetLabelSeriesList(v []*string) *BatchFetchAccountLabelRequest
	GetLabelSeriesList() []*string
	SetOrganization(v string) *BatchFetchAccountLabelRequest
	GetOrganization() *string
	SetPk(v int64) *BatchFetchAccountLabelRequest
	GetPk() *int64
	SetToken(v string) *BatchFetchAccountLabelRequest
	GetToken() *string
	SetUserName(v string) *BatchFetchAccountLabelRequest
	GetUserName() *string
}

type BatchFetchAccountLabelRequest struct {
	Instant *bool `json:"Instant,omitempty" xml:"Instant,omitempty"`
	// This parameter is required.
	LabelSeriesList []*string `json:"LabelSeriesList,omitempty" xml:"LabelSeriesList,omitempty" type:"Repeated"`
	// This parameter is required.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// This parameter is required.
	Pk *int64 `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// This parameter is required.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s BatchFetchAccountLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchFetchAccountLabelRequest) GoString() string {
	return s.String()
}

func (s *BatchFetchAccountLabelRequest) GetInstant() *bool {
	return s.Instant
}

func (s *BatchFetchAccountLabelRequest) GetLabelSeriesList() []*string {
	return s.LabelSeriesList
}

func (s *BatchFetchAccountLabelRequest) GetOrganization() *string {
	return s.Organization
}

func (s *BatchFetchAccountLabelRequest) GetPk() *int64 {
	return s.Pk
}

func (s *BatchFetchAccountLabelRequest) GetToken() *string {
	return s.Token
}

func (s *BatchFetchAccountLabelRequest) GetUserName() *string {
	return s.UserName
}

func (s *BatchFetchAccountLabelRequest) SetInstant(v bool) *BatchFetchAccountLabelRequest {
	s.Instant = &v
	return s
}

func (s *BatchFetchAccountLabelRequest) SetLabelSeriesList(v []*string) *BatchFetchAccountLabelRequest {
	s.LabelSeriesList = v
	return s
}

func (s *BatchFetchAccountLabelRequest) SetOrganization(v string) *BatchFetchAccountLabelRequest {
	s.Organization = &v
	return s
}

func (s *BatchFetchAccountLabelRequest) SetPk(v int64) *BatchFetchAccountLabelRequest {
	s.Pk = &v
	return s
}

func (s *BatchFetchAccountLabelRequest) SetToken(v string) *BatchFetchAccountLabelRequest {
	s.Token = &v
	return s
}

func (s *BatchFetchAccountLabelRequest) SetUserName(v string) *BatchFetchAccountLabelRequest {
	s.UserName = &v
	return s
}

func (s *BatchFetchAccountLabelRequest) Validate() error {
	return dara.Validate(s)
}
