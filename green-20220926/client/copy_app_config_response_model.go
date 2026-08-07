// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyAppConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyAppConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyAppConfigResponse
	GetStatusCode() *int32
	SetBody(v *CopyAppConfigResponseBody) *CopyAppConfigResponse
	GetBody() *CopyAppConfigResponseBody
}

type CopyAppConfigResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyAppConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyAppConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyAppConfigResponse) GoString() string {
	return s.String()
}

func (s *CopyAppConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyAppConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyAppConfigResponse) GetBody() *CopyAppConfigResponseBody {
	return s.Body
}

func (s *CopyAppConfigResponse) SetHeaders(v map[string]*string) *CopyAppConfigResponse {
	s.Headers = v
	return s
}

func (s *CopyAppConfigResponse) SetStatusCode(v int32) *CopyAppConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyAppConfigResponse) SetBody(v *CopyAppConfigResponseBody) *CopyAppConfigResponse {
	s.Body = v
	return s
}

func (s *CopyAppConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
