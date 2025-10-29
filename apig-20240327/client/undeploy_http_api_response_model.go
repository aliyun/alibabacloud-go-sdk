// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUndeployHttpApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UndeployHttpApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UndeployHttpApiResponse
	GetStatusCode() *int32
	SetBody(v *UndeployHttpApiResponseBody) *UndeployHttpApiResponse
	GetBody() *UndeployHttpApiResponseBody
}

type UndeployHttpApiResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UndeployHttpApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UndeployHttpApiResponse) String() string {
	return dara.Prettify(s)
}

func (s UndeployHttpApiResponse) GoString() string {
	return s.String()
}

func (s *UndeployHttpApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UndeployHttpApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UndeployHttpApiResponse) GetBody() *UndeployHttpApiResponseBody {
	return s.Body
}

func (s *UndeployHttpApiResponse) SetHeaders(v map[string]*string) *UndeployHttpApiResponse {
	s.Headers = v
	return s
}

func (s *UndeployHttpApiResponse) SetStatusCode(v int32) *UndeployHttpApiResponse {
	s.StatusCode = &v
	return s
}

func (s *UndeployHttpApiResponse) SetBody(v *UndeployHttpApiResponseBody) *UndeployHttpApiResponse {
	s.Body = v
	return s
}

func (s *UndeployHttpApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
