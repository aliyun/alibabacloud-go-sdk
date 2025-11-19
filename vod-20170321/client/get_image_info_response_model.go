// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetImageInfoResponseBody) *GetImageInfoResponse
	GetBody() *GetImageInfoResponseBody
}

type GetImageInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageInfoResponse) GoString() string {
	return s.String()
}

func (s *GetImageInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageInfoResponse) GetBody() *GetImageInfoResponseBody {
	return s.Body
}

func (s *GetImageInfoResponse) SetHeaders(v map[string]*string) *GetImageInfoResponse {
	s.Headers = v
	return s
}

func (s *GetImageInfoResponse) SetStatusCode(v int32) *GetImageInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageInfoResponse) SetBody(v *GetImageInfoResponseBody) *GetImageInfoResponse {
	s.Body = v
	return s
}

func (s *GetImageInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
