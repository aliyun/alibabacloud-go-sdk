// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLivyComputeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartLivyComputeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartLivyComputeResponse
	GetStatusCode() *int32
	SetBody(v *StartLivyComputeResponseBody) *StartLivyComputeResponse
	GetBody() *StartLivyComputeResponseBody
}

type StartLivyComputeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartLivyComputeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartLivyComputeResponse) String() string {
	return dara.Prettify(s)
}

func (s StartLivyComputeResponse) GoString() string {
	return s.String()
}

func (s *StartLivyComputeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartLivyComputeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartLivyComputeResponse) GetBody() *StartLivyComputeResponseBody {
	return s.Body
}

func (s *StartLivyComputeResponse) SetHeaders(v map[string]*string) *StartLivyComputeResponse {
	s.Headers = v
	return s
}

func (s *StartLivyComputeResponse) SetStatusCode(v int32) *StartLivyComputeResponse {
	s.StatusCode = &v
	return s
}

func (s *StartLivyComputeResponse) SetBody(v *StartLivyComputeResponseBody) *StartLivyComputeResponse {
	s.Body = v
	return s
}

func (s *StartLivyComputeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
