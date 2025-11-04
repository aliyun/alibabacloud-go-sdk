// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamToSearchLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStreamToSearchLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStreamToSearchLibResponse
	GetStatusCode() *int32
	SetBody(v *CreateStreamToSearchLibResponseBody) *CreateStreamToSearchLibResponse
	GetBody() *CreateStreamToSearchLibResponseBody
}

type CreateStreamToSearchLibResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStreamToSearchLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStreamToSearchLibResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamToSearchLibResponse) GoString() string {
	return s.String()
}

func (s *CreateStreamToSearchLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStreamToSearchLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStreamToSearchLibResponse) GetBody() *CreateStreamToSearchLibResponseBody {
	return s.Body
}

func (s *CreateStreamToSearchLibResponse) SetHeaders(v map[string]*string) *CreateStreamToSearchLibResponse {
	s.Headers = v
	return s
}

func (s *CreateStreamToSearchLibResponse) SetStatusCode(v int32) *CreateStreamToSearchLibResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStreamToSearchLibResponse) SetBody(v *CreateStreamToSearchLibResponseBody) *CreateStreamToSearchLibResponse {
	s.Body = v
	return s
}

func (s *CreateStreamToSearchLibResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
