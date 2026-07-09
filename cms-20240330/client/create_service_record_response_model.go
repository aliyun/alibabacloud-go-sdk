// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceRecordResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceRecordResponseBody) *CreateServiceRecordResponse
	GetBody() *CreateServiceRecordResponseBody
}

type CreateServiceRecordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceRecordResponse) GetBody() *CreateServiceRecordResponseBody {
	return s.Body
}

func (s *CreateServiceRecordResponse) SetHeaders(v map[string]*string) *CreateServiceRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceRecordResponse) SetStatusCode(v int32) *CreateServiceRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceRecordResponse) SetBody(v *CreateServiceRecordResponseBody) *CreateServiceRecordResponse {
	s.Body = v
	return s
}

func (s *CreateServiceRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
