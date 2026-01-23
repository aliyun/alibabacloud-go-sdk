// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityIdentifyResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecurityIdentifyResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecurityIdentifyResultResponse
	GetStatusCode() *int32
	SetBody(v *GetSecurityIdentifyResultResponseBody) *GetSecurityIdentifyResultResponse
	GetBody() *GetSecurityIdentifyResultResponseBody
}

type GetSecurityIdentifyResultResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecurityIdentifyResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecurityIdentifyResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityIdentifyResultResponse) GoString() string {
	return s.String()
}

func (s *GetSecurityIdentifyResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecurityIdentifyResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecurityIdentifyResultResponse) GetBody() *GetSecurityIdentifyResultResponseBody {
	return s.Body
}

func (s *GetSecurityIdentifyResultResponse) SetHeaders(v map[string]*string) *GetSecurityIdentifyResultResponse {
	s.Headers = v
	return s
}

func (s *GetSecurityIdentifyResultResponse) SetStatusCode(v int32) *GetSecurityIdentifyResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecurityIdentifyResultResponse) SetBody(v *GetSecurityIdentifyResultResponseBody) *GetSecurityIdentifyResultResponse {
	s.Body = v
	return s
}

func (s *GetSecurityIdentifyResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
