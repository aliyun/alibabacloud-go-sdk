// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJinaAiReaderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJinaAiReaderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJinaAiReaderResponse
	GetStatusCode() *int32
	SetBody(v *GetJinaAiReaderResponseBody) *GetJinaAiReaderResponse
	GetBody() *GetJinaAiReaderResponseBody
}

type GetJinaAiReaderResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJinaAiReaderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJinaAiReaderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJinaAiReaderResponse) GoString() string {
	return s.String()
}

func (s *GetJinaAiReaderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJinaAiReaderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJinaAiReaderResponse) GetBody() *GetJinaAiReaderResponseBody {
	return s.Body
}

func (s *GetJinaAiReaderResponse) SetHeaders(v map[string]*string) *GetJinaAiReaderResponse {
	s.Headers = v
	return s
}

func (s *GetJinaAiReaderResponse) SetStatusCode(v int32) *GetJinaAiReaderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJinaAiReaderResponse) SetBody(v *GetJinaAiReaderResponseBody) *GetJinaAiReaderResponse {
	s.Body = v
	return s
}

func (s *GetJinaAiReaderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
