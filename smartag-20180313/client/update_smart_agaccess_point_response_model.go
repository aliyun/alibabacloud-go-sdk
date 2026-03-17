// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAGAccessPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSmartAGAccessPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSmartAGAccessPointResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSmartAGAccessPointResponseBody) *UpdateSmartAGAccessPointResponse
	GetBody() *UpdateSmartAGAccessPointResponseBody
}

type UpdateSmartAGAccessPointResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmartAGAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmartAGAccessPointResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAGAccessPointResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmartAGAccessPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSmartAGAccessPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSmartAGAccessPointResponse) GetBody() *UpdateSmartAGAccessPointResponseBody {
	return s.Body
}

func (s *UpdateSmartAGAccessPointResponse) SetHeaders(v map[string]*string) *UpdateSmartAGAccessPointResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmartAGAccessPointResponse) SetStatusCode(v int32) *UpdateSmartAGAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmartAGAccessPointResponse) SetBody(v *UpdateSmartAGAccessPointResponseBody) *UpdateSmartAGAccessPointResponse {
	s.Body = v
	return s
}

func (s *UpdateSmartAGAccessPointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
