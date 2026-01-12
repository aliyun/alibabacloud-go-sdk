// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryCollectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMemoryCollectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMemoryCollectionsResponse
	GetStatusCode() *int32
	SetBody(v *ListMemoryCollectionsResult) *ListMemoryCollectionsResponse
	GetBody() *ListMemoryCollectionsResult
}

type ListMemoryCollectionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMemoryCollectionsResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMemoryCollectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryCollectionsResponse) GoString() string {
	return s.String()
}

func (s *ListMemoryCollectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMemoryCollectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMemoryCollectionsResponse) GetBody() *ListMemoryCollectionsResult {
	return s.Body
}

func (s *ListMemoryCollectionsResponse) SetHeaders(v map[string]*string) *ListMemoryCollectionsResponse {
	s.Headers = v
	return s
}

func (s *ListMemoryCollectionsResponse) SetStatusCode(v int32) *ListMemoryCollectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMemoryCollectionsResponse) SetBody(v *ListMemoryCollectionsResult) *ListMemoryCollectionsResponse {
	s.Body = v
	return s
}

func (s *ListMemoryCollectionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
