// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCopyrightJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitCopyrightJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitCopyrightJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitCopyrightJobResponseBody) *SubmitCopyrightJobResponse
	GetBody() *SubmitCopyrightJobResponseBody
}

type SubmitCopyrightJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitCopyrightJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitCopyrightJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitCopyrightJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitCopyrightJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitCopyrightJobResponse) GetBody() *SubmitCopyrightJobResponseBody {
	return s.Body
}

func (s *SubmitCopyrightJobResponse) SetHeaders(v map[string]*string) *SubmitCopyrightJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitCopyrightJobResponse) SetStatusCode(v int32) *SubmitCopyrightJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitCopyrightJobResponse) SetBody(v *SubmitCopyrightJobResponseBody) *SubmitCopyrightJobResponse {
	s.Body = v
	return s
}

func (s *SubmitCopyrightJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
