// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSavedQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQueryId(v string) *DeleteSavedQueryRequest
	GetQueryId() *string
}

type DeleteSavedQueryRequest struct {
	// The ID of the template.
	//
	// You can call the [ListSavedQueries](~~ListSavedQueries~~) operation to obtain the template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sq-GeAck****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s DeleteSavedQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSavedQueryRequest) GoString() string {
	return s.String()
}

func (s *DeleteSavedQueryRequest) GetQueryId() *string {
	return s.QueryId
}

func (s *DeleteSavedQueryRequest) SetQueryId(v string) *DeleteSavedQueryRequest {
	s.QueryId = &v
	return s
}

func (s *DeleteSavedQueryRequest) Validate() error {
	return dara.Validate(s)
}
