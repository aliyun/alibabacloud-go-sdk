// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBizMetricByNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBizMetricByNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBizMetricByNameResponse
	GetStatusCode() *int32
	SetBody(v *GetBizMetricByNameResponseBody) *GetBizMetricByNameResponse
	GetBody() *GetBizMetricByNameResponseBody
}

type GetBizMetricByNameResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBizMetricByNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBizMetricByNameResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBizMetricByNameResponse) GoString() string {
	return s.String()
}

func (s *GetBizMetricByNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBizMetricByNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBizMetricByNameResponse) GetBody() *GetBizMetricByNameResponseBody {
	return s.Body
}

func (s *GetBizMetricByNameResponse) SetHeaders(v map[string]*string) *GetBizMetricByNameResponse {
	s.Headers = v
	return s
}

func (s *GetBizMetricByNameResponse) SetStatusCode(v int32) *GetBizMetricByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBizMetricByNameResponse) SetBody(v *GetBizMetricByNameResponseBody) *GetBizMetricByNameResponse {
	s.Body = v
	return s
}

func (s *GetBizMetricByNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
