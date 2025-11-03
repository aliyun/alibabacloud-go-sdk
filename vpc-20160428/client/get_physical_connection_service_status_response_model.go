// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalConnectionServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPhysicalConnectionServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPhysicalConnectionServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetPhysicalConnectionServiceStatusResponseBody) *GetPhysicalConnectionServiceStatusResponse
	GetBody() *GetPhysicalConnectionServiceStatusResponseBody
}

type GetPhysicalConnectionServiceStatusResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhysicalConnectionServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhysicalConnectionServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalConnectionServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetPhysicalConnectionServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPhysicalConnectionServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPhysicalConnectionServiceStatusResponse) GetBody() *GetPhysicalConnectionServiceStatusResponseBody {
	return s.Body
}

func (s *GetPhysicalConnectionServiceStatusResponse) SetHeaders(v map[string]*string) *GetPhysicalConnectionServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetPhysicalConnectionServiceStatusResponse) SetStatusCode(v int32) *GetPhysicalConnectionServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhysicalConnectionServiceStatusResponse) SetBody(v *GetPhysicalConnectionServiceStatusResponseBody) *GetPhysicalConnectionServiceStatusResponse {
	s.Body = v
	return s
}

func (s *GetPhysicalConnectionServiceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
