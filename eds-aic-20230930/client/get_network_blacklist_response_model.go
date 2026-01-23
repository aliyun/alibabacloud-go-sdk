// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkBlacklistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNetworkBlacklistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNetworkBlacklistResponse
	GetStatusCode() *int32
	SetBody(v *GetNetworkBlacklistResponseBody) *GetNetworkBlacklistResponse
	GetBody() *GetNetworkBlacklistResponseBody
}

type GetNetworkBlacklistResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNetworkBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNetworkBlacklistResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkBlacklistResponse) GoString() string {
	return s.String()
}

func (s *GetNetworkBlacklistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNetworkBlacklistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNetworkBlacklistResponse) GetBody() *GetNetworkBlacklistResponseBody {
	return s.Body
}

func (s *GetNetworkBlacklistResponse) SetHeaders(v map[string]*string) *GetNetworkBlacklistResponse {
	s.Headers = v
	return s
}

func (s *GetNetworkBlacklistResponse) SetStatusCode(v int32) *GetNetworkBlacklistResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNetworkBlacklistResponse) SetBody(v *GetNetworkBlacklistResponseBody) *GetNetworkBlacklistResponse {
	s.Body = v
	return s
}

func (s *GetNetworkBlacklistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
