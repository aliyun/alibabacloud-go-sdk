// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOpLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOpLogResponse
	GetStatusCode() *int32
	SetBody(v *GetOpLogResponseBody) *GetOpLogResponse
	GetBody() *GetOpLogResponseBody
}

type GetOpLogResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOpLogResponse) GoString() string {
	return s.String()
}

func (s *GetOpLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOpLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOpLogResponse) GetBody() *GetOpLogResponseBody {
	return s.Body
}

func (s *GetOpLogResponse) SetHeaders(v map[string]*string) *GetOpLogResponse {
	s.Headers = v
	return s
}

func (s *GetOpLogResponse) SetStatusCode(v int32) *GetOpLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpLogResponse) SetBody(v *GetOpLogResponseBody) *GetOpLogResponse {
	s.Body = v
	return s
}

func (s *GetOpLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
