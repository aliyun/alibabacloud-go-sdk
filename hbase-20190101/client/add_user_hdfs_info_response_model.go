// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserHdfsInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUserHdfsInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUserHdfsInfoResponse
	GetStatusCode() *int32
	SetBody(v *AddUserHdfsInfoResponseBody) *AddUserHdfsInfoResponse
	GetBody() *AddUserHdfsInfoResponseBody
}

type AddUserHdfsInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserHdfsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserHdfsInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUserHdfsInfoResponse) GoString() string {
	return s.String()
}

func (s *AddUserHdfsInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUserHdfsInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUserHdfsInfoResponse) GetBody() *AddUserHdfsInfoResponseBody {
	return s.Body
}

func (s *AddUserHdfsInfoResponse) SetHeaders(v map[string]*string) *AddUserHdfsInfoResponse {
	s.Headers = v
	return s
}

func (s *AddUserHdfsInfoResponse) SetStatusCode(v int32) *AddUserHdfsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserHdfsInfoResponse) SetBody(v *AddUserHdfsInfoResponseBody) *AddUserHdfsInfoResponse {
	s.Body = v
	return s
}

func (s *AddUserHdfsInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
