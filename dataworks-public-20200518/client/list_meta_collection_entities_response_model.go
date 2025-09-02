// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaCollectionEntitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMetaCollectionEntitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMetaCollectionEntitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListMetaCollectionEntitiesResponseBody) *ListMetaCollectionEntitiesResponse
	GetBody() *ListMetaCollectionEntitiesResponseBody
}

type ListMetaCollectionEntitiesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMetaCollectionEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMetaCollectionEntitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMetaCollectionEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListMetaCollectionEntitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMetaCollectionEntitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMetaCollectionEntitiesResponse) GetBody() *ListMetaCollectionEntitiesResponseBody {
	return s.Body
}

func (s *ListMetaCollectionEntitiesResponse) SetHeaders(v map[string]*string) *ListMetaCollectionEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListMetaCollectionEntitiesResponse) SetStatusCode(v int32) *ListMetaCollectionEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMetaCollectionEntitiesResponse) SetBody(v *ListMetaCollectionEntitiesResponseBody) *ListMetaCollectionEntitiesResponse {
	s.Body = v
	return s
}

func (s *ListMetaCollectionEntitiesResponse) Validate() error {
	return dara.Validate(s)
}
