// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYouhuiForOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateYouhuiForOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateYouhuiForOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateYouhuiForOrderResponseBody) *CreateYouhuiForOrderResponse
	GetBody() *CreateYouhuiForOrderResponseBody
}

type CreateYouhuiForOrderResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateYouhuiForOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateYouhuiForOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateYouhuiForOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateYouhuiForOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateYouhuiForOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateYouhuiForOrderResponse) GetBody() *CreateYouhuiForOrderResponseBody {
	return s.Body
}

func (s *CreateYouhuiForOrderResponse) SetHeaders(v map[string]*string) *CreateYouhuiForOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateYouhuiForOrderResponse) SetStatusCode(v int32) *CreateYouhuiForOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateYouhuiForOrderResponse) SetBody(v *CreateYouhuiForOrderResponseBody) *CreateYouhuiForOrderResponse {
	s.Body = v
	return s
}

func (s *CreateYouhuiForOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
