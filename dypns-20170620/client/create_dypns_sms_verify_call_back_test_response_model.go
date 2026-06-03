// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDypnsSmsVerifyCallBackTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDypnsSmsVerifyCallBackTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDypnsSmsVerifyCallBackTestResponse
	GetStatusCode() *int32
	SetBody(v *CreateDypnsSmsVerifyCallBackTestResponseBody) *CreateDypnsSmsVerifyCallBackTestResponse
	GetBody() *CreateDypnsSmsVerifyCallBackTestResponseBody
}

type CreateDypnsSmsVerifyCallBackTestResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDypnsSmsVerifyCallBackTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDypnsSmsVerifyCallBackTestResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDypnsSmsVerifyCallBackTestResponse) GoString() string {
	return s.String()
}

func (s *CreateDypnsSmsVerifyCallBackTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDypnsSmsVerifyCallBackTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDypnsSmsVerifyCallBackTestResponse) GetBody() *CreateDypnsSmsVerifyCallBackTestResponseBody {
	return s.Body
}

func (s *CreateDypnsSmsVerifyCallBackTestResponse) SetHeaders(v map[string]*string) *CreateDypnsSmsVerifyCallBackTestResponse {
	s.Headers = v
	return s
}

func (s *CreateDypnsSmsVerifyCallBackTestResponse) SetStatusCode(v int32) *CreateDypnsSmsVerifyCallBackTestResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDypnsSmsVerifyCallBackTestResponse) SetBody(v *CreateDypnsSmsVerifyCallBackTestResponseBody) *CreateDypnsSmsVerifyCallBackTestResponse {
	s.Body = v
	return s
}

func (s *CreateDypnsSmsVerifyCallBackTestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
