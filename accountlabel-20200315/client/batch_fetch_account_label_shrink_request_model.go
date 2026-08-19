// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchFetchAccountLabelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstant(v bool) *BatchFetchAccountLabelShrinkRequest
	GetInstant() *bool
	SetLabelSeriesListShrink(v string) *BatchFetchAccountLabelShrinkRequest
	GetLabelSeriesListShrink() *string
	SetOrganization(v string) *BatchFetchAccountLabelShrinkRequest
	GetOrganization() *string
	SetPk(v int64) *BatchFetchAccountLabelShrinkRequest
	GetPk() *int64
	SetToken(v string) *BatchFetchAccountLabelShrinkRequest
	GetToken() *string
	SetUserName(v string) *BatchFetchAccountLabelShrinkRequest
	GetUserName() *string
}

type BatchFetchAccountLabelShrinkRequest struct {
	Instant *bool `json:"Instant,omitempty" xml:"Instant,omitempty"`
	// This parameter is required.
	LabelSeriesListShrink *string `json:"LabelSeriesList,omitempty" xml:"LabelSeriesList,omitempty"`
	// This parameter is required.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// This parameter is required.
	Pk *int64 `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// This parameter is required.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s BatchFetchAccountLabelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchFetchAccountLabelShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchFetchAccountLabelShrinkRequest) GetInstant() *bool {
	return s.Instant
}

func (s *BatchFetchAccountLabelShrinkRequest) GetLabelSeriesListShrink() *string {
	return s.LabelSeriesListShrink
}

func (s *BatchFetchAccountLabelShrinkRequest) GetOrganization() *string {
	return s.Organization
}

func (s *BatchFetchAccountLabelShrinkRequest) GetPk() *int64 {
	return s.Pk
}

func (s *BatchFetchAccountLabelShrinkRequest) GetToken() *string {
	return s.Token
}

func (s *BatchFetchAccountLabelShrinkRequest) GetUserName() *string {
	return s.UserName
}

func (s *BatchFetchAccountLabelShrinkRequest) SetInstant(v bool) *BatchFetchAccountLabelShrinkRequest {
	s.Instant = &v
	return s
}

func (s *BatchFetchAccountLabelShrinkRequest) SetLabelSeriesListShrink(v string) *BatchFetchAccountLabelShrinkRequest {
	s.LabelSeriesListShrink = &v
	return s
}

func (s *BatchFetchAccountLabelShrinkRequest) SetOrganization(v string) *BatchFetchAccountLabelShrinkRequest {
	s.Organization = &v
	return s
}

func (s *BatchFetchAccountLabelShrinkRequest) SetPk(v int64) *BatchFetchAccountLabelShrinkRequest {
	s.Pk = &v
	return s
}

func (s *BatchFetchAccountLabelShrinkRequest) SetToken(v string) *BatchFetchAccountLabelShrinkRequest {
	s.Token = &v
	return s
}

func (s *BatchFetchAccountLabelShrinkRequest) SetUserName(v string) *BatchFetchAccountLabelShrinkRequest {
	s.UserName = &v
	return s
}

func (s *BatchFetchAccountLabelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
