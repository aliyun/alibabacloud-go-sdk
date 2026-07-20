// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGuideSubStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryGuideSubStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryGuideSubStatusResponse
	GetStatusCode() *int32
	SetBody(v *QueryGuideSubStatusResponseBody) *QueryGuideSubStatusResponse
	GetBody() *QueryGuideSubStatusResponseBody
}

type QueryGuideSubStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGuideSubStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGuideSubStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryGuideSubStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryGuideSubStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryGuideSubStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryGuideSubStatusResponse) GetBody() *QueryGuideSubStatusResponseBody {
	return s.Body
}

func (s *QueryGuideSubStatusResponse) SetHeaders(v map[string]*string) *QueryGuideSubStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryGuideSubStatusResponse) SetStatusCode(v int32) *QueryGuideSubStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGuideSubStatusResponse) SetBody(v *QueryGuideSubStatusResponseBody) *QueryGuideSubStatusResponse {
	s.Body = v
	return s
}

func (s *QueryGuideSubStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
