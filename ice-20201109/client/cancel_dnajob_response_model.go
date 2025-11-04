// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDNAJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelDNAJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelDNAJobResponse
	GetStatusCode() *int32
	SetBody(v *CancelDNAJobResponseBody) *CancelDNAJobResponse
	GetBody() *CancelDNAJobResponseBody
}

type CancelDNAJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelDNAJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelDNAJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelDNAJobResponse) GoString() string {
	return s.String()
}

func (s *CancelDNAJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelDNAJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelDNAJobResponse) GetBody() *CancelDNAJobResponseBody {
	return s.Body
}

func (s *CancelDNAJobResponse) SetHeaders(v map[string]*string) *CancelDNAJobResponse {
	s.Headers = v
	return s
}

func (s *CancelDNAJobResponse) SetStatusCode(v int32) *CancelDNAJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelDNAJobResponse) SetBody(v *CancelDNAJobResponseBody) *CancelDNAJobResponse {
	s.Body = v
	return s
}

func (s *CancelDNAJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
