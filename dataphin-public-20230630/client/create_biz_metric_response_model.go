// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBizMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBizMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBizMetricResponse
	GetStatusCode() *int32
	SetBody(v *CreateBizMetricResponseBody) *CreateBizMetricResponse
	GetBody() *CreateBizMetricResponseBody
}

type CreateBizMetricResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBizMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBizMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBizMetricResponse) GoString() string {
	return s.String()
}

func (s *CreateBizMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBizMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBizMetricResponse) GetBody() *CreateBizMetricResponseBody {
	return s.Body
}

func (s *CreateBizMetricResponse) SetHeaders(v map[string]*string) *CreateBizMetricResponse {
	s.Headers = v
	return s
}

func (s *CreateBizMetricResponse) SetStatusCode(v int32) *CreateBizMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBizMetricResponse) SetBody(v *CreateBizMetricResponseBody) *CreateBizMetricResponse {
	s.Body = v
	return s
}

func (s *CreateBizMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
