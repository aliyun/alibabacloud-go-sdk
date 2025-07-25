// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyGtmConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyGtmConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyGtmConfigResponse
	GetStatusCode() *int32
	SetBody(v *CopyGtmConfigResponseBody) *CopyGtmConfigResponse
	GetBody() *CopyGtmConfigResponseBody
}

type CopyGtmConfigResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyGtmConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyGtmConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyGtmConfigResponse) GoString() string {
	return s.String()
}

func (s *CopyGtmConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyGtmConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyGtmConfigResponse) GetBody() *CopyGtmConfigResponseBody {
	return s.Body
}

func (s *CopyGtmConfigResponse) SetHeaders(v map[string]*string) *CopyGtmConfigResponse {
	s.Headers = v
	return s
}

func (s *CopyGtmConfigResponse) SetStatusCode(v int32) *CopyGtmConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyGtmConfigResponse) SetBody(v *CopyGtmConfigResponseBody) *CopyGtmConfigResponse {
	s.Body = v
	return s
}

func (s *CopyGtmConfigResponse) Validate() error {
	return dara.Validate(s)
}
