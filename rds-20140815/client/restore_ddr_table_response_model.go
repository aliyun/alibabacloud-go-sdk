// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreDdrTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestoreDdrTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestoreDdrTableResponse
	GetStatusCode() *int32
	SetBody(v *RestoreDdrTableResponseBody) *RestoreDdrTableResponse
	GetBody() *RestoreDdrTableResponseBody
}

type RestoreDdrTableResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreDdrTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreDdrTableResponse) String() string {
	return dara.Prettify(s)
}

func (s RestoreDdrTableResponse) GoString() string {
	return s.String()
}

func (s *RestoreDdrTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestoreDdrTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestoreDdrTableResponse) GetBody() *RestoreDdrTableResponseBody {
	return s.Body
}

func (s *RestoreDdrTableResponse) SetHeaders(v map[string]*string) *RestoreDdrTableResponse {
	s.Headers = v
	return s
}

func (s *RestoreDdrTableResponse) SetStatusCode(v int32) *RestoreDdrTableResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreDdrTableResponse) SetBody(v *RestoreDdrTableResponseBody) *RestoreDdrTableResponse {
	s.Body = v
	return s
}

func (s *RestoreDdrTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
