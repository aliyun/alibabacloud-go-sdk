// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJMeterLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJMeterLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJMeterLogsResponse
	GetStatusCode() *int32
	SetBody(v *GetJMeterLogsResponseBody) *GetJMeterLogsResponse
	GetBody() *GetJMeterLogsResponseBody
}

type GetJMeterLogsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJMeterLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJMeterLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJMeterLogsResponse) GoString() string {
	return s.String()
}

func (s *GetJMeterLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJMeterLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJMeterLogsResponse) GetBody() *GetJMeterLogsResponseBody {
	return s.Body
}

func (s *GetJMeterLogsResponse) SetHeaders(v map[string]*string) *GetJMeterLogsResponse {
	s.Headers = v
	return s
}

func (s *GetJMeterLogsResponse) SetStatusCode(v int32) *GetJMeterLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJMeterLogsResponse) SetBody(v *GetJMeterLogsResponseBody) *GetJMeterLogsResponse {
	s.Body = v
	return s
}

func (s *GetJMeterLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
