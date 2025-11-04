// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaConvertJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitMediaConvertJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitMediaConvertJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitMediaConvertJobResponseBody) *SubmitMediaConvertJobResponse
	GetBody() *SubmitMediaConvertJobResponseBody
}

type SubmitMediaConvertJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitMediaConvertJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitMediaConvertJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaConvertJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitMediaConvertJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitMediaConvertJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitMediaConvertJobResponse) GetBody() *SubmitMediaConvertJobResponseBody {
	return s.Body
}

func (s *SubmitMediaConvertJobResponse) SetHeaders(v map[string]*string) *SubmitMediaConvertJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitMediaConvertJobResponse) SetStatusCode(v int32) *SubmitMediaConvertJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitMediaConvertJobResponse) SetBody(v *SubmitMediaConvertJobResponseBody) *SubmitMediaConvertJobResponse {
	s.Body = v
	return s
}

func (s *SubmitMediaConvertJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
