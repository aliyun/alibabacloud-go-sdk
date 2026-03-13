// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAliUidByInstanceIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserAliUidByInstanceIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserAliUidByInstanceIdResponse
	GetStatusCode() *int32
	SetBody(v *GetUserAliUidByInstanceIdResponseBody) *GetUserAliUidByInstanceIdResponse
	GetBody() *GetUserAliUidByInstanceIdResponseBody
}

type GetUserAliUidByInstanceIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserAliUidByInstanceIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserAliUidByInstanceIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserAliUidByInstanceIdResponse) GoString() string {
	return s.String()
}

func (s *GetUserAliUidByInstanceIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserAliUidByInstanceIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserAliUidByInstanceIdResponse) GetBody() *GetUserAliUidByInstanceIdResponseBody {
	return s.Body
}

func (s *GetUserAliUidByInstanceIdResponse) SetHeaders(v map[string]*string) *GetUserAliUidByInstanceIdResponse {
	s.Headers = v
	return s
}

func (s *GetUserAliUidByInstanceIdResponse) SetStatusCode(v int32) *GetUserAliUidByInstanceIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserAliUidByInstanceIdResponse) SetBody(v *GetUserAliUidByInstanceIdResponseBody) *GetUserAliUidByInstanceIdResponse {
	s.Body = v
	return s
}

func (s *GetUserAliUidByInstanceIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
