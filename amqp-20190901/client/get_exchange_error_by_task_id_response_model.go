// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExchangeErrorByTaskIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExchangeErrorByTaskIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExchangeErrorByTaskIdResponse
	GetStatusCode() *int32
	SetBody(v *GetExchangeErrorByTaskIdResponseBody) *GetExchangeErrorByTaskIdResponse
	GetBody() *GetExchangeErrorByTaskIdResponseBody
}

type GetExchangeErrorByTaskIdResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExchangeErrorByTaskIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExchangeErrorByTaskIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExchangeErrorByTaskIdResponse) GoString() string {
	return s.String()
}

func (s *GetExchangeErrorByTaskIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExchangeErrorByTaskIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExchangeErrorByTaskIdResponse) GetBody() *GetExchangeErrorByTaskIdResponseBody {
	return s.Body
}

func (s *GetExchangeErrorByTaskIdResponse) SetHeaders(v map[string]*string) *GetExchangeErrorByTaskIdResponse {
	s.Headers = v
	return s
}

func (s *GetExchangeErrorByTaskIdResponse) SetStatusCode(v int32) *GetExchangeErrorByTaskIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExchangeErrorByTaskIdResponse) SetBody(v *GetExchangeErrorByTaskIdResponseBody) *GetExchangeErrorByTaskIdResponse {
	s.Body = v
	return s
}

func (s *GetExchangeErrorByTaskIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
