// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotfixTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotfixTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotfixTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetHotfixTaskResponseBody) *GetHotfixTaskResponse
	GetBody() *GetHotfixTaskResponseBody
}

type GetHotfixTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotfixTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotfixTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotfixTaskResponse) GoString() string {
	return s.String()
}

func (s *GetHotfixTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotfixTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotfixTaskResponse) GetBody() *GetHotfixTaskResponseBody {
	return s.Body
}

func (s *GetHotfixTaskResponse) SetHeaders(v map[string]*string) *GetHotfixTaskResponse {
	s.Headers = v
	return s
}

func (s *GetHotfixTaskResponse) SetStatusCode(v int32) *GetHotfixTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotfixTaskResponse) SetBody(v *GetHotfixTaskResponseBody) *GetHotfixTaskResponse {
	s.Body = v
	return s
}

func (s *GetHotfixTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
