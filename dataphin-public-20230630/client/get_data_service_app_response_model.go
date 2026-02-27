// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServiceAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServiceAppResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServiceAppResponseBody) *GetDataServiceAppResponse
	GetBody() *GetDataServiceAppResponseBody
}

type GetDataServiceAppResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceAppResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServiceAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServiceAppResponse) GetBody() *GetDataServiceAppResponseBody {
	return s.Body
}

func (s *GetDataServiceAppResponse) SetHeaders(v map[string]*string) *GetDataServiceAppResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceAppResponse) SetStatusCode(v int32) *GetDataServiceAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceAppResponse) SetBody(v *GetDataServiceAppResponseBody) *GetDataServiceAppResponse {
	s.Body = v
	return s
}

func (s *GetDataServiceAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
