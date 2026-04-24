// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelOperatorOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModelOperatorOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModelOperatorOrderResponse
	GetStatusCode() *int32
	SetBody(v *GetModelOperatorOrderResponseBody) *GetModelOperatorOrderResponse
	GetBody() *GetModelOperatorOrderResponseBody
}

type GetModelOperatorOrderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelOperatorOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelOperatorOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelOperatorOrderResponse) GoString() string {
	return s.String()
}

func (s *GetModelOperatorOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModelOperatorOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModelOperatorOrderResponse) GetBody() *GetModelOperatorOrderResponseBody {
	return s.Body
}

func (s *GetModelOperatorOrderResponse) SetHeaders(v map[string]*string) *GetModelOperatorOrderResponse {
	s.Headers = v
	return s
}

func (s *GetModelOperatorOrderResponse) SetStatusCode(v int32) *GetModelOperatorOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelOperatorOrderResponse) SetBody(v *GetModelOperatorOrderResponseBody) *GetModelOperatorOrderResponse {
	s.Body = v
	return s
}

func (s *GetModelOperatorOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
