// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChartReleaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteChartReleaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteChartReleaseResponse
	GetStatusCode() *int32
	SetBody(v *DeleteChartReleaseResponseBody) *DeleteChartReleaseResponse
	GetBody() *DeleteChartReleaseResponseBody
}

type DeleteChartReleaseResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChartReleaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChartReleaseResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteChartReleaseResponse) GoString() string {
	return s.String()
}

func (s *DeleteChartReleaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteChartReleaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteChartReleaseResponse) GetBody() *DeleteChartReleaseResponseBody {
	return s.Body
}

func (s *DeleteChartReleaseResponse) SetHeaders(v map[string]*string) *DeleteChartReleaseResponse {
	s.Headers = v
	return s
}

func (s *DeleteChartReleaseResponse) SetStatusCode(v int32) *DeleteChartReleaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChartReleaseResponse) SetBody(v *DeleteChartReleaseResponseBody) *DeleteChartReleaseResponse {
	s.Body = v
	return s
}

func (s *DeleteChartReleaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
