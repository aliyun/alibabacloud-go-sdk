// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogDeliverySLRResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLogDeliverySLRResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLogDeliverySLRResponse
	GetStatusCode() *int32
	SetBody(v *CreateLogDeliverySLRResponseBody) *CreateLogDeliverySLRResponse
	GetBody() *CreateLogDeliverySLRResponseBody
}

type CreateLogDeliverySLRResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLogDeliverySLRResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLogDeliverySLRResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLogDeliverySLRResponse) GoString() string {
	return s.String()
}

func (s *CreateLogDeliverySLRResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLogDeliverySLRResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLogDeliverySLRResponse) GetBody() *CreateLogDeliverySLRResponseBody {
	return s.Body
}

func (s *CreateLogDeliverySLRResponse) SetHeaders(v map[string]*string) *CreateLogDeliverySLRResponse {
	s.Headers = v
	return s
}

func (s *CreateLogDeliverySLRResponse) SetStatusCode(v int32) *CreateLogDeliverySLRResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLogDeliverySLRResponse) SetBody(v *CreateLogDeliverySLRResponseBody) *CreateLogDeliverySLRResponse {
	s.Body = v
	return s
}

func (s *CreateLogDeliverySLRResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
