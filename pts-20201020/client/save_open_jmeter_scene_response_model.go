// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveOpenJMeterSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveOpenJMeterSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveOpenJMeterSceneResponse
	GetStatusCode() *int32
	SetBody(v *SaveOpenJMeterSceneResponseBody) *SaveOpenJMeterSceneResponse
	GetBody() *SaveOpenJMeterSceneResponseBody
}

type SaveOpenJMeterSceneResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveOpenJMeterSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveOpenJMeterSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveOpenJMeterSceneResponse) GoString() string {
	return s.String()
}

func (s *SaveOpenJMeterSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveOpenJMeterSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveOpenJMeterSceneResponse) GetBody() *SaveOpenJMeterSceneResponseBody {
	return s.Body
}

func (s *SaveOpenJMeterSceneResponse) SetHeaders(v map[string]*string) *SaveOpenJMeterSceneResponse {
	s.Headers = v
	return s
}

func (s *SaveOpenJMeterSceneResponse) SetStatusCode(v int32) *SaveOpenJMeterSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveOpenJMeterSceneResponse) SetBody(v *SaveOpenJMeterSceneResponseBody) *SaveOpenJMeterSceneResponse {
	s.Body = v
	return s
}

func (s *SaveOpenJMeterSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
