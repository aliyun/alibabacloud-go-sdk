// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSavedQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQueryId(v string) *GetSavedQueryRequest
	GetQueryId() *string
}

type GetSavedQueryRequest struct {
	// The template ID.
	//
	// >  You can call the [ListSavedQueries](~~ListSavedQueries~~) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sq-GeAck****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s GetSavedQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSavedQueryRequest) GoString() string {
	return s.String()
}

func (s *GetSavedQueryRequest) GetQueryId() *string {
	return s.QueryId
}

func (s *GetSavedQueryRequest) SetQueryId(v string) *GetSavedQueryRequest {
	s.QueryId = &v
	return s
}

func (s *GetSavedQueryRequest) Validate() error {
	return dara.Validate(s)
}
