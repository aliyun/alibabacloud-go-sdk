// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaDNAResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaDNAResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaDNAResultResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaDNAResultResponseBody) *GetMediaDNAResultResponse
	GetBody() *GetMediaDNAResultResponseBody
}

type GetMediaDNAResultResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaDNAResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaDNAResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaDNAResultResponse) GoString() string {
	return s.String()
}

func (s *GetMediaDNAResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaDNAResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaDNAResultResponse) GetBody() *GetMediaDNAResultResponseBody {
	return s.Body
}

func (s *GetMediaDNAResultResponse) SetHeaders(v map[string]*string) *GetMediaDNAResultResponse {
	s.Headers = v
	return s
}

func (s *GetMediaDNAResultResponse) SetStatusCode(v int32) *GetMediaDNAResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaDNAResultResponse) SetBody(v *GetMediaDNAResultResponseBody) *GetMediaDNAResultResponse {
	s.Body = v
	return s
}

func (s *GetMediaDNAResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
