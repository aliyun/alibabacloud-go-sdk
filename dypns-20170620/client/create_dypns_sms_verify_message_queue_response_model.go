// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDypnsSmsVerifyMessageQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDypnsSmsVerifyMessageQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDypnsSmsVerifyMessageQueueResponse
	GetStatusCode() *int32
	SetBody(v *CreateDypnsSmsVerifyMessageQueueResponseBody) *CreateDypnsSmsVerifyMessageQueueResponse
	GetBody() *CreateDypnsSmsVerifyMessageQueueResponseBody
}

type CreateDypnsSmsVerifyMessageQueueResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDypnsSmsVerifyMessageQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDypnsSmsVerifyMessageQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDypnsSmsVerifyMessageQueueResponse) GoString() string {
	return s.String()
}

func (s *CreateDypnsSmsVerifyMessageQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDypnsSmsVerifyMessageQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDypnsSmsVerifyMessageQueueResponse) GetBody() *CreateDypnsSmsVerifyMessageQueueResponseBody {
	return s.Body
}

func (s *CreateDypnsSmsVerifyMessageQueueResponse) SetHeaders(v map[string]*string) *CreateDypnsSmsVerifyMessageQueueResponse {
	s.Headers = v
	return s
}

func (s *CreateDypnsSmsVerifyMessageQueueResponse) SetStatusCode(v int32) *CreateDypnsSmsVerifyMessageQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageQueueResponse) SetBody(v *CreateDypnsSmsVerifyMessageQueueResponseBody) *CreateDypnsSmsVerifyMessageQueueResponse {
	s.Body = v
	return s
}

func (s *CreateDypnsSmsVerifyMessageQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
