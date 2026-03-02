// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLogUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLogUrlResponse
	GetStatusCode() *int32
	SetBody(v *PdpSlsLogInfo) *GetLogUrlResponse
	GetBody() *PdpSlsLogInfo
}

type GetLogUrlResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpSlsLogInfo     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLogUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLogUrlResponse) GoString() string {
	return s.String()
}

func (s *GetLogUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLogUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLogUrlResponse) GetBody() *PdpSlsLogInfo {
	return s.Body
}

func (s *GetLogUrlResponse) SetHeaders(v map[string]*string) *GetLogUrlResponse {
	s.Headers = v
	return s
}

func (s *GetLogUrlResponse) SetStatusCode(v int32) *GetLogUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogUrlResponse) SetBody(v *PdpSlsLogInfo) *GetLogUrlResponse {
	s.Body = v
	return s
}

func (s *GetLogUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
