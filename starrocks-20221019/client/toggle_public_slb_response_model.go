// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTogglePublicSlbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TogglePublicSlbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TogglePublicSlbResponse
	GetStatusCode() *int32
	SetBody(v *TogglePublicSlbResponseBody) *TogglePublicSlbResponse
	GetBody() *TogglePublicSlbResponseBody
}

type TogglePublicSlbResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TogglePublicSlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TogglePublicSlbResponse) String() string {
	return dara.Prettify(s)
}

func (s TogglePublicSlbResponse) GoString() string {
	return s.String()
}

func (s *TogglePublicSlbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TogglePublicSlbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TogglePublicSlbResponse) GetBody() *TogglePublicSlbResponseBody {
	return s.Body
}

func (s *TogglePublicSlbResponse) SetHeaders(v map[string]*string) *TogglePublicSlbResponse {
	s.Headers = v
	return s
}

func (s *TogglePublicSlbResponse) SetStatusCode(v int32) *TogglePublicSlbResponse {
	s.StatusCode = &v
	return s
}

func (s *TogglePublicSlbResponse) SetBody(v *TogglePublicSlbResponseBody) *TogglePublicSlbResponse {
	s.Body = v
	return s
}

func (s *TogglePublicSlbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
