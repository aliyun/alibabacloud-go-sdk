// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakeFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataLakeFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataLakeFunctionResponse
	GetStatusCode() *int32
	SetBody(v *GetDataLakeFunctionResponseBody) *GetDataLakeFunctionResponse
	GetBody() *GetDataLakeFunctionResponseBody
}

type GetDataLakeFunctionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataLakeFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataLakeFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakeFunctionResponse) GoString() string {
	return s.String()
}

func (s *GetDataLakeFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataLakeFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataLakeFunctionResponse) GetBody() *GetDataLakeFunctionResponseBody {
	return s.Body
}

func (s *GetDataLakeFunctionResponse) SetHeaders(v map[string]*string) *GetDataLakeFunctionResponse {
	s.Headers = v
	return s
}

func (s *GetDataLakeFunctionResponse) SetStatusCode(v int32) *GetDataLakeFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataLakeFunctionResponse) SetBody(v *GetDataLakeFunctionResponseBody) *GetDataLakeFunctionResponse {
	s.Body = v
	return s
}

func (s *GetDataLakeFunctionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
