// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAppSeoIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *SubmitAppSeoIndexRequest
	GetBizId() *string
	SetDomain(v string) *SubmitAppSeoIndexRequest
	GetDomain() *string
	SetSeType(v string) *SubmitAppSeoIndexRequest
	GetSeType() *string
	SetSubmitLater(v bool) *SubmitAppSeoIndexRequest
	GetSubmitLater() *bool
}

type SubmitAppSeoIndexRequest struct {
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// yjdw.bpu.edu.cn-waf
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// type
	SeType      *string `json:"SeType,omitempty" xml:"SeType,omitempty"`
	SubmitLater *bool   `json:"SubmitLater,omitempty" xml:"SubmitLater,omitempty"`
}

func (s SubmitAppSeoIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAppSeoIndexRequest) GoString() string {
	return s.String()
}

func (s *SubmitAppSeoIndexRequest) GetBizId() *string {
	return s.BizId
}

func (s *SubmitAppSeoIndexRequest) GetDomain() *string {
	return s.Domain
}

func (s *SubmitAppSeoIndexRequest) GetSeType() *string {
	return s.SeType
}

func (s *SubmitAppSeoIndexRequest) GetSubmitLater() *bool {
	return s.SubmitLater
}

func (s *SubmitAppSeoIndexRequest) SetBizId(v string) *SubmitAppSeoIndexRequest {
	s.BizId = &v
	return s
}

func (s *SubmitAppSeoIndexRequest) SetDomain(v string) *SubmitAppSeoIndexRequest {
	s.Domain = &v
	return s
}

func (s *SubmitAppSeoIndexRequest) SetSeType(v string) *SubmitAppSeoIndexRequest {
	s.SeType = &v
	return s
}

func (s *SubmitAppSeoIndexRequest) SetSubmitLater(v bool) *SubmitAppSeoIndexRequest {
	s.SubmitLater = &v
	return s
}

func (s *SubmitAppSeoIndexRequest) Validate() error {
	return dara.Validate(s)
}
