// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserHdfsInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserHdfsInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserHdfsInfoResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserHdfsInfoResponseBody) *DeleteUserHdfsInfoResponse
	GetBody() *DeleteUserHdfsInfoResponseBody
}

type DeleteUserHdfsInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserHdfsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserHdfsInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserHdfsInfoResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserHdfsInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserHdfsInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserHdfsInfoResponse) GetBody() *DeleteUserHdfsInfoResponseBody {
	return s.Body
}

func (s *DeleteUserHdfsInfoResponse) SetHeaders(v map[string]*string) *DeleteUserHdfsInfoResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserHdfsInfoResponse) SetStatusCode(v int32) *DeleteUserHdfsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserHdfsInfoResponse) SetBody(v *DeleteUserHdfsInfoResponseBody) *DeleteUserHdfsInfoResponse {
	s.Body = v
	return s
}

func (s *DeleteUserHdfsInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
