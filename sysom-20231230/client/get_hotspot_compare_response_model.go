// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotCompareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotspotCompareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotspotCompareResponse
	GetStatusCode() *int32
	SetBody(v *GetHotspotCompareResponseBody) *GetHotspotCompareResponse
	GetBody() *GetHotspotCompareResponseBody
}

type GetHotspotCompareResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotspotCompareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotspotCompareResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotCompareResponse) GoString() string {
	return s.String()
}

func (s *GetHotspotCompareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotspotCompareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotspotCompareResponse) GetBody() *GetHotspotCompareResponseBody {
	return s.Body
}

func (s *GetHotspotCompareResponse) SetHeaders(v map[string]*string) *GetHotspotCompareResponse {
	s.Headers = v
	return s
}

func (s *GetHotspotCompareResponse) SetStatusCode(v int32) *GetHotspotCompareResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotspotCompareResponse) SetBody(v *GetHotspotCompareResponseBody) *GetHotspotCompareResponse {
	s.Body = v
	return s
}

func (s *GetHotspotCompareResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
