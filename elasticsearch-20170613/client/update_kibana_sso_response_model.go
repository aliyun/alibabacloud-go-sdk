// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKibanaSsoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKibanaSsoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKibanaSsoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKibanaSsoResponseBody) *UpdateKibanaSsoResponse
	GetBody() *UpdateKibanaSsoResponseBody
}

type UpdateKibanaSsoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKibanaSsoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKibanaSsoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKibanaSsoResponse) GoString() string {
	return s.String()
}

func (s *UpdateKibanaSsoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKibanaSsoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKibanaSsoResponse) GetBody() *UpdateKibanaSsoResponseBody {
	return s.Body
}

func (s *UpdateKibanaSsoResponse) SetHeaders(v map[string]*string) *UpdateKibanaSsoResponse {
	s.Headers = v
	return s
}

func (s *UpdateKibanaSsoResponse) SetStatusCode(v int32) *UpdateKibanaSsoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKibanaSsoResponse) SetBody(v *UpdateKibanaSsoResponseBody) *UpdateKibanaSsoResponse {
	s.Body = v
	return s
}

func (s *UpdateKibanaSsoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
