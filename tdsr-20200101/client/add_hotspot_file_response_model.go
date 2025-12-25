// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddHotspotFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddHotspotFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddHotspotFileResponse
	GetStatusCode() *int32
	SetBody(v *AddHotspotFileResponseBody) *AddHotspotFileResponse
	GetBody() *AddHotspotFileResponseBody
}

type AddHotspotFileResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddHotspotFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddHotspotFileResponse) String() string {
	return dara.Prettify(s)
}

func (s AddHotspotFileResponse) GoString() string {
	return s.String()
}

func (s *AddHotspotFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddHotspotFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddHotspotFileResponse) GetBody() *AddHotspotFileResponseBody {
	return s.Body
}

func (s *AddHotspotFileResponse) SetHeaders(v map[string]*string) *AddHotspotFileResponse {
	s.Headers = v
	return s
}

func (s *AddHotspotFileResponse) SetStatusCode(v int32) *AddHotspotFileResponse {
	s.StatusCode = &v
	return s
}

func (s *AddHotspotFileResponse) SetBody(v *AddHotspotFileResponseBody) *AddHotspotFileResponse {
	s.Body = v
	return s
}

func (s *AddHotspotFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
