// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaCensorJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitMediaCensorJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitMediaCensorJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitMediaCensorJobResponseBody) *SubmitMediaCensorJobResponse
	GetBody() *SubmitMediaCensorJobResponseBody
}

type SubmitMediaCensorJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitMediaCensorJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitMediaCensorJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaCensorJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitMediaCensorJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitMediaCensorJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitMediaCensorJobResponse) GetBody() *SubmitMediaCensorJobResponseBody {
	return s.Body
}

func (s *SubmitMediaCensorJobResponse) SetHeaders(v map[string]*string) *SubmitMediaCensorJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitMediaCensorJobResponse) SetStatusCode(v int32) *SubmitMediaCensorJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitMediaCensorJobResponse) SetBody(v *SubmitMediaCensorJobResponseBody) *SubmitMediaCensorJobResponse {
	s.Body = v
	return s
}

func (s *SubmitMediaCensorJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
