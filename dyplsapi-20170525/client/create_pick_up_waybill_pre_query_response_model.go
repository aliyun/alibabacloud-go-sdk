// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePickUpWaybillPreQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePickUpWaybillPreQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePickUpWaybillPreQueryResponse
	GetStatusCode() *int32
	SetBody(v *CreatePickUpWaybillPreQueryResponseBody) *CreatePickUpWaybillPreQueryResponse
	GetBody() *CreatePickUpWaybillPreQueryResponseBody
}

type CreatePickUpWaybillPreQueryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePickUpWaybillPreQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePickUpWaybillPreQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryResponse) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePickUpWaybillPreQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePickUpWaybillPreQueryResponse) GetBody() *CreatePickUpWaybillPreQueryResponseBody {
	return s.Body
}

func (s *CreatePickUpWaybillPreQueryResponse) SetHeaders(v map[string]*string) *CreatePickUpWaybillPreQueryResponse {
	s.Headers = v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponse) SetStatusCode(v int32) *CreatePickUpWaybillPreQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponse) SetBody(v *CreatePickUpWaybillPreQueryResponseBody) *CreatePickUpWaybillPreQueryResponse {
	s.Body = v
	return s
}

func (s *CreatePickUpWaybillPreQueryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
