// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveModelConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveModelConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveModelConfigResponse
	GetStatusCode() *int32
	SetBody(v *SaveModelConfigResponseBody) *SaveModelConfigResponse
	GetBody() *SaveModelConfigResponseBody
}

type SaveModelConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveModelConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveModelConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveModelConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveModelConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveModelConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveModelConfigResponse) GetBody() *SaveModelConfigResponseBody {
	return s.Body
}

func (s *SaveModelConfigResponse) SetHeaders(v map[string]*string) *SaveModelConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveModelConfigResponse) SetStatusCode(v int32) *SaveModelConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveModelConfigResponse) SetBody(v *SaveModelConfigResponseBody) *SaveModelConfigResponse {
	s.Body = v
	return s
}

func (s *SaveModelConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
