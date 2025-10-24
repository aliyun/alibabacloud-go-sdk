// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelJobResponse
	GetStatusCode() *int32
	SetBody(v *CancelJobResponseBody) *CancelJobResponse
	GetBody() *CancelJobResponseBody
}

type CancelJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelJobResponse) GoString() string {
	return s.String()
}

func (s *CancelJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelJobResponse) GetBody() *CancelJobResponseBody {
	return s.Body
}

func (s *CancelJobResponse) SetHeaders(v map[string]*string) *CancelJobResponse {
	s.Headers = v
	return s
}

func (s *CancelJobResponse) SetStatusCode(v int32) *CancelJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelJobResponse) SetBody(v *CancelJobResponseBody) *CancelJobResponse {
	s.Body = v
	return s
}

func (s *CancelJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
