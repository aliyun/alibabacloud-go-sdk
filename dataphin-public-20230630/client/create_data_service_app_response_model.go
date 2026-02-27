// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataServiceAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataServiceAppResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataServiceAppResponseBody) *CreateDataServiceAppResponse
	GetBody() *CreateDataServiceAppResponseBody
}

type CreateDataServiceAppResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataServiceAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataServiceAppResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceAppResponse) GoString() string {
	return s.String()
}

func (s *CreateDataServiceAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataServiceAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataServiceAppResponse) GetBody() *CreateDataServiceAppResponseBody {
	return s.Body
}

func (s *CreateDataServiceAppResponse) SetHeaders(v map[string]*string) *CreateDataServiceAppResponse {
	s.Headers = v
	return s
}

func (s *CreateDataServiceAppResponse) SetStatusCode(v int32) *CreateDataServiceAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataServiceAppResponse) SetBody(v *CreateDataServiceAppResponseBody) *CreateDataServiceAppResponse {
	s.Body = v
	return s
}

func (s *CreateDataServiceAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
