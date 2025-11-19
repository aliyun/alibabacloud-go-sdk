// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMezzanineInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMezzanineInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMezzanineInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetMezzanineInfoResponseBody) *GetMezzanineInfoResponse
	GetBody() *GetMezzanineInfoResponseBody
}

type GetMezzanineInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMezzanineInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMezzanineInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMezzanineInfoResponse) GoString() string {
	return s.String()
}

func (s *GetMezzanineInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMezzanineInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMezzanineInfoResponse) GetBody() *GetMezzanineInfoResponseBody {
	return s.Body
}

func (s *GetMezzanineInfoResponse) SetHeaders(v map[string]*string) *GetMezzanineInfoResponse {
	s.Headers = v
	return s
}

func (s *GetMezzanineInfoResponse) SetStatusCode(v int32) *GetMezzanineInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMezzanineInfoResponse) SetBody(v *GetMezzanineInfoResponseBody) *GetMezzanineInfoResponse {
	s.Body = v
	return s
}

func (s *GetMezzanineInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
