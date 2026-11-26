// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpEventConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHttpEventConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHttpEventConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetHttpEventConfigResponseBody) *GetHttpEventConfigResponse
	GetBody() *GetHttpEventConfigResponseBody
}

type GetHttpEventConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHttpEventConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHttpEventConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHttpEventConfigResponse) GoString() string {
	return s.String()
}

func (s *GetHttpEventConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHttpEventConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHttpEventConfigResponse) GetBody() *GetHttpEventConfigResponseBody {
	return s.Body
}

func (s *GetHttpEventConfigResponse) SetHeaders(v map[string]*string) *GetHttpEventConfigResponse {
	s.Headers = v
	return s
}

func (s *GetHttpEventConfigResponse) SetStatusCode(v int32) *GetHttpEventConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHttpEventConfigResponse) SetBody(v *GetHttpEventConfigResponseBody) *GetHttpEventConfigResponse {
	s.Body = v
	return s
}

func (s *GetHttpEventConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
