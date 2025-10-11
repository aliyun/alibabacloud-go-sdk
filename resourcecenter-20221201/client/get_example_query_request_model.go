// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExampleQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQueryId(v string) *GetExampleQueryRequest
	GetQueryId() *string
}

type GetExampleQueryRequest struct {
	// The ID of the template.
	//
	// >  You can call the [ListExampleQueries](~~ListExampleQueries~~) operation to obtain the template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sq-0PfKy****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s GetExampleQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExampleQueryRequest) GoString() string {
	return s.String()
}

func (s *GetExampleQueryRequest) GetQueryId() *string {
	return s.QueryId
}

func (s *GetExampleQueryRequest) SetQueryId(v string) *GetExampleQueryRequest {
	s.QueryId = &v
	return s
}

func (s *GetExampleQueryRequest) Validate() error {
	return dara.Validate(s)
}
