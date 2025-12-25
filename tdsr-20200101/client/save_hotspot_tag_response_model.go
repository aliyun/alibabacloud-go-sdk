// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveHotspotTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveHotspotTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveHotspotTagResponse
	GetStatusCode() *int32
	SetBody(v *SaveHotspotTagResponseBody) *SaveHotspotTagResponse
	GetBody() *SaveHotspotTagResponseBody
}

type SaveHotspotTagResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveHotspotTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveHotspotTagResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveHotspotTagResponse) GoString() string {
	return s.String()
}

func (s *SaveHotspotTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveHotspotTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveHotspotTagResponse) GetBody() *SaveHotspotTagResponseBody {
	return s.Body
}

func (s *SaveHotspotTagResponse) SetHeaders(v map[string]*string) *SaveHotspotTagResponse {
	s.Headers = v
	return s
}

func (s *SaveHotspotTagResponse) SetStatusCode(v int32) *SaveHotspotTagResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveHotspotTagResponse) SetBody(v *SaveHotspotTagResponseBody) *SaveHotspotTagResponse {
	s.Body = v
	return s
}

func (s *SaveHotspotTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
