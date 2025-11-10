// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumeTimespanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConsumeTimespanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConsumeTimespanResponse
	GetStatusCode() *int32
	SetBody(v *GetConsumeTimespanResponseBody) *GetConsumeTimespanResponse
	GetBody() *GetConsumeTimespanResponseBody
}

type GetConsumeTimespanResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsumeTimespanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsumeTimespanResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConsumeTimespanResponse) GoString() string {
	return s.String()
}

func (s *GetConsumeTimespanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConsumeTimespanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConsumeTimespanResponse) GetBody() *GetConsumeTimespanResponseBody {
	return s.Body
}

func (s *GetConsumeTimespanResponse) SetHeaders(v map[string]*string) *GetConsumeTimespanResponse {
	s.Headers = v
	return s
}

func (s *GetConsumeTimespanResponse) SetStatusCode(v int32) *GetConsumeTimespanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsumeTimespanResponse) SetBody(v *GetConsumeTimespanResponseBody) *GetConsumeTimespanResponse {
	s.Body = v
	return s
}

func (s *GetConsumeTimespanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
