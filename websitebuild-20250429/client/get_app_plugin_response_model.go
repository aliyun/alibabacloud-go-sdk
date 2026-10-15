// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppPluginResponse
	GetStatusCode() *int32
	SetBody(v *GetAppPluginResponseBody) *GetAppPluginResponse
	GetBody() *GetAppPluginResponseBody
}

type GetAppPluginResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppPluginResponse) GoString() string {
	return s.String()
}

func (s *GetAppPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppPluginResponse) GetBody() *GetAppPluginResponseBody {
	return s.Body
}

func (s *GetAppPluginResponse) SetHeaders(v map[string]*string) *GetAppPluginResponse {
	s.Headers = v
	return s
}

func (s *GetAppPluginResponse) SetStatusCode(v int32) *GetAppPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppPluginResponse) SetBody(v *GetAppPluginResponseBody) *GetAppPluginResponse {
	s.Body = v
	return s
}

func (s *GetAppPluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
