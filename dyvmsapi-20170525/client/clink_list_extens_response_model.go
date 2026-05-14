// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListExtensResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkListExtensResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkListExtensResponse
	GetStatusCode() *int32
	SetBody(v *ClinkListExtensResponseBody) *ClinkListExtensResponse
	GetBody() *ClinkListExtensResponseBody
}

type ClinkListExtensResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkListExtensResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkListExtensResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkListExtensResponse) GoString() string {
	return s.String()
}

func (s *ClinkListExtensResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkListExtensResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkListExtensResponse) GetBody() *ClinkListExtensResponseBody {
	return s.Body
}

func (s *ClinkListExtensResponse) SetHeaders(v map[string]*string) *ClinkListExtensResponse {
	s.Headers = v
	return s
}

func (s *ClinkListExtensResponse) SetStatusCode(v int32) *ClinkListExtensResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkListExtensResponse) SetBody(v *ClinkListExtensResponseBody) *ClinkListExtensResponse {
	s.Body = v
	return s
}

func (s *ClinkListExtensResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
