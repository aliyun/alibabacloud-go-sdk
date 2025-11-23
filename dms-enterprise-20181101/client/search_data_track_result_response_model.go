// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDataTrackResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchDataTrackResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchDataTrackResultResponse
	GetStatusCode() *int32
	SetBody(v *SearchDataTrackResultResponseBody) *SearchDataTrackResultResponse
	GetBody() *SearchDataTrackResultResponseBody
}

type SearchDataTrackResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchDataTrackResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchDataTrackResultResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchDataTrackResultResponse) GoString() string {
	return s.String()
}

func (s *SearchDataTrackResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchDataTrackResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchDataTrackResultResponse) GetBody() *SearchDataTrackResultResponseBody {
	return s.Body
}

func (s *SearchDataTrackResultResponse) SetHeaders(v map[string]*string) *SearchDataTrackResultResponse {
	s.Headers = v
	return s
}

func (s *SearchDataTrackResultResponse) SetStatusCode(v int32) *SearchDataTrackResultResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchDataTrackResultResponse) SetBody(v *SearchDataTrackResultResponseBody) *SearchDataTrackResultResponse {
	s.Body = v
	return s
}

func (s *SearchDataTrackResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
