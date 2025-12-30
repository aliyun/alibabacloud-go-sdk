// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBizMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBizMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBizMetricResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBizMetricResponseBody) *DeleteBizMetricResponse
	GetBody() *DeleteBizMetricResponseBody
}

type DeleteBizMetricResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBizMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBizMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBizMetricResponse) GoString() string {
	return s.String()
}

func (s *DeleteBizMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBizMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBizMetricResponse) GetBody() *DeleteBizMetricResponseBody {
	return s.Body
}

func (s *DeleteBizMetricResponse) SetHeaders(v map[string]*string) *DeleteBizMetricResponse {
	s.Headers = v
	return s
}

func (s *DeleteBizMetricResponse) SetStatusCode(v int32) *DeleteBizMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBizMetricResponse) SetBody(v *DeleteBizMetricResponseBody) *DeleteBizMetricResponse {
	s.Body = v
	return s
}

func (s *DeleteBizMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
