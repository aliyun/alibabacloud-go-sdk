// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAntiBruteForceRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v []*int64) *DeleteAntiBruteForceRuleRequest
	GetIds() []*int64
}

type DeleteAntiBruteForceRuleRequest struct {
	// The IDs of the defense rules against brute-force attacks to delete.
	//
	// This parameter is required.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s DeleteAntiBruteForceRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAntiBruteForceRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAntiBruteForceRuleRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *DeleteAntiBruteForceRuleRequest) SetIds(v []*int64) *DeleteAntiBruteForceRuleRequest {
	s.Ids = v
	return s
}

func (s *DeleteAntiBruteForceRuleRequest) Validate() error {
	return dara.Validate(s)
}
