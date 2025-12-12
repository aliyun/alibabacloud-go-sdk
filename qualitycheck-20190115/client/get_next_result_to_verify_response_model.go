// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNextResultToVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNextResultToVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNextResultToVerifyResponse
	GetStatusCode() *int32
	SetBody(v *GetNextResultToVerifyResponseBody) *GetNextResultToVerifyResponse
	GetBody() *GetNextResultToVerifyResponseBody
}

type GetNextResultToVerifyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNextResultToVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNextResultToVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNextResultToVerifyResponse) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNextResultToVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNextResultToVerifyResponse) GetBody() *GetNextResultToVerifyResponseBody {
	return s.Body
}

func (s *GetNextResultToVerifyResponse) SetHeaders(v map[string]*string) *GetNextResultToVerifyResponse {
	s.Headers = v
	return s
}

func (s *GetNextResultToVerifyResponse) SetStatusCode(v int32) *GetNextResultToVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNextResultToVerifyResponse) SetBody(v *GetNextResultToVerifyResponseBody) *GetNextResultToVerifyResponse {
	s.Body = v
	return s
}

func (s *GetNextResultToVerifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
