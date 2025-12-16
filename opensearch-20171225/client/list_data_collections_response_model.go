// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCollectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataCollectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataCollectionsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataCollectionsResponseBody) *ListDataCollectionsResponse
	GetBody() *ListDataCollectionsResponseBody
}

type ListDataCollectionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataCollectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataCollectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataCollectionsResponse) GoString() string {
	return s.String()
}

func (s *ListDataCollectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataCollectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataCollectionsResponse) GetBody() *ListDataCollectionsResponseBody {
	return s.Body
}

func (s *ListDataCollectionsResponse) SetHeaders(v map[string]*string) *ListDataCollectionsResponse {
	s.Headers = v
	return s
}

func (s *ListDataCollectionsResponse) SetStatusCode(v int32) *ListDataCollectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataCollectionsResponse) SetBody(v *ListDataCollectionsResponseBody) *ListDataCollectionsResponse {
	s.Body = v
	return s
}

func (s *ListDataCollectionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
