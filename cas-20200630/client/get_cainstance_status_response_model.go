// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCAInstanceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCAInstanceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCAInstanceStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetCAInstanceStatusResponseBody) *GetCAInstanceStatusResponse
	GetBody() *GetCAInstanceStatusResponseBody
}

type GetCAInstanceStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCAInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCAInstanceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCAInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetCAInstanceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCAInstanceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCAInstanceStatusResponse) GetBody() *GetCAInstanceStatusResponseBody {
	return s.Body
}

func (s *GetCAInstanceStatusResponse) SetHeaders(v map[string]*string) *GetCAInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetCAInstanceStatusResponse) SetStatusCode(v int32) *GetCAInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCAInstanceStatusResponse) SetBody(v *GetCAInstanceStatusResponseBody) *GetCAInstanceStatusResponse {
	s.Body = v
	return s
}

func (s *GetCAInstanceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
