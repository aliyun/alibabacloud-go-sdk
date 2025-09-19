// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientUserDefineRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetClientUserDefineRuleRequest
	GetId() *int64
}

type GetClientUserDefineRuleRequest struct {
	// The ID of the custom defense rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetClientUserDefineRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClientUserDefineRuleRequest) GoString() string {
	return s.String()
}

func (s *GetClientUserDefineRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *GetClientUserDefineRuleRequest) SetId(v int64) *GetClientUserDefineRuleRequest {
	s.Id = &v
	return s
}

func (s *GetClientUserDefineRuleRequest) Validate() error {
	return dara.Validate(s)
}
