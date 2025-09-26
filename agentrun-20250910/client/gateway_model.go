// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGateway interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *Gateway
	GetCreatedAt() *string
	SetGatewayId(v string) *Gateway
	GetGatewayId() *string
	SetInternetUrl(v string) *Gateway
	GetInternetUrl() *string
	SetIntranetUrl(v string) *Gateway
	GetIntranetUrl() *string
	SetName(v string) *Gateway
	GetName() *string
	SetStatus(v string) *Gateway
	GetStatus() *string
	SetUpdatedAt(v string) *Gateway
	GetUpdatedAt() *string
}

type Gateway struct {
	CreatedAt   *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	GatewayId   *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	InternetUrl *string `json:"internetUrl,omitempty" xml:"internetUrl,omitempty"`
	IntranetUrl *string `json:"intranetUrl,omitempty" xml:"intranetUrl,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
	UpdatedAt   *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s Gateway) String() string {
	return dara.Prettify(s)
}

func (s Gateway) GoString() string {
	return s.String()
}

func (s *Gateway) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Gateway) GetGatewayId() *string {
	return s.GatewayId
}

func (s *Gateway) GetInternetUrl() *string {
	return s.InternetUrl
}

func (s *Gateway) GetIntranetUrl() *string {
	return s.IntranetUrl
}

func (s *Gateway) GetName() *string {
	return s.Name
}

func (s *Gateway) GetStatus() *string {
	return s.Status
}

func (s *Gateway) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *Gateway) SetCreatedAt(v string) *Gateway {
	s.CreatedAt = &v
	return s
}

func (s *Gateway) SetGatewayId(v string) *Gateway {
	s.GatewayId = &v
	return s
}

func (s *Gateway) SetInternetUrl(v string) *Gateway {
	s.InternetUrl = &v
	return s
}

func (s *Gateway) SetIntranetUrl(v string) *Gateway {
	s.IntranetUrl = &v
	return s
}

func (s *Gateway) SetName(v string) *Gateway {
	s.Name = &v
	return s
}

func (s *Gateway) SetStatus(v string) *Gateway {
	s.Status = &v
	return s
}

func (s *Gateway) SetUpdatedAt(v string) *Gateway {
	s.UpdatedAt = &v
	return s
}

func (s *Gateway) Validate() error {
	return dara.Validate(s)
}
