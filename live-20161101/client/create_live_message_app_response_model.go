// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveMessageAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLiveMessageAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLiveMessageAppResponse
	GetStatusCode() *int32
	SetBody(v *CreateLiveMessageAppResponseBody) *CreateLiveMessageAppResponse
	GetBody() *CreateLiveMessageAppResponseBody
}

type CreateLiveMessageAppResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLiveMessageAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLiveMessageAppResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveMessageAppResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveMessageAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLiveMessageAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLiveMessageAppResponse) GetBody() *CreateLiveMessageAppResponseBody {
	return s.Body
}

func (s *CreateLiveMessageAppResponse) SetHeaders(v map[string]*string) *CreateLiveMessageAppResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveMessageAppResponse) SetStatusCode(v int32) *CreateLiveMessageAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLiveMessageAppResponse) SetBody(v *CreateLiveMessageAppResponseBody) *CreateLiveMessageAppResponse {
	s.Body = v
	return s
}

func (s *CreateLiveMessageAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
