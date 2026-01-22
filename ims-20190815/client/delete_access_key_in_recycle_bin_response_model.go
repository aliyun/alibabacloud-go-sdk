// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessKeyInRecycleBinResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAccessKeyInRecycleBinResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAccessKeyInRecycleBinResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAccessKeyInRecycleBinResponseBody) *DeleteAccessKeyInRecycleBinResponse
	GetBody() *DeleteAccessKeyInRecycleBinResponseBody
}

type DeleteAccessKeyInRecycleBinResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccessKeyInRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccessKeyInRecycleBinResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessKeyInRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessKeyInRecycleBinResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAccessKeyInRecycleBinResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAccessKeyInRecycleBinResponse) GetBody() *DeleteAccessKeyInRecycleBinResponseBody {
	return s.Body
}

func (s *DeleteAccessKeyInRecycleBinResponse) SetHeaders(v map[string]*string) *DeleteAccessKeyInRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessKeyInRecycleBinResponse) SetStatusCode(v int32) *DeleteAccessKeyInRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessKeyInRecycleBinResponse) SetBody(v *DeleteAccessKeyInRecycleBinResponseBody) *DeleteAccessKeyInRecycleBinResponse {
	s.Body = v
	return s
}

func (s *DeleteAccessKeyInRecycleBinResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
