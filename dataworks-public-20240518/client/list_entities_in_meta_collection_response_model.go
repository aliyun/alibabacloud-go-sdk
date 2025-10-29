// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEntitiesInMetaCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEntitiesInMetaCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEntitiesInMetaCollectionResponse
	GetStatusCode() *int32
	SetBody(v *ListEntitiesInMetaCollectionResponseBody) *ListEntitiesInMetaCollectionResponse
	GetBody() *ListEntitiesInMetaCollectionResponseBody
}

type ListEntitiesInMetaCollectionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEntitiesInMetaCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEntitiesInMetaCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesInMetaCollectionResponse) GoString() string {
	return s.String()
}

func (s *ListEntitiesInMetaCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEntitiesInMetaCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEntitiesInMetaCollectionResponse) GetBody() *ListEntitiesInMetaCollectionResponseBody {
	return s.Body
}

func (s *ListEntitiesInMetaCollectionResponse) SetHeaders(v map[string]*string) *ListEntitiesInMetaCollectionResponse {
	s.Headers = v
	return s
}

func (s *ListEntitiesInMetaCollectionResponse) SetStatusCode(v int32) *ListEntitiesInMetaCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEntitiesInMetaCollectionResponse) SetBody(v *ListEntitiesInMetaCollectionResponseBody) *ListEntitiesInMetaCollectionResponse {
	s.Body = v
	return s
}

func (s *ListEntitiesInMetaCollectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
