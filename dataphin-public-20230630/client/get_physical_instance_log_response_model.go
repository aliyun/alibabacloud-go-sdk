// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalInstanceLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPhysicalInstanceLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPhysicalInstanceLogResponse
	GetStatusCode() *int32
	SetBody(v *GetPhysicalInstanceLogResponseBody) *GetPhysicalInstanceLogResponse
	GetBody() *GetPhysicalInstanceLogResponseBody
}

type GetPhysicalInstanceLogResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhysicalInstanceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhysicalInstanceLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalInstanceLogResponse) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPhysicalInstanceLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPhysicalInstanceLogResponse) GetBody() *GetPhysicalInstanceLogResponseBody {
	return s.Body
}

func (s *GetPhysicalInstanceLogResponse) SetHeaders(v map[string]*string) *GetPhysicalInstanceLogResponse {
	s.Headers = v
	return s
}

func (s *GetPhysicalInstanceLogResponse) SetStatusCode(v int32) *GetPhysicalInstanceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhysicalInstanceLogResponse) SetBody(v *GetPhysicalInstanceLogResponseBody) *GetPhysicalInstanceLogResponse {
	s.Body = v
	return s
}

func (s *GetPhysicalInstanceLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
