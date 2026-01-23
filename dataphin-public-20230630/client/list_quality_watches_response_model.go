// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityWatchesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQualityWatchesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQualityWatchesResponse
	GetStatusCode() *int32
	SetBody(v *ListQualityWatchesResponseBody) *ListQualityWatchesResponse
	GetBody() *ListQualityWatchesResponseBody
}

type ListQualityWatchesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQualityWatchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQualityWatchesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchesResponse) GoString() string {
	return s.String()
}

func (s *ListQualityWatchesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQualityWatchesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQualityWatchesResponse) GetBody() *ListQualityWatchesResponseBody {
	return s.Body
}

func (s *ListQualityWatchesResponse) SetHeaders(v map[string]*string) *ListQualityWatchesResponse {
	s.Headers = v
	return s
}

func (s *ListQualityWatchesResponse) SetStatusCode(v int32) *ListQualityWatchesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQualityWatchesResponse) SetBody(v *ListQualityWatchesResponseBody) *ListQualityWatchesResponse {
	s.Body = v
	return s
}

func (s *ListQualityWatchesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
