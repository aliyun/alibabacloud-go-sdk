// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBizEntitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBizEntitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBizEntitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListBizEntitiesResponseBody) *ListBizEntitiesResponse
	GetBody() *ListBizEntitiesResponseBody
}

type ListBizEntitiesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBizEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBizEntitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBizEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListBizEntitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBizEntitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBizEntitiesResponse) GetBody() *ListBizEntitiesResponseBody {
	return s.Body
}

func (s *ListBizEntitiesResponse) SetHeaders(v map[string]*string) *ListBizEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListBizEntitiesResponse) SetStatusCode(v int32) *ListBizEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBizEntitiesResponse) SetBody(v *ListBizEntitiesResponseBody) *ListBizEntitiesResponse {
	s.Body = v
	return s
}

func (s *ListBizEntitiesResponse) Validate() error {
	return dara.Validate(s)
}
