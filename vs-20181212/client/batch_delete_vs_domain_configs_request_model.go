// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteVsDomainConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *BatchDeleteVsDomainConfigsRequest
	GetDomainNames() *string
	SetFunctionNames(v string) *BatchDeleteVsDomainConfigsRequest
	GetFunctionNames() *string
	SetOwnerId(v int64) *BatchDeleteVsDomainConfigsRequest
	GetOwnerId() *int64
}

type BatchDeleteVsDomainConfigsRequest struct {
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
	// https_force
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s BatchDeleteVsDomainConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteVsDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteVsDomainConfigsRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *BatchDeleteVsDomainConfigsRequest) GetFunctionNames() *string {
	return s.FunctionNames
}

func (s *BatchDeleteVsDomainConfigsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchDeleteVsDomainConfigsRequest) SetDomainNames(v string) *BatchDeleteVsDomainConfigsRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchDeleteVsDomainConfigsRequest) SetFunctionNames(v string) *BatchDeleteVsDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

func (s *BatchDeleteVsDomainConfigsRequest) SetOwnerId(v int64) *BatchDeleteVsDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchDeleteVsDomainConfigsRequest) Validate() error {
	return dara.Validate(s)
}
