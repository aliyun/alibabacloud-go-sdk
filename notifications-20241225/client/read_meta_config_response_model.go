// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMetaConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadMetaConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadMetaConfigResponse
	GetStatusCode() *int32
	SetBody(v *ReadMetaConfigResponseBody) *ReadMetaConfigResponse
	GetBody() *ReadMetaConfigResponseBody
}

type ReadMetaConfigResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadMetaConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadMetaConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadMetaConfigResponse) GoString() string {
	return s.String()
}

func (s *ReadMetaConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadMetaConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadMetaConfigResponse) GetBody() *ReadMetaConfigResponseBody {
	return s.Body
}

func (s *ReadMetaConfigResponse) SetHeaders(v map[string]*string) *ReadMetaConfigResponse {
	s.Headers = v
	return s
}

func (s *ReadMetaConfigResponse) SetStatusCode(v int32) *ReadMetaConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadMetaConfigResponse) SetBody(v *ReadMetaConfigResponseBody) *ReadMetaConfigResponse {
	s.Body = v
	return s
}

func (s *ReadMetaConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
