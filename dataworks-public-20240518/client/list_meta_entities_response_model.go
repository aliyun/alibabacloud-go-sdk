// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaEntitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMetaEntitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMetaEntitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListMetaEntitiesResponseBody) *ListMetaEntitiesResponse
	GetBody() *ListMetaEntitiesResponseBody
}

type ListMetaEntitiesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMetaEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMetaEntitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMetaEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListMetaEntitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMetaEntitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMetaEntitiesResponse) GetBody() *ListMetaEntitiesResponseBody {
	return s.Body
}

func (s *ListMetaEntitiesResponse) SetHeaders(v map[string]*string) *ListMetaEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListMetaEntitiesResponse) SetStatusCode(v int32) *ListMetaEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMetaEntitiesResponse) SetBody(v *ListMetaEntitiesResponseBody) *ListMetaEntitiesResponse {
	s.Body = v
	return s
}

func (s *ListMetaEntitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
