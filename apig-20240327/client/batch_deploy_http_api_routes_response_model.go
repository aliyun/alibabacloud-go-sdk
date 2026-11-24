// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeployHttpApiRoutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeployHttpApiRoutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeployHttpApiRoutesResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeployHttpApiRoutesResponseBody) *BatchDeployHttpApiRoutesResponse
	GetBody() *BatchDeployHttpApiRoutesResponseBody
}

type BatchDeployHttpApiRoutesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeployHttpApiRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeployHttpApiRoutesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeployHttpApiRoutesResponse) GoString() string {
	return s.String()
}

func (s *BatchDeployHttpApiRoutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeployHttpApiRoutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeployHttpApiRoutesResponse) GetBody() *BatchDeployHttpApiRoutesResponseBody {
	return s.Body
}

func (s *BatchDeployHttpApiRoutesResponse) SetHeaders(v map[string]*string) *BatchDeployHttpApiRoutesResponse {
	s.Headers = v
	return s
}

func (s *BatchDeployHttpApiRoutesResponse) SetStatusCode(v int32) *BatchDeployHttpApiRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeployHttpApiRoutesResponse) SetBody(v *BatchDeployHttpApiRoutesResponseBody) *BatchDeployHttpApiRoutesResponse {
	s.Body = v
	return s
}

func (s *BatchDeployHttpApiRoutesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
