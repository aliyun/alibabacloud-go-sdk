// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYikeProductionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateYikeProductionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateYikeProductionResponse
	GetStatusCode() *int32
	SetBody(v *CreateYikeProductionResponseBody) *CreateYikeProductionResponse
	GetBody() *CreateYikeProductionResponseBody
}

type CreateYikeProductionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateYikeProductionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateYikeProductionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateYikeProductionResponse) GoString() string {
	return s.String()
}

func (s *CreateYikeProductionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateYikeProductionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateYikeProductionResponse) GetBody() *CreateYikeProductionResponseBody {
	return s.Body
}

func (s *CreateYikeProductionResponse) SetHeaders(v map[string]*string) *CreateYikeProductionResponse {
	s.Headers = v
	return s
}

func (s *CreateYikeProductionResponse) SetStatusCode(v int32) *CreateYikeProductionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateYikeProductionResponse) SetBody(v *CreateYikeProductionResponseBody) *CreateYikeProductionResponse {
	s.Body = v
	return s
}

func (s *CreateYikeProductionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
