// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetFileProtectRuleRequest
	GetId() *int64
}

type GetFileProtectRuleRequest struct {
	// The ID of the rule.
	//
	// example:
	//
	// 245
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetFileProtectRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectRuleRequest) GoString() string {
	return s.String()
}

func (s *GetFileProtectRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *GetFileProtectRuleRequest) SetId(v int64) *GetFileProtectRuleRequest {
	s.Id = &v
	return s
}

func (s *GetFileProtectRuleRequest) Validate() error {
	return dara.Validate(s)
}
