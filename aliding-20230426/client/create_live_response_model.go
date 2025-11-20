// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLiveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLiveResponse
	GetStatusCode() *int32
	SetBody(v *CreateLiveResponseBody) *CreateLiveResponse
	GetBody() *CreateLiveResponseBody
}

type CreateLiveResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLiveResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLiveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLiveResponse) GetBody() *CreateLiveResponseBody {
	return s.Body
}

func (s *CreateLiveResponse) SetHeaders(v map[string]*string) *CreateLiveResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveResponse) SetStatusCode(v int32) *CreateLiveResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLiveResponse) SetBody(v *CreateLiveResponseBody) *CreateLiveResponse {
	s.Body = v
	return s
}

func (s *CreateLiveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
