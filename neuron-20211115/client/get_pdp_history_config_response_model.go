// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPdpHistoryConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPdpHistoryConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPdpHistoryConfigResponse
	GetStatusCode() *int32
	SetBody(v *PdpHistoryConfig) *GetPdpHistoryConfigResponse
	GetBody() *PdpHistoryConfig
}

type GetPdpHistoryConfigResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpHistoryConfig  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPdpHistoryConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPdpHistoryConfigResponse) GoString() string {
	return s.String()
}

func (s *GetPdpHistoryConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPdpHistoryConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPdpHistoryConfigResponse) GetBody() *PdpHistoryConfig {
	return s.Body
}

func (s *GetPdpHistoryConfigResponse) SetHeaders(v map[string]*string) *GetPdpHistoryConfigResponse {
	s.Headers = v
	return s
}

func (s *GetPdpHistoryConfigResponse) SetStatusCode(v int32) *GetPdpHistoryConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPdpHistoryConfigResponse) SetBody(v *PdpHistoryConfig) *GetPdpHistoryConfigResponse {
	s.Body = v
	return s
}

func (s *GetPdpHistoryConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
