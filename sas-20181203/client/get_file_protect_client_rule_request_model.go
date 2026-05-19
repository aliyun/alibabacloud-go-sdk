// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectClientRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetFileProtectClientRuleRequest
	GetId() *int64
}

type GetFileProtectClientRuleRequest struct {
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetFileProtectClientRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectClientRuleRequest) GoString() string {
	return s.String()
}

func (s *GetFileProtectClientRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *GetFileProtectClientRuleRequest) SetId(v int64) *GetFileProtectClientRuleRequest {
	s.Id = &v
	return s
}

func (s *GetFileProtectClientRuleRequest) Validate() error {
	return dara.Validate(s)
}
