// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleComplaintResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HandleComplaintResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HandleComplaintResponse
	GetStatusCode() *int32
	SetBody(v *HandleComplaintResponseBody) *HandleComplaintResponse
	GetBody() *HandleComplaintResponseBody
}

type HandleComplaintResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HandleComplaintResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HandleComplaintResponse) String() string {
	return dara.Prettify(s)
}

func (s HandleComplaintResponse) GoString() string {
	return s.String()
}

func (s *HandleComplaintResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HandleComplaintResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HandleComplaintResponse) GetBody() *HandleComplaintResponseBody {
	return s.Body
}

func (s *HandleComplaintResponse) SetHeaders(v map[string]*string) *HandleComplaintResponse {
	s.Headers = v
	return s
}

func (s *HandleComplaintResponse) SetStatusCode(v int32) *HandleComplaintResponse {
	s.StatusCode = &v
	return s
}

func (s *HandleComplaintResponse) SetBody(v *HandleComplaintResponseBody) *HandleComplaintResponse {
	s.Body = v
	return s
}

func (s *HandleComplaintResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
