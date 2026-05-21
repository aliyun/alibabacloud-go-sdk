// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDatasetResourceUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenDatasetResourceUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenDatasetResourceUrlResponse
	GetStatusCode() *int32
	SetBody(v *OpenDatasetResourceUrlResponseBody) *OpenDatasetResourceUrlResponse
	GetBody() *OpenDatasetResourceUrlResponseBody
}

type OpenDatasetResourceUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenDatasetResourceUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenDatasetResourceUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenDatasetResourceUrlResponse) GoString() string {
	return s.String()
}

func (s *OpenDatasetResourceUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenDatasetResourceUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenDatasetResourceUrlResponse) GetBody() *OpenDatasetResourceUrlResponseBody {
	return s.Body
}

func (s *OpenDatasetResourceUrlResponse) SetHeaders(v map[string]*string) *OpenDatasetResourceUrlResponse {
	s.Headers = v
	return s
}

func (s *OpenDatasetResourceUrlResponse) SetStatusCode(v int32) *OpenDatasetResourceUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenDatasetResourceUrlResponse) SetBody(v *OpenDatasetResourceUrlResponseBody) *OpenDatasetResourceUrlResponse {
	s.Body = v
	return s
}

func (s *OpenDatasetResourceUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
