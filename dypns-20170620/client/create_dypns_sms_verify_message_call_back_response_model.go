// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDypnsSmsVerifyMessageCallBackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDypnsSmsVerifyMessageCallBackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDypnsSmsVerifyMessageCallBackResponse
	GetStatusCode() *int32
	SetBody(v *CreateDypnsSmsVerifyMessageCallBackResponseBody) *CreateDypnsSmsVerifyMessageCallBackResponse
	GetBody() *CreateDypnsSmsVerifyMessageCallBackResponseBody
}

type CreateDypnsSmsVerifyMessageCallBackResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDypnsSmsVerifyMessageCallBackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDypnsSmsVerifyMessageCallBackResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDypnsSmsVerifyMessageCallBackResponse) GoString() string {
	return s.String()
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponse) GetBody() *CreateDypnsSmsVerifyMessageCallBackResponseBody {
	return s.Body
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponse) SetHeaders(v map[string]*string) *CreateDypnsSmsVerifyMessageCallBackResponse {
	s.Headers = v
	return s
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponse) SetStatusCode(v int32) *CreateDypnsSmsVerifyMessageCallBackResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponse) SetBody(v *CreateDypnsSmsVerifyMessageCallBackResponseBody) *CreateDypnsSmsVerifyMessageCallBackResponse {
	s.Body = v
	return s
}

func (s *CreateDypnsSmsVerifyMessageCallBackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
