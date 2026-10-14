// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAifgUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAifgUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAifgUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetAifgUrlResponseBody) *GetAifgUrlResponse
	GetBody() *GetAifgUrlResponseBody
}

type GetAifgUrlResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAifgUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAifgUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAifgUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAifgUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAifgUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAifgUrlResponse) GetBody() *GetAifgUrlResponseBody {
	return s.Body
}

func (s *GetAifgUrlResponse) SetHeaders(v map[string]*string) *GetAifgUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAifgUrlResponse) SetStatusCode(v int32) *GetAifgUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAifgUrlResponse) SetBody(v *GetAifgUrlResponseBody) *GetAifgUrlResponse {
	s.Body = v
	return s
}

func (s *GetAifgUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
