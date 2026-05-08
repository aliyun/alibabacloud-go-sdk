// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateYikeProductionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateYikeProductionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateYikeProductionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateYikeProductionResponseBody) *UpdateYikeProductionResponse
	GetBody() *UpdateYikeProductionResponseBody
}

type UpdateYikeProductionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateYikeProductionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateYikeProductionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateYikeProductionResponse) GoString() string {
	return s.String()
}

func (s *UpdateYikeProductionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateYikeProductionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateYikeProductionResponse) GetBody() *UpdateYikeProductionResponseBody {
	return s.Body
}

func (s *UpdateYikeProductionResponse) SetHeaders(v map[string]*string) *UpdateYikeProductionResponse {
	s.Headers = v
	return s
}

func (s *UpdateYikeProductionResponse) SetStatusCode(v int32) *UpdateYikeProductionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateYikeProductionResponse) SetBody(v *UpdateYikeProductionResponseBody) *UpdateYikeProductionResponse {
	s.Body = v
	return s
}

func (s *UpdateYikeProductionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
