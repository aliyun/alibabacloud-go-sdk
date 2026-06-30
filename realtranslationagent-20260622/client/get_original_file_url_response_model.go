// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginalFileUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOriginalFileUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOriginalFileUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetOriginalFileUrlResponseBody) *GetOriginalFileUrlResponse
	GetBody() *GetOriginalFileUrlResponseBody
}

type GetOriginalFileUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOriginalFileUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOriginalFileUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOriginalFileUrlResponse) GoString() string {
	return s.String()
}

func (s *GetOriginalFileUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOriginalFileUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOriginalFileUrlResponse) GetBody() *GetOriginalFileUrlResponseBody {
	return s.Body
}

func (s *GetOriginalFileUrlResponse) SetHeaders(v map[string]*string) *GetOriginalFileUrlResponse {
	s.Headers = v
	return s
}

func (s *GetOriginalFileUrlResponse) SetStatusCode(v int32) *GetOriginalFileUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOriginalFileUrlResponse) SetBody(v *GetOriginalFileUrlResponseBody) *GetOriginalFileUrlResponse {
	s.Body = v
	return s
}

func (s *GetOriginalFileUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
