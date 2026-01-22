// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeprovisionApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeprovisionApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeprovisionApplicationResponse
	GetStatusCode() *int32
	SetBody(v *DeprovisionApplicationResponseBody) *DeprovisionApplicationResponse
	GetBody() *DeprovisionApplicationResponseBody
}

type DeprovisionApplicationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeprovisionApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeprovisionApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeprovisionApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeprovisionApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeprovisionApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeprovisionApplicationResponse) GetBody() *DeprovisionApplicationResponseBody {
	return s.Body
}

func (s *DeprovisionApplicationResponse) SetHeaders(v map[string]*string) *DeprovisionApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeprovisionApplicationResponse) SetStatusCode(v int32) *DeprovisionApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeprovisionApplicationResponse) SetBody(v *DeprovisionApplicationResponseBody) *DeprovisionApplicationResponse {
	s.Body = v
	return s
}

func (s *DeprovisionApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
