// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOcspStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ModifyOcspStatusRequest
	GetDomain() *string
	SetEnable(v int32) *ModifyOcspStatusRequest
	GetEnable() *int32
}

type ModifyOcspStatusRequest struct {
	// The domain name for which you want to configure the Static Page Caching policy.
	//
	// > You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all the domain names that are added to Anti-DDoS Pro or Anti-DDoS Premium.
	//
	// This parameter is required.
	//
	// example:
	//
	// click.linktech.cn
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether to enable the OCSP feature. Valid values:
	//
	// 	- **1**: yes.
	//
	// 	- **0**: no.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s ModifyOcspStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyOcspStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyOcspStatusRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyOcspStatusRequest) GetEnable() *int32 {
	return s.Enable
}

func (s *ModifyOcspStatusRequest) SetDomain(v string) *ModifyOcspStatusRequest {
	s.Domain = &v
	return s
}

func (s *ModifyOcspStatusRequest) SetEnable(v int32) *ModifyOcspStatusRequest {
	s.Enable = &v
	return s
}

func (s *ModifyOcspStatusRequest) Validate() error {
	return dara.Validate(s)
}
