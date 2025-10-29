// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMessageAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMessageAppResponse
	GetStatusCode() *int32
	SetBody(v *GetMessageAppResponseBody) *GetMessageAppResponse
	GetBody() *GetMessageAppResponseBody
}

type GetMessageAppResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMessageAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMessageAppResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMessageAppResponse) GoString() string {
	return s.String()
}

func (s *GetMessageAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMessageAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMessageAppResponse) GetBody() *GetMessageAppResponseBody {
	return s.Body
}

func (s *GetMessageAppResponse) SetHeaders(v map[string]*string) *GetMessageAppResponse {
	s.Headers = v
	return s
}

func (s *GetMessageAppResponse) SetStatusCode(v int32) *GetMessageAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMessageAppResponse) SetBody(v *GetMessageAppResponseBody) *GetMessageAppResponse {
	s.Body = v
	return s
}

func (s *GetMessageAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
