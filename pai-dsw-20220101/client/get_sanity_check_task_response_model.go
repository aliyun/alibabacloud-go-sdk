// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSanityCheckTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSanityCheckTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSanityCheckTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetSanityCheckTaskResponseBody) *GetSanityCheckTaskResponse
	GetBody() *GetSanityCheckTaskResponseBody
}

type GetSanityCheckTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSanityCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSanityCheckTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSanityCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *GetSanityCheckTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSanityCheckTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSanityCheckTaskResponse) GetBody() *GetSanityCheckTaskResponseBody {
	return s.Body
}

func (s *GetSanityCheckTaskResponse) SetHeaders(v map[string]*string) *GetSanityCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *GetSanityCheckTaskResponse) SetStatusCode(v int32) *GetSanityCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSanityCheckTaskResponse) SetBody(v *GetSanityCheckTaskResponseBody) *GetSanityCheckTaskResponse {
	s.Body = v
	return s
}

func (s *GetSanityCheckTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
