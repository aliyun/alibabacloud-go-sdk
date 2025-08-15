// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAdvancedQueryHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQueryId(v string) *DeleteAdvancedQueryHistoryRequest
	GetQueryId() *string
}

type DeleteAdvancedQueryHistoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// query-uIkIvLiVSuCKqg0yoa****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s DeleteAdvancedQueryHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAdvancedQueryHistoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteAdvancedQueryHistoryRequest) GetQueryId() *string {
	return s.QueryId
}

func (s *DeleteAdvancedQueryHistoryRequest) SetQueryId(v string) *DeleteAdvancedQueryHistoryRequest {
	s.QueryId = &v
	return s
}

func (s *DeleteAdvancedQueryHistoryRequest) Validate() error {
	return dara.Validate(s)
}
