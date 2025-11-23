// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataArchiveOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataArchiveOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataArchiveOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataArchiveOrderResponseBody) *CreateDataArchiveOrderResponse
	GetBody() *CreateDataArchiveOrderResponseBody
}

type CreateDataArchiveOrderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataArchiveOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataArchiveOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataArchiveOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateDataArchiveOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataArchiveOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataArchiveOrderResponse) GetBody() *CreateDataArchiveOrderResponseBody {
	return s.Body
}

func (s *CreateDataArchiveOrderResponse) SetHeaders(v map[string]*string) *CreateDataArchiveOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateDataArchiveOrderResponse) SetStatusCode(v int32) *CreateDataArchiveOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataArchiveOrderResponse) SetBody(v *CreateDataArchiveOrderResponseBody) *CreateDataArchiveOrderResponse {
	s.Body = v
	return s
}

func (s *CreateDataArchiveOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
