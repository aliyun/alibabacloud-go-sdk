// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLayer7CustomPortsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListLayer7CustomPortsRequest
	GetLang() *string
	SetResourceGroupId(v string) *ListLayer7CustomPortsRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *ListLayer7CustomPortsRequest
	GetSourceIp() *string
}

type ListLayer7CustomPortsRequest struct {
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ListLayer7CustomPortsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLayer7CustomPortsRequest) GoString() string {
	return s.String()
}

func (s *ListLayer7CustomPortsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListLayer7CustomPortsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListLayer7CustomPortsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ListLayer7CustomPortsRequest) SetLang(v string) *ListLayer7CustomPortsRequest {
	s.Lang = &v
	return s
}

func (s *ListLayer7CustomPortsRequest) SetResourceGroupId(v string) *ListLayer7CustomPortsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLayer7CustomPortsRequest) SetSourceIp(v string) *ListLayer7CustomPortsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListLayer7CustomPortsRequest) Validate() error {
	return dara.Validate(s)
}
