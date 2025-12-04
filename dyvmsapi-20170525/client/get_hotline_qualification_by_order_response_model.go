// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineQualificationByOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotlineQualificationByOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotlineQualificationByOrderResponse
	GetStatusCode() *int32
	SetBody(v *GetHotlineQualificationByOrderResponseBody) *GetHotlineQualificationByOrderResponse
	GetBody() *GetHotlineQualificationByOrderResponseBody
}

type GetHotlineQualificationByOrderResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotlineQualificationByOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotlineQualificationByOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineQualificationByOrderResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineQualificationByOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotlineQualificationByOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotlineQualificationByOrderResponse) GetBody() *GetHotlineQualificationByOrderResponseBody {
	return s.Body
}

func (s *GetHotlineQualificationByOrderResponse) SetHeaders(v map[string]*string) *GetHotlineQualificationByOrderResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineQualificationByOrderResponse) SetStatusCode(v int32) *GetHotlineQualificationByOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotlineQualificationByOrderResponse) SetBody(v *GetHotlineQualificationByOrderResponseBody) *GetHotlineQualificationByOrderResponse {
	s.Body = v
	return s
}

func (s *GetHotlineQualificationByOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
