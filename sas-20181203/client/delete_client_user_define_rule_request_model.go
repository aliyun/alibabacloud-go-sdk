// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientUserDefineRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdList(v []*int64) *DeleteClientUserDefineRuleRequest
	GetIdList() []*int64
}

type DeleteClientUserDefineRuleRequest struct {
	// The IDs of the custom defense rules.
	//
	// This parameter is required.
	IdList []*int64 `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
}

func (s DeleteClientUserDefineRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientUserDefineRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteClientUserDefineRuleRequest) GetIdList() []*int64 {
	return s.IdList
}

func (s *DeleteClientUserDefineRuleRequest) SetIdList(v []*int64) *DeleteClientUserDefineRuleRequest {
	s.IdList = v
	return s
}

func (s *DeleteClientUserDefineRuleRequest) Validate() error {
	return dara.Validate(s)
}
