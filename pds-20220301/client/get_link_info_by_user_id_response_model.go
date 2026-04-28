// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLinkInfoByUserIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLinkInfoByUserIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLinkInfoByUserIdResponse
	GetStatusCode() *int32
	SetBody(v *GetLinkInfoByUserIdResponseBody) *GetLinkInfoByUserIdResponse
	GetBody() *GetLinkInfoByUserIdResponseBody
}

type GetLinkInfoByUserIdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLinkInfoByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLinkInfoByUserIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLinkInfoByUserIdResponse) GoString() string {
	return s.String()
}

func (s *GetLinkInfoByUserIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLinkInfoByUserIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLinkInfoByUserIdResponse) GetBody() *GetLinkInfoByUserIdResponseBody {
	return s.Body
}

func (s *GetLinkInfoByUserIdResponse) SetHeaders(v map[string]*string) *GetLinkInfoByUserIdResponse {
	s.Headers = v
	return s
}

func (s *GetLinkInfoByUserIdResponse) SetStatusCode(v int32) *GetLinkInfoByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLinkInfoByUserIdResponse) SetBody(v *GetLinkInfoByUserIdResponseBody) *GetLinkInfoByUserIdResponse {
	s.Body = v
	return s
}

func (s *GetLinkInfoByUserIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
