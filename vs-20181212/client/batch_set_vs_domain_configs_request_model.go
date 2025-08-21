// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetVsDomainConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *BatchSetVsDomainConfigsRequest
	GetDomainNames() *string
	SetFunctions(v string) *BatchSetVsDomainConfigsRequest
	GetFunctions() *string
	SetOwnerId(v int64) *BatchSetVsDomainConfigsRequest
	GetOwnerId() *int64
}

type BatchSetVsDomainConfigsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{“functionArgs”:[{“argName”:”domain_name”,”argValue”:”api.hellodtworld.com”}],”functionName”:”set_req_host_header”}]
	Functions *string `json:"Functions,omitempty" xml:"Functions,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s BatchSetVsDomainConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSetVsDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *BatchSetVsDomainConfigsRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *BatchSetVsDomainConfigsRequest) GetFunctions() *string {
	return s.Functions
}

func (s *BatchSetVsDomainConfigsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchSetVsDomainConfigsRequest) SetDomainNames(v string) *BatchSetVsDomainConfigsRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchSetVsDomainConfigsRequest) SetFunctions(v string) *BatchSetVsDomainConfigsRequest {
	s.Functions = &v
	return s
}

func (s *BatchSetVsDomainConfigsRequest) SetOwnerId(v int64) *BatchSetVsDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchSetVsDomainConfigsRequest) Validate() error {
	return dara.Validate(s)
}
