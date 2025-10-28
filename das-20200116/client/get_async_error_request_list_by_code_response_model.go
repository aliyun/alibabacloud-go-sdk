// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncErrorRequestListByCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAsyncErrorRequestListByCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAsyncErrorRequestListByCodeResponse
	GetStatusCode() *int32
	SetBody(v *GetAsyncErrorRequestListByCodeResponseBody) *GetAsyncErrorRequestListByCodeResponse
	GetBody() *GetAsyncErrorRequestListByCodeResponseBody
}

type GetAsyncErrorRequestListByCodeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsyncErrorRequestListByCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsyncErrorRequestListByCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncErrorRequestListByCodeResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncErrorRequestListByCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAsyncErrorRequestListByCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAsyncErrorRequestListByCodeResponse) GetBody() *GetAsyncErrorRequestListByCodeResponseBody {
	return s.Body
}

func (s *GetAsyncErrorRequestListByCodeResponse) SetHeaders(v map[string]*string) *GetAsyncErrorRequestListByCodeResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponse) SetStatusCode(v int32) *GetAsyncErrorRequestListByCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponse) SetBody(v *GetAsyncErrorRequestListByCodeResponseBody) *GetAsyncErrorRequestListByCodeResponse {
	s.Body = v
	return s
}

func (s *GetAsyncErrorRequestListByCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
