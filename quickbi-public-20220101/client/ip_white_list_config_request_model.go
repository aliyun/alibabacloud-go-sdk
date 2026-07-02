// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIpWhiteListConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpWhiteList(v string) *IpWhiteListConfigRequest
	GetIpWhiteList() *string
	SetOperation(v string) *IpWhiteListConfigRequest
	GetOperation() *string
}

type IpWhiteListConfigRequest struct {
	// Required for increase and delete operations. The IP whitelist. Separate multiple IP addresses with commas.
	//
	// example:
	//
	// 60.205.254.120
	IpWhiteList *string `json:"IpWhiteList,omitempty" xml:"IpWhiteList,omitempty"`
	// The operation type. Valid values:
	//
	// - ADD: incrementally adds entries.
	//
	// - DELETE: deletes entries.
	//
	// - QUERY: queries entries.
	//
	// This parameter is required.
	//
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
}

func (s IpWhiteListConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s IpWhiteListConfigRequest) GoString() string {
	return s.String()
}

func (s *IpWhiteListConfigRequest) GetIpWhiteList() *string {
	return s.IpWhiteList
}

func (s *IpWhiteListConfigRequest) GetOperation() *string {
	return s.Operation
}

func (s *IpWhiteListConfigRequest) SetIpWhiteList(v string) *IpWhiteListConfigRequest {
	s.IpWhiteList = &v
	return s
}

func (s *IpWhiteListConfigRequest) SetOperation(v string) *IpWhiteListConfigRequest {
	s.Operation = &v
	return s
}

func (s *IpWhiteListConfigRequest) Validate() error {
	return dara.Validate(s)
}
