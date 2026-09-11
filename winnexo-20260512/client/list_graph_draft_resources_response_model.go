// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGraphDraftResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGraphDraftResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGraphDraftResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListGraphDraftResourcesResponseBody) *ListGraphDraftResourcesResponse
	GetBody() *ListGraphDraftResourcesResponseBody
}

type ListGraphDraftResourcesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGraphDraftResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGraphDraftResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGraphDraftResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListGraphDraftResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGraphDraftResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGraphDraftResourcesResponse) GetBody() *ListGraphDraftResourcesResponseBody {
	return s.Body
}

func (s *ListGraphDraftResourcesResponse) SetHeaders(v map[string]*string) *ListGraphDraftResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListGraphDraftResourcesResponse) SetStatusCode(v int32) *ListGraphDraftResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGraphDraftResourcesResponse) SetBody(v *ListGraphDraftResourcesResponseBody) *ListGraphDraftResourcesResponse {
	s.Body = v
	return s
}

func (s *ListGraphDraftResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
