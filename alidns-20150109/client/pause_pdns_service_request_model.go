// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPausePdnsServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *PausePdnsServiceRequest
	GetLang() *string
	SetServiceType(v string) *PausePdnsServiceRequest
	GetServiceType() *string
}

type PausePdnsServiceRequest struct {
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s PausePdnsServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s PausePdnsServiceRequest) GoString() string {
	return s.String()
}

func (s *PausePdnsServiceRequest) GetLang() *string {
	return s.Lang
}

func (s *PausePdnsServiceRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *PausePdnsServiceRequest) SetLang(v string) *PausePdnsServiceRequest {
	s.Lang = &v
	return s
}

func (s *PausePdnsServiceRequest) SetServiceType(v string) *PausePdnsServiceRequest {
	s.ServiceType = &v
	return s
}

func (s *PausePdnsServiceRequest) Validate() error {
	return dara.Validate(s)
}
