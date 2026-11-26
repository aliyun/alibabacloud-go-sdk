// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUndeployHttpApiRoutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchUndeployHttpApiRoutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchUndeployHttpApiRoutesResponse
	GetStatusCode() *int32
	SetBody(v *BatchUndeployHttpApiRoutesResponseBody) *BatchUndeployHttpApiRoutesResponse
	GetBody() *BatchUndeployHttpApiRoutesResponseBody
}

type BatchUndeployHttpApiRoutesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUndeployHttpApiRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUndeployHttpApiRoutesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchUndeployHttpApiRoutesResponse) GoString() string {
	return s.String()
}

func (s *BatchUndeployHttpApiRoutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchUndeployHttpApiRoutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchUndeployHttpApiRoutesResponse) GetBody() *BatchUndeployHttpApiRoutesResponseBody {
	return s.Body
}

func (s *BatchUndeployHttpApiRoutesResponse) SetHeaders(v map[string]*string) *BatchUndeployHttpApiRoutesResponse {
	s.Headers = v
	return s
}

func (s *BatchUndeployHttpApiRoutesResponse) SetStatusCode(v int32) *BatchUndeployHttpApiRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUndeployHttpApiRoutesResponse) SetBody(v *BatchUndeployHttpApiRoutesResponseBody) *BatchUndeployHttpApiRoutesResponse {
	s.Body = v
	return s
}

func (s *BatchUndeployHttpApiRoutesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
