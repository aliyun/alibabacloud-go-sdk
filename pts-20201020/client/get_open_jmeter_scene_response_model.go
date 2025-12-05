// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpenJMeterSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOpenJMeterSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOpenJMeterSceneResponse
	GetStatusCode() *int32
	SetBody(v *GetOpenJMeterSceneResponseBody) *GetOpenJMeterSceneResponse
	GetBody() *GetOpenJMeterSceneResponseBody
}

type GetOpenJMeterSceneResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpenJMeterSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpenJMeterSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOpenJMeterSceneResponse) GoString() string {
	return s.String()
}

func (s *GetOpenJMeterSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOpenJMeterSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOpenJMeterSceneResponse) GetBody() *GetOpenJMeterSceneResponseBody {
	return s.Body
}

func (s *GetOpenJMeterSceneResponse) SetHeaders(v map[string]*string) *GetOpenJMeterSceneResponse {
	s.Headers = v
	return s
}

func (s *GetOpenJMeterSceneResponse) SetStatusCode(v int32) *GetOpenJMeterSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpenJMeterSceneResponse) SetBody(v *GetOpenJMeterSceneResponseBody) *GetOpenJMeterSceneResponse {
	s.Body = v
	return s
}

func (s *GetOpenJMeterSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
