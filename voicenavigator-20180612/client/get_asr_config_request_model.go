// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsrConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigLevel(v int32) *GetAsrConfigRequest
	GetConfigLevel() *int32
	SetEntryId(v string) *GetAsrConfigRequest
	GetEntryId() *string
}

type GetAsrConfigRequest struct {
	// example:
	//
	// 1
	ConfigLevel *int32 `json:"ConfigLevel,omitempty" xml:"ConfigLevel,omitempty"`
	// example:
	//
	// 3b1d3031-8b6e-460a-8640-d330f2ca50b8
	EntryId *string `json:"EntryId,omitempty" xml:"EntryId,omitempty"`
}

func (s GetAsrConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAsrConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAsrConfigRequest) GetConfigLevel() *int32 {
	return s.ConfigLevel
}

func (s *GetAsrConfigRequest) GetEntryId() *string {
	return s.EntryId
}

func (s *GetAsrConfigRequest) SetConfigLevel(v int32) *GetAsrConfigRequest {
	s.ConfigLevel = &v
	return s
}

func (s *GetAsrConfigRequest) SetEntryId(v string) *GetAsrConfigRequest {
	s.EntryId = &v
	return s
}

func (s *GetAsrConfigRequest) Validate() error {
	return dara.Validate(s)
}
