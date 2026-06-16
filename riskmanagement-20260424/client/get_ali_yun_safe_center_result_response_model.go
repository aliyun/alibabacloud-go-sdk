// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAliYunSafeCenterResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAliYunSafeCenterResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAliYunSafeCenterResultResponse
	GetStatusCode() *int32
	SetBody(v *GetAliYunSafeCenterResultResponseBody) *GetAliYunSafeCenterResultResponse
	GetBody() *GetAliYunSafeCenterResultResponseBody
}

type GetAliYunSafeCenterResultResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAliYunSafeCenterResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAliYunSafeCenterResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultResponse) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAliYunSafeCenterResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAliYunSafeCenterResultResponse) GetBody() *GetAliYunSafeCenterResultResponseBody {
	return s.Body
}

func (s *GetAliYunSafeCenterResultResponse) SetHeaders(v map[string]*string) *GetAliYunSafeCenterResultResponse {
	s.Headers = v
	return s
}

func (s *GetAliYunSafeCenterResultResponse) SetStatusCode(v int32) *GetAliYunSafeCenterResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAliYunSafeCenterResultResponse) SetBody(v *GetAliYunSafeCenterResultResponseBody) *GetAliYunSafeCenterResultResponse {
	s.Body = v
	return s
}

func (s *GetAliYunSafeCenterResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
