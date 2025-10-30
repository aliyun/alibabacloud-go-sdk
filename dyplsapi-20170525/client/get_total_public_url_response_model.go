// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTotalPublicUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTotalPublicUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTotalPublicUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetTotalPublicUrlResponseBody) *GetTotalPublicUrlResponse
	GetBody() *GetTotalPublicUrlResponseBody
}

type GetTotalPublicUrlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTotalPublicUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTotalPublicUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTotalPublicUrlResponse) GoString() string {
	return s.String()
}

func (s *GetTotalPublicUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTotalPublicUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTotalPublicUrlResponse) GetBody() *GetTotalPublicUrlResponseBody {
	return s.Body
}

func (s *GetTotalPublicUrlResponse) SetHeaders(v map[string]*string) *GetTotalPublicUrlResponse {
	s.Headers = v
	return s
}

func (s *GetTotalPublicUrlResponse) SetStatusCode(v int32) *GetTotalPublicUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTotalPublicUrlResponse) SetBody(v *GetTotalPublicUrlResponseBody) *GetTotalPublicUrlResponse {
	s.Body = v
	return s
}

func (s *GetTotalPublicUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
