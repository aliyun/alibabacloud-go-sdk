// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileDetectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFileDetectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFileDetectResponse
	GetStatusCode() *int32
	SetBody(v *CreateFileDetectResponseBody) *CreateFileDetectResponse
	GetBody() *CreateFileDetectResponseBody
}

type CreateFileDetectResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFileDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFileDetectResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFileDetectResponse) GoString() string {
	return s.String()
}

func (s *CreateFileDetectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFileDetectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFileDetectResponse) GetBody() *CreateFileDetectResponseBody {
	return s.Body
}

func (s *CreateFileDetectResponse) SetHeaders(v map[string]*string) *CreateFileDetectResponse {
	s.Headers = v
	return s
}

func (s *CreateFileDetectResponse) SetStatusCode(v int32) *CreateFileDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileDetectResponse) SetBody(v *CreateFileDetectResponseBody) *CreateFileDetectResponse {
	s.Body = v
	return s
}

func (s *CreateFileDetectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
