// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGroupSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGroupSourceResponse
	GetStatusCode() *int32
	SetBody(v *GetGroupSourceResponseBody) *GetGroupSourceResponse
	GetBody() *GetGroupSourceResponseBody
}

type GetGroupSourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGroupSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGroupSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGroupSourceResponse) GoString() string {
	return s.String()
}

func (s *GetGroupSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGroupSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGroupSourceResponse) GetBody() *GetGroupSourceResponseBody {
	return s.Body
}

func (s *GetGroupSourceResponse) SetHeaders(v map[string]*string) *GetGroupSourceResponse {
	s.Headers = v
	return s
}

func (s *GetGroupSourceResponse) SetStatusCode(v int32) *GetGroupSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupSourceResponse) SetBody(v *GetGroupSourceResponseBody) *GetGroupSourceResponse {
	s.Body = v
	return s
}

func (s *GetGroupSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
