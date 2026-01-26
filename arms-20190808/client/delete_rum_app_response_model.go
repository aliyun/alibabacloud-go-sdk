// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRumAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRumAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRumAppResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRumAppResponseBody) *DeleteRumAppResponse
	GetBody() *DeleteRumAppResponseBody
}

type DeleteRumAppResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRumAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRumAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRumAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteRumAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRumAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRumAppResponse) GetBody() *DeleteRumAppResponseBody {
	return s.Body
}

func (s *DeleteRumAppResponse) SetHeaders(v map[string]*string) *DeleteRumAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteRumAppResponse) SetStatusCode(v int32) *DeleteRumAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRumAppResponse) SetBody(v *DeleteRumAppResponseBody) *DeleteRumAppResponse {
	s.Body = v
	return s
}

func (s *DeleteRumAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
