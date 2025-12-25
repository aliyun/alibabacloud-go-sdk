// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetThreadDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetThreadDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetThreadDataResponse
	GetStatusCode() *int32
	SetBody(v *GetThreadDataResponseBody) *GetThreadDataResponse
	GetBody() *GetThreadDataResponseBody
}

type GetThreadDataResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetThreadDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetThreadDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetThreadDataResponse) GoString() string {
	return s.String()
}

func (s *GetThreadDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetThreadDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetThreadDataResponse) GetBody() *GetThreadDataResponseBody {
	return s.Body
}

func (s *GetThreadDataResponse) SetHeaders(v map[string]*string) *GetThreadDataResponse {
	s.Headers = v
	return s
}

func (s *GetThreadDataResponse) SetStatusCode(v int32) *GetThreadDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetThreadDataResponse) SetBody(v *GetThreadDataResponseBody) *GetThreadDataResponse {
	s.Body = v
	return s
}

func (s *GetThreadDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
