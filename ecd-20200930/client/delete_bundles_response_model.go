// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBundlesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBundlesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBundlesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBundlesResponseBody) *DeleteBundlesResponse
	GetBody() *DeleteBundlesResponseBody
}

type DeleteBundlesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBundlesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBundlesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBundlesResponse) GoString() string {
	return s.String()
}

func (s *DeleteBundlesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBundlesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBundlesResponse) GetBody() *DeleteBundlesResponseBody {
	return s.Body
}

func (s *DeleteBundlesResponse) SetHeaders(v map[string]*string) *DeleteBundlesResponse {
	s.Headers = v
	return s
}

func (s *DeleteBundlesResponse) SetStatusCode(v int32) *DeleteBundlesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBundlesResponse) SetBody(v *DeleteBundlesResponseBody) *DeleteBundlesResponse {
	s.Body = v
	return s
}

func (s *DeleteBundlesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
