// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaCollectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMetaCollectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMetaCollectionsResponse
	GetStatusCode() *int32
	SetBody(v *ListMetaCollectionsResponseBody) *ListMetaCollectionsResponse
	GetBody() *ListMetaCollectionsResponseBody
}

type ListMetaCollectionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMetaCollectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMetaCollectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMetaCollectionsResponse) GoString() string {
	return s.String()
}

func (s *ListMetaCollectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMetaCollectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMetaCollectionsResponse) GetBody() *ListMetaCollectionsResponseBody {
	return s.Body
}

func (s *ListMetaCollectionsResponse) SetHeaders(v map[string]*string) *ListMetaCollectionsResponse {
	s.Headers = v
	return s
}

func (s *ListMetaCollectionsResponse) SetStatusCode(v int32) *ListMetaCollectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMetaCollectionsResponse) SetBody(v *ListMetaCollectionsResponseBody) *ListMetaCollectionsResponse {
	s.Body = v
	return s
}

func (s *ListMetaCollectionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
