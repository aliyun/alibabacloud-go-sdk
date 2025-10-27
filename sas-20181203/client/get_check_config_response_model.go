// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCheckConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCheckConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetCheckConfigResponseBody) *GetCheckConfigResponse
	GetBody() *GetCheckConfigResponseBody
}

type GetCheckConfigResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCheckConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCheckConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCheckConfigResponse) GoString() string {
	return s.String()
}

func (s *GetCheckConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCheckConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCheckConfigResponse) GetBody() *GetCheckConfigResponseBody {
	return s.Body
}

func (s *GetCheckConfigResponse) SetHeaders(v map[string]*string) *GetCheckConfigResponse {
	s.Headers = v
	return s
}

func (s *GetCheckConfigResponse) SetStatusCode(v int32) *GetCheckConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckConfigResponse) SetBody(v *GetCheckConfigResponseBody) *GetCheckConfigResponse {
	s.Body = v
	return s
}

func (s *GetCheckConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
