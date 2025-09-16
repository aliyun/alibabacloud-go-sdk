// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPostQueryResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v map[string]interface{}) *ListPostQueryResultRequest
	GetBody() map[string]interface{}
	SetType(v string) *ListPostQueryResultRequest
	GetType() *string
}

type ListPostQueryResultRequest struct {
	// The request body.
	//
	// example:
	//
	// {}
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
	// The query type. Valid values: sql: SQL query. ha3: Havenask query.
	//
	// example:
	//
	// ha3
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListPostQueryResultRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPostQueryResultRequest) GoString() string {
	return s.String()
}

func (s *ListPostQueryResultRequest) GetBody() map[string]interface{} {
	return s.Body
}

func (s *ListPostQueryResultRequest) GetType() *string {
	return s.Type
}

func (s *ListPostQueryResultRequest) SetBody(v map[string]interface{}) *ListPostQueryResultRequest {
	s.Body = v
	return s
}

func (s *ListPostQueryResultRequest) SetType(v string) *ListPostQueryResultRequest {
	s.Type = &v
	return s
}

func (s *ListPostQueryResultRequest) Validate() error {
	return dara.Validate(s)
}
