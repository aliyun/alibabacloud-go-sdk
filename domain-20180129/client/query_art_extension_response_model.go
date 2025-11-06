// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryArtExtensionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryArtExtensionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryArtExtensionResponse
	GetStatusCode() *int32
	SetBody(v *QueryArtExtensionResponseBody) *QueryArtExtensionResponse
	GetBody() *QueryArtExtensionResponseBody
}

type QueryArtExtensionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryArtExtensionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryArtExtensionResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryArtExtensionResponse) GoString() string {
	return s.String()
}

func (s *QueryArtExtensionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryArtExtensionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryArtExtensionResponse) GetBody() *QueryArtExtensionResponseBody {
	return s.Body
}

func (s *QueryArtExtensionResponse) SetHeaders(v map[string]*string) *QueryArtExtensionResponse {
	s.Headers = v
	return s
}

func (s *QueryArtExtensionResponse) SetStatusCode(v int32) *QueryArtExtensionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryArtExtensionResponse) SetBody(v *QueryArtExtensionResponseBody) *QueryArtExtensionResponse {
	s.Body = v
	return s
}

func (s *QueryArtExtensionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
