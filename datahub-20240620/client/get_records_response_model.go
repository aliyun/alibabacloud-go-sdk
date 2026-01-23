// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRecordsResponse
	GetStatusCode() *int32
	SetBody(v *GetRecordsResponseBody) *GetRecordsResponse
	GetBody() *GetRecordsResponseBody
}

type GetRecordsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRecordsResponse) GoString() string {
	return s.String()
}

func (s *GetRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRecordsResponse) GetBody() *GetRecordsResponseBody {
	return s.Body
}

func (s *GetRecordsResponse) SetHeaders(v map[string]*string) *GetRecordsResponse {
	s.Headers = v
	return s
}

func (s *GetRecordsResponse) SetStatusCode(v int32) *GetRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecordsResponse) SetBody(v *GetRecordsResponseBody) *GetRecordsResponse {
	s.Body = v
	return s
}

func (s *GetRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
