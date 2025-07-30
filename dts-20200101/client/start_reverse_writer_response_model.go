// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartReverseWriterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartReverseWriterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartReverseWriterResponse
	GetStatusCode() *int32
	SetBody(v *StartReverseWriterResponseBody) *StartReverseWriterResponse
	GetBody() *StartReverseWriterResponseBody
}

type StartReverseWriterResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartReverseWriterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartReverseWriterResponse) String() string {
	return dara.Prettify(s)
}

func (s StartReverseWriterResponse) GoString() string {
	return s.String()
}

func (s *StartReverseWriterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartReverseWriterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartReverseWriterResponse) GetBody() *StartReverseWriterResponseBody {
	return s.Body
}

func (s *StartReverseWriterResponse) SetHeaders(v map[string]*string) *StartReverseWriterResponse {
	s.Headers = v
	return s
}

func (s *StartReverseWriterResponse) SetStatusCode(v int32) *StartReverseWriterResponse {
	s.StatusCode = &v
	return s
}

func (s *StartReverseWriterResponse) SetBody(v *StartReverseWriterResponseBody) *StartReverseWriterResponse {
	s.Body = v
	return s
}

func (s *StartReverseWriterResponse) Validate() error {
	return dara.Validate(s)
}
