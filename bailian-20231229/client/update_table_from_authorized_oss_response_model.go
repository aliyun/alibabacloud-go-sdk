// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableFromAuthorizedOssResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTableFromAuthorizedOssResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTableFromAuthorizedOssResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTableFromAuthorizedOssResponseBody) *UpdateTableFromAuthorizedOssResponse
	GetBody() *UpdateTableFromAuthorizedOssResponseBody
}

type UpdateTableFromAuthorizedOssResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTableFromAuthorizedOssResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTableFromAuthorizedOssResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableFromAuthorizedOssResponse) GoString() string {
	return s.String()
}

func (s *UpdateTableFromAuthorizedOssResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTableFromAuthorizedOssResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTableFromAuthorizedOssResponse) GetBody() *UpdateTableFromAuthorizedOssResponseBody {
	return s.Body
}

func (s *UpdateTableFromAuthorizedOssResponse) SetHeaders(v map[string]*string) *UpdateTableFromAuthorizedOssResponse {
	s.Headers = v
	return s
}

func (s *UpdateTableFromAuthorizedOssResponse) SetStatusCode(v int32) *UpdateTableFromAuthorizedOssResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTableFromAuthorizedOssResponse) SetBody(v *UpdateTableFromAuthorizedOssResponseBody) *UpdateTableFromAuthorizedOssResponse {
	s.Body = v
	return s
}

func (s *UpdateTableFromAuthorizedOssResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
