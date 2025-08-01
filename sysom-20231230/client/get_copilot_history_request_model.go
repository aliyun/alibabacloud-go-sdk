// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCopilotHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *GetCopilotHistoryRequest
	GetCount() *int64
}

type GetCopilotHistoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 100
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
}

func (s GetCopilotHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCopilotHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetCopilotHistoryRequest) GetCount() *int64 {
	return s.Count
}

func (s *GetCopilotHistoryRequest) SetCount(v int64) *GetCopilotHistoryRequest {
	s.Count = &v
	return s
}

func (s *GetCopilotHistoryRequest) Validate() error {
	return dara.Validate(s)
}
