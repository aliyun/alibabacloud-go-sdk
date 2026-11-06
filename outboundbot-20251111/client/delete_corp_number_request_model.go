// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCorpNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *DeleteCorpNumberRequest
	GetForce() *bool
	SetNumber(v string) *DeleteCorpNumberRequest
	GetNumber() *string
	SetTenantId(v string) *DeleteCorpNumberRequest
	GetTenantId() *string
}

type DeleteCorpNumberRequest struct {
	// Specifies whether to force delete an existing number.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The outbound number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 057100001234
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1234567890
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DeleteCorpNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCorpNumberRequest) GoString() string {
	return s.String()
}

func (s *DeleteCorpNumberRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteCorpNumberRequest) GetNumber() *string {
	return s.Number
}

func (s *DeleteCorpNumberRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteCorpNumberRequest) SetForce(v bool) *DeleteCorpNumberRequest {
	s.Force = &v
	return s
}

func (s *DeleteCorpNumberRequest) SetNumber(v string) *DeleteCorpNumberRequest {
	s.Number = &v
	return s
}

func (s *DeleteCorpNumberRequest) SetTenantId(v string) *DeleteCorpNumberRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteCorpNumberRequest) Validate() error {
	return dara.Validate(s)
}
