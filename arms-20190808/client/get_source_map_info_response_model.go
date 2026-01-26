// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceMapInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSourceMapInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSourceMapInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetSourceMapInfoResponseBody) *GetSourceMapInfoResponse
	GetBody() *GetSourceMapInfoResponseBody
}

type GetSourceMapInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSourceMapInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSourceMapInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSourceMapInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSourceMapInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSourceMapInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSourceMapInfoResponse) GetBody() *GetSourceMapInfoResponseBody {
	return s.Body
}

func (s *GetSourceMapInfoResponse) SetHeaders(v map[string]*string) *GetSourceMapInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSourceMapInfoResponse) SetStatusCode(v int32) *GetSourceMapInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSourceMapInfoResponse) SetBody(v *GetSourceMapInfoResponseBody) *GetSourceMapInfoResponse {
	s.Body = v
	return s
}

func (s *GetSourceMapInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
