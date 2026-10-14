// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateC2paSignTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateC2paSignTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateC2paSignTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateC2paSignTaskResponseBody) *CreateC2paSignTaskResponse
	GetBody() *CreateC2paSignTaskResponseBody
}

type CreateC2paSignTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateC2paSignTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateC2paSignTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateC2paSignTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateC2paSignTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateC2paSignTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateC2paSignTaskResponse) GetBody() *CreateC2paSignTaskResponseBody {
	return s.Body
}

func (s *CreateC2paSignTaskResponse) SetHeaders(v map[string]*string) *CreateC2paSignTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateC2paSignTaskResponse) SetStatusCode(v int32) *CreateC2paSignTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateC2paSignTaskResponse) SetBody(v *CreateC2paSignTaskResponseBody) *CreateC2paSignTaskResponse {
	s.Body = v
	return s
}

func (s *CreateC2paSignTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
