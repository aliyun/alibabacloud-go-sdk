// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRemoteADBDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRemoteADBDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRemoteADBDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRemoteADBDataSourceResponseBody) *DeleteRemoteADBDataSourceResponse
	GetBody() *DeleteRemoteADBDataSourceResponseBody
}

type DeleteRemoteADBDataSourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRemoteADBDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRemoteADBDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRemoteADBDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteRemoteADBDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRemoteADBDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRemoteADBDataSourceResponse) GetBody() *DeleteRemoteADBDataSourceResponseBody {
	return s.Body
}

func (s *DeleteRemoteADBDataSourceResponse) SetHeaders(v map[string]*string) *DeleteRemoteADBDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteRemoteADBDataSourceResponse) SetStatusCode(v int32) *DeleteRemoteADBDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRemoteADBDataSourceResponse) SetBody(v *DeleteRemoteADBDataSourceResponseBody) *DeleteRemoteADBDataSourceResponse {
	s.Body = v
	return s
}

func (s *DeleteRemoteADBDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
