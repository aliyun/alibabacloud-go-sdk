// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFormationCrawlerScheduleStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFormationCrawlerScheduleStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFormationCrawlerScheduleStateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFormationCrawlerScheduleStateResponseBody) *UpdateFormationCrawlerScheduleStateResponse
	GetBody() *UpdateFormationCrawlerScheduleStateResponseBody
}

type UpdateFormationCrawlerScheduleStateResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFormationCrawlerScheduleStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFormationCrawlerScheduleStateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFormationCrawlerScheduleStateResponse) GoString() string {
	return s.String()
}

func (s *UpdateFormationCrawlerScheduleStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFormationCrawlerScheduleStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFormationCrawlerScheduleStateResponse) GetBody() *UpdateFormationCrawlerScheduleStateResponseBody {
	return s.Body
}

func (s *UpdateFormationCrawlerScheduleStateResponse) SetHeaders(v map[string]*string) *UpdateFormationCrawlerScheduleStateResponse {
	s.Headers = v
	return s
}

func (s *UpdateFormationCrawlerScheduleStateResponse) SetStatusCode(v int32) *UpdateFormationCrawlerScheduleStateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFormationCrawlerScheduleStateResponse) SetBody(v *UpdateFormationCrawlerScheduleStateResponseBody) *UpdateFormationCrawlerScheduleStateResponse {
	s.Body = v
	return s
}

func (s *UpdateFormationCrawlerScheduleStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
