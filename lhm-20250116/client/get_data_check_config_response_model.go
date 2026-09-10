// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataCheckConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataCheckConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetDataCheckConfigResponseBody) *GetDataCheckConfigResponse
	GetBody() *GetDataCheckConfigResponseBody
}

type GetDataCheckConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataCheckConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataCheckConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckConfigResponse) GoString() string {
	return s.String()
}

func (s *GetDataCheckConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataCheckConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataCheckConfigResponse) GetBody() *GetDataCheckConfigResponseBody {
	return s.Body
}

func (s *GetDataCheckConfigResponse) SetHeaders(v map[string]*string) *GetDataCheckConfigResponse {
	s.Headers = v
	return s
}

func (s *GetDataCheckConfigResponse) SetStatusCode(v int32) *GetDataCheckConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataCheckConfigResponse) SetBody(v *GetDataCheckConfigResponseBody) *GetDataCheckConfigResponse {
	s.Body = v
	return s
}

func (s *GetDataCheckConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
