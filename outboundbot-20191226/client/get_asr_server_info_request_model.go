// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsrServerInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntryId(v string) *GetAsrServerInfoRequest
	GetEntryId() *string
	SetStrategyLevel(v int32) *GetAsrServerInfoRequest
	GetStrategyLevel() *int32
}

type GetAsrServerInfoRequest struct {
	// Instance ID
	//
	// example:
	//
	// 024f8cf0-c842-4c01-b74b-c8667e4579c7
	EntryId *string `json:"EntryId,omitempty" xml:"EntryId,omitempty"`
	// Policy level
	//
	// example:
	//
	// 2
	StrategyLevel *int32 `json:"StrategyLevel,omitempty" xml:"StrategyLevel,omitempty"`
}

func (s GetAsrServerInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAsrServerInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAsrServerInfoRequest) GetEntryId() *string {
	return s.EntryId
}

func (s *GetAsrServerInfoRequest) GetStrategyLevel() *int32 {
	return s.StrategyLevel
}

func (s *GetAsrServerInfoRequest) SetEntryId(v string) *GetAsrServerInfoRequest {
	s.EntryId = &v
	return s
}

func (s *GetAsrServerInfoRequest) SetStrategyLevel(v int32) *GetAsrServerInfoRequest {
	s.StrategyLevel = &v
	return s
}

func (s *GetAsrServerInfoRequest) Validate() error {
	return dara.Validate(s)
}
