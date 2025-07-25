// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIspFlushCacheTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SubmitIspFlushCacheTaskRequest
	GetClientToken() *string
	SetDomainName(v string) *SubmitIspFlushCacheTaskRequest
	GetDomainName() *string
	SetIsp(v []*string) *SubmitIspFlushCacheTaskRequest
	GetIsp() []*string
	SetLang(v string) *SubmitIspFlushCacheTaskRequest
	GetLang() *string
}

type SubmitIspFlushCacheTaskRequest struct {
	// This parameter is required.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	Isp  []*string `json:"Isp,omitempty" xml:"Isp,omitempty" type:"Repeated"`
	Lang *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s SubmitIspFlushCacheTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitIspFlushCacheTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitIspFlushCacheTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SubmitIspFlushCacheTaskRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SubmitIspFlushCacheTaskRequest) GetIsp() []*string {
	return s.Isp
}

func (s *SubmitIspFlushCacheTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *SubmitIspFlushCacheTaskRequest) SetClientToken(v string) *SubmitIspFlushCacheTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *SubmitIspFlushCacheTaskRequest) SetDomainName(v string) *SubmitIspFlushCacheTaskRequest {
	s.DomainName = &v
	return s
}

func (s *SubmitIspFlushCacheTaskRequest) SetIsp(v []*string) *SubmitIspFlushCacheTaskRequest {
	s.Isp = v
	return s
}

func (s *SubmitIspFlushCacheTaskRequest) SetLang(v string) *SubmitIspFlushCacheTaskRequest {
	s.Lang = &v
	return s
}

func (s *SubmitIspFlushCacheTaskRequest) Validate() error {
	return dara.Validate(s)
}
