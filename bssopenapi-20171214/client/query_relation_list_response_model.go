// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRelationListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRelationListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRelationListResponse
	GetStatusCode() *int32
	SetBody(v *QueryRelationListResponseBody) *QueryRelationListResponse
	GetBody() *QueryRelationListResponseBody
}

type QueryRelationListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRelationListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRelationListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRelationListResponse) GoString() string {
	return s.String()
}

func (s *QueryRelationListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRelationListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRelationListResponse) GetBody() *QueryRelationListResponseBody {
	return s.Body
}

func (s *QueryRelationListResponse) SetHeaders(v map[string]*string) *QueryRelationListResponse {
	s.Headers = v
	return s
}

func (s *QueryRelationListResponse) SetStatusCode(v int32) *QueryRelationListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRelationListResponse) SetBody(v *QueryRelationListResponseBody) *QueryRelationListResponse {
	s.Body = v
	return s
}

func (s *QueryRelationListResponse) Validate() error {
	return dara.Validate(s)
}
