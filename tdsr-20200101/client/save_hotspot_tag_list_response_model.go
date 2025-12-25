// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveHotspotTagListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveHotspotTagListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveHotspotTagListResponse
	GetStatusCode() *int32
	SetBody(v *SaveHotspotTagListResponseBody) *SaveHotspotTagListResponse
	GetBody() *SaveHotspotTagListResponseBody
}

type SaveHotspotTagListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveHotspotTagListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveHotspotTagListResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveHotspotTagListResponse) GoString() string {
	return s.String()
}

func (s *SaveHotspotTagListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveHotspotTagListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveHotspotTagListResponse) GetBody() *SaveHotspotTagListResponseBody {
	return s.Body
}

func (s *SaveHotspotTagListResponse) SetHeaders(v map[string]*string) *SaveHotspotTagListResponse {
	s.Headers = v
	return s
}

func (s *SaveHotspotTagListResponse) SetStatusCode(v int32) *SaveHotspotTagListResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveHotspotTagListResponse) SetBody(v *SaveHotspotTagListResponseBody) *SaveHotspotTagListResponse {
	s.Body = v
	return s
}

func (s *SaveHotspotTagListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
