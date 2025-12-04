// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSLBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSLBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSLBResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSLBResponseBody) *DeleteSLBResponse
	GetBody() *DeleteSLBResponseBody
}

type DeleteSLBResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSLBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSLBResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSLBResponse) GoString() string {
	return s.String()
}

func (s *DeleteSLBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSLBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSLBResponse) GetBody() *DeleteSLBResponseBody {
	return s.Body
}

func (s *DeleteSLBResponse) SetHeaders(v map[string]*string) *DeleteSLBResponse {
	s.Headers = v
	return s
}

func (s *DeleteSLBResponse) SetStatusCode(v int32) *DeleteSLBResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSLBResponse) SetBody(v *DeleteSLBResponseBody) *DeleteSLBResponse {
	s.Body = v
	return s
}

func (s *DeleteSLBResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
