// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCallCenterProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCallCenterProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCallCenterProviderResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCallCenterProviderResponseBody) *DeleteCallCenterProviderResponse
	GetBody() *DeleteCallCenterProviderResponseBody
}

type DeleteCallCenterProviderResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCallCenterProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCallCenterProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCallCenterProviderResponse) GoString() string {
	return s.String()
}

func (s *DeleteCallCenterProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCallCenterProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCallCenterProviderResponse) GetBody() *DeleteCallCenterProviderResponseBody {
	return s.Body
}

func (s *DeleteCallCenterProviderResponse) SetHeaders(v map[string]*string) *DeleteCallCenterProviderResponse {
	s.Headers = v
	return s
}

func (s *DeleteCallCenterProviderResponse) SetStatusCode(v int32) *DeleteCallCenterProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCallCenterProviderResponse) SetBody(v *DeleteCallCenterProviderResponseBody) *DeleteCallCenterProviderResponse {
	s.Body = v
	return s
}

func (s *DeleteCallCenterProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
