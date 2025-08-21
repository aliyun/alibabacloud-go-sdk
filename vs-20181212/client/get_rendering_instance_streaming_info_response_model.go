// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRenderingInstanceStreamingInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRenderingInstanceStreamingInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRenderingInstanceStreamingInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetRenderingInstanceStreamingInfoResponseBody) *GetRenderingInstanceStreamingInfoResponse
	GetBody() *GetRenderingInstanceStreamingInfoResponseBody
}

type GetRenderingInstanceStreamingInfoResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRenderingInstanceStreamingInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRenderingInstanceStreamingInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRenderingInstanceStreamingInfoResponse) GoString() string {
	return s.String()
}

func (s *GetRenderingInstanceStreamingInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRenderingInstanceStreamingInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRenderingInstanceStreamingInfoResponse) GetBody() *GetRenderingInstanceStreamingInfoResponseBody {
	return s.Body
}

func (s *GetRenderingInstanceStreamingInfoResponse) SetHeaders(v map[string]*string) *GetRenderingInstanceStreamingInfoResponse {
	s.Headers = v
	return s
}

func (s *GetRenderingInstanceStreamingInfoResponse) SetStatusCode(v int32) *GetRenderingInstanceStreamingInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRenderingInstanceStreamingInfoResponse) SetBody(v *GetRenderingInstanceStreamingInfoResponseBody) *GetRenderingInstanceStreamingInfoResponse {
	s.Body = v
	return s
}

func (s *GetRenderingInstanceStreamingInfoResponse) Validate() error {
	return dara.Validate(s)
}
