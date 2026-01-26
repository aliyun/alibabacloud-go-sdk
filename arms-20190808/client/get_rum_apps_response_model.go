// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRumAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRumAppsResponse
	GetStatusCode() *int32
	SetBody(v *GetRumAppsResponseBody) *GetRumAppsResponse
	GetBody() *GetRumAppsResponseBody
}

type GetRumAppsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRumAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRumAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRumAppsResponse) GoString() string {
	return s.String()
}

func (s *GetRumAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRumAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRumAppsResponse) GetBody() *GetRumAppsResponseBody {
	return s.Body
}

func (s *GetRumAppsResponse) SetHeaders(v map[string]*string) *GetRumAppsResponse {
	s.Headers = v
	return s
}

func (s *GetRumAppsResponse) SetStatusCode(v int32) *GetRumAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRumAppsResponse) SetBody(v *GetRumAppsResponseBody) *GetRumAppsResponse {
	s.Body = v
	return s
}

func (s *GetRumAppsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
