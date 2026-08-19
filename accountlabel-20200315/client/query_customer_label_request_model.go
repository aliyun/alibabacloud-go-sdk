// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomerLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstant(v bool) *QueryCustomerLabelRequest
	GetInstant() *bool
	SetLabelSeries(v string) *QueryCustomerLabelRequest
	GetLabelSeries() *string
	SetPK(v int64) *QueryCustomerLabelRequest
	GetPK() *int64
	SetToken(v string) *QueryCustomerLabelRequest
	GetToken() *string
}

type QueryCustomerLabelRequest struct {
	Instant     *bool   `json:"Instant,omitempty" xml:"Instant,omitempty"`
	LabelSeries *string `json:"LabelSeries,omitempty" xml:"LabelSeries,omitempty"`
	// This parameter is required.
	PK    *int64  `json:"PK,omitempty" xml:"PK,omitempty"`
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s QueryCustomerLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomerLabelRequest) GoString() string {
	return s.String()
}

func (s *QueryCustomerLabelRequest) GetInstant() *bool {
	return s.Instant
}

func (s *QueryCustomerLabelRequest) GetLabelSeries() *string {
	return s.LabelSeries
}

func (s *QueryCustomerLabelRequest) GetPK() *int64 {
	return s.PK
}

func (s *QueryCustomerLabelRequest) GetToken() *string {
	return s.Token
}

func (s *QueryCustomerLabelRequest) SetInstant(v bool) *QueryCustomerLabelRequest {
	s.Instant = &v
	return s
}

func (s *QueryCustomerLabelRequest) SetLabelSeries(v string) *QueryCustomerLabelRequest {
	s.LabelSeries = &v
	return s
}

func (s *QueryCustomerLabelRequest) SetPK(v int64) *QueryCustomerLabelRequest {
	s.PK = &v
	return s
}

func (s *QueryCustomerLabelRequest) SetToken(v string) *QueryCustomerLabelRequest {
	s.Token = &v
	return s
}

func (s *QueryCustomerLabelRequest) Validate() error {
	return dara.Validate(s)
}
