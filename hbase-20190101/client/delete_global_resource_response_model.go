// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGlobalResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGlobalResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGlobalResourceResponseBody) *DeleteGlobalResourceResponse
	GetBody() *DeleteGlobalResourceResponseBody
}

type DeleteGlobalResourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGlobalResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGlobalResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGlobalResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGlobalResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGlobalResourceResponse) GetBody() *DeleteGlobalResourceResponseBody {
	return s.Body
}

func (s *DeleteGlobalResourceResponse) SetHeaders(v map[string]*string) *DeleteGlobalResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGlobalResourceResponse) SetStatusCode(v int32) *DeleteGlobalResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGlobalResourceResponse) SetBody(v *DeleteGlobalResourceResponseBody) *DeleteGlobalResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteGlobalResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
