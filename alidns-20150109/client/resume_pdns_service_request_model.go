// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumePdnsServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ResumePdnsServiceRequest
	GetLang() *string
	SetServiceType(v string) *ResumePdnsServiceRequest
	GetServiceType() *string
}

type ResumePdnsServiceRequest struct {
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s ResumePdnsServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumePdnsServiceRequest) GoString() string {
	return s.String()
}

func (s *ResumePdnsServiceRequest) GetLang() *string {
	return s.Lang
}

func (s *ResumePdnsServiceRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *ResumePdnsServiceRequest) SetLang(v string) *ResumePdnsServiceRequest {
	s.Lang = &v
	return s
}

func (s *ResumePdnsServiceRequest) SetServiceType(v string) *ResumePdnsServiceRequest {
	s.ServiceType = &v
	return s
}

func (s *ResumePdnsServiceRequest) Validate() error {
	return dara.Validate(s)
}
