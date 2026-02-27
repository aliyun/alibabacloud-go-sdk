// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataServiceAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataServiceAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataServiceAppResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataServiceAppResponseBody) *UpdateDataServiceAppResponse
	GetBody() *UpdateDataServiceAppResponseBody
}

type UpdateDataServiceAppResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataServiceAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataServiceAppResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataServiceAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataServiceAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataServiceAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataServiceAppResponse) GetBody() *UpdateDataServiceAppResponseBody {
	return s.Body
}

func (s *UpdateDataServiceAppResponse) SetHeaders(v map[string]*string) *UpdateDataServiceAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataServiceAppResponse) SetStatusCode(v int32) *UpdateDataServiceAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataServiceAppResponse) SetBody(v *UpdateDataServiceAppResponseBody) *UpdateDataServiceAppResponse {
	s.Body = v
	return s
}

func (s *UpdateDataServiceAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
