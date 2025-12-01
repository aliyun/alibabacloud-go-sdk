// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuspEventPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSuspEventPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSuspEventPageResponse
	GetStatusCode() *int32
	SetBody(v *GetSuspEventPageResponseBody) *GetSuspEventPageResponse
	GetBody() *GetSuspEventPageResponseBody
}

type GetSuspEventPageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSuspEventPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSuspEventPageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSuspEventPageResponse) GoString() string {
	return s.String()
}

func (s *GetSuspEventPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSuspEventPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSuspEventPageResponse) GetBody() *GetSuspEventPageResponseBody {
	return s.Body
}

func (s *GetSuspEventPageResponse) SetHeaders(v map[string]*string) *GetSuspEventPageResponse {
	s.Headers = v
	return s
}

func (s *GetSuspEventPageResponse) SetStatusCode(v int32) *GetSuspEventPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSuspEventPageResponse) SetBody(v *GetSuspEventPageResponseBody) *GetSuspEventPageResponse {
	s.Body = v
	return s
}

func (s *GetSuspEventPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
