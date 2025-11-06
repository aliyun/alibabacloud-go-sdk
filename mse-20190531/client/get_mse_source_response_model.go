// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMseSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMseSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMseSourceResponse
	GetStatusCode() *int32
	SetBody(v *GetMseSourceResponseBody) *GetMseSourceResponse
	GetBody() *GetMseSourceResponseBody
}

type GetMseSourceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMseSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMseSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMseSourceResponse) GoString() string {
	return s.String()
}

func (s *GetMseSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMseSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMseSourceResponse) GetBody() *GetMseSourceResponseBody {
	return s.Body
}

func (s *GetMseSourceResponse) SetHeaders(v map[string]*string) *GetMseSourceResponse {
	s.Headers = v
	return s
}

func (s *GetMseSourceResponse) SetStatusCode(v int32) *GetMseSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMseSourceResponse) SetBody(v *GetMseSourceResponseBody) *GetMseSourceResponse {
	s.Body = v
	return s
}

func (s *GetMseSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
