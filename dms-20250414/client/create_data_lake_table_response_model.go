// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLakeTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataLakeTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataLakeTableResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataLakeTableResponseBody) *CreateDataLakeTableResponse
	GetBody() *CreateDataLakeTableResponseBody
}

type CreateDataLakeTableResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataLakeTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataLakeTableResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLakeTableResponse) GoString() string {
	return s.String()
}

func (s *CreateDataLakeTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataLakeTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataLakeTableResponse) GetBody() *CreateDataLakeTableResponseBody {
	return s.Body
}

func (s *CreateDataLakeTableResponse) SetHeaders(v map[string]*string) *CreateDataLakeTableResponse {
	s.Headers = v
	return s
}

func (s *CreateDataLakeTableResponse) SetStatusCode(v int32) *CreateDataLakeTableResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataLakeTableResponse) SetBody(v *CreateDataLakeTableResponseBody) *CreateDataLakeTableResponse {
	s.Body = v
	return s
}

func (s *CreateDataLakeTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
