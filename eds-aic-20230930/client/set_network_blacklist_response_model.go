// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetNetworkBlacklistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetNetworkBlacklistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetNetworkBlacklistResponse
	GetStatusCode() *int32
	SetBody(v *SetNetworkBlacklistResponseBody) *SetNetworkBlacklistResponse
	GetBody() *SetNetworkBlacklistResponseBody
}

type SetNetworkBlacklistResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetNetworkBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetNetworkBlacklistResponse) String() string {
	return dara.Prettify(s)
}

func (s SetNetworkBlacklistResponse) GoString() string {
	return s.String()
}

func (s *SetNetworkBlacklistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetNetworkBlacklistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetNetworkBlacklistResponse) GetBody() *SetNetworkBlacklistResponseBody {
	return s.Body
}

func (s *SetNetworkBlacklistResponse) SetHeaders(v map[string]*string) *SetNetworkBlacklistResponse {
	s.Headers = v
	return s
}

func (s *SetNetworkBlacklistResponse) SetStatusCode(v int32) *SetNetworkBlacklistResponse {
	s.StatusCode = &v
	return s
}

func (s *SetNetworkBlacklistResponse) SetBody(v *SetNetworkBlacklistResponseBody) *SetNetworkBlacklistResponse {
	s.Body = v
	return s
}

func (s *SetNetworkBlacklistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
