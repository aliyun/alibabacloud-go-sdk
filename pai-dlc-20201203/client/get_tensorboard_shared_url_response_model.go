// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTensorboardSharedUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTensorboardSharedUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTensorboardSharedUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetTensorboardSharedUrlResponseBody) *GetTensorboardSharedUrlResponse
	GetBody() *GetTensorboardSharedUrlResponseBody
}

type GetTensorboardSharedUrlResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTensorboardSharedUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTensorboardSharedUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTensorboardSharedUrlResponse) GoString() string {
	return s.String()
}

func (s *GetTensorboardSharedUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTensorboardSharedUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTensorboardSharedUrlResponse) GetBody() *GetTensorboardSharedUrlResponseBody {
	return s.Body
}

func (s *GetTensorboardSharedUrlResponse) SetHeaders(v map[string]*string) *GetTensorboardSharedUrlResponse {
	s.Headers = v
	return s
}

func (s *GetTensorboardSharedUrlResponse) SetStatusCode(v int32) *GetTensorboardSharedUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTensorboardSharedUrlResponse) SetBody(v *GetTensorboardSharedUrlResponseBody) *GetTensorboardSharedUrlResponse {
	s.Body = v
	return s
}

func (s *GetTensorboardSharedUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
