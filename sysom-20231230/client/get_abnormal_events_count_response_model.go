// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAbnormalEventsCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAbnormalEventsCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAbnormalEventsCountResponse
	GetStatusCode() *int32
	SetBody(v *GetAbnormalEventsCountResponseBody) *GetAbnormalEventsCountResponse
	GetBody() *GetAbnormalEventsCountResponseBody
}

type GetAbnormalEventsCountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAbnormalEventsCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAbnormalEventsCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAbnormalEventsCountResponse) GoString() string {
	return s.String()
}

func (s *GetAbnormalEventsCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAbnormalEventsCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAbnormalEventsCountResponse) GetBody() *GetAbnormalEventsCountResponseBody {
	return s.Body
}

func (s *GetAbnormalEventsCountResponse) SetHeaders(v map[string]*string) *GetAbnormalEventsCountResponse {
	s.Headers = v
	return s
}

func (s *GetAbnormalEventsCountResponse) SetStatusCode(v int32) *GetAbnormalEventsCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAbnormalEventsCountResponse) SetBody(v *GetAbnormalEventsCountResponseBody) *GetAbnormalEventsCountResponse {
	s.Body = v
	return s
}

func (s *GetAbnormalEventsCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
