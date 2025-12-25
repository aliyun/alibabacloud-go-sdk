// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotspotTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotspotTagResponse
	GetStatusCode() *int32
	SetBody(v *GetHotspotTagResponseBody) *GetHotspotTagResponse
	GetBody() *GetHotspotTagResponseBody
}

type GetHotspotTagResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotspotTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotspotTagResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotTagResponse) GoString() string {
	return s.String()
}

func (s *GetHotspotTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotspotTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotspotTagResponse) GetBody() *GetHotspotTagResponseBody {
	return s.Body
}

func (s *GetHotspotTagResponse) SetHeaders(v map[string]*string) *GetHotspotTagResponse {
	s.Headers = v
	return s
}

func (s *GetHotspotTagResponse) SetStatusCode(v int32) *GetHotspotTagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotspotTagResponse) SetBody(v *GetHotspotTagResponseBody) *GetHotspotTagResponse {
	s.Body = v
	return s
}

func (s *GetHotspotTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
