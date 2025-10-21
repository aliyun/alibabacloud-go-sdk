// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUdfFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUdfFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUdfFunctionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUdfFunctionResponseBody) *DeleteUdfFunctionResponse
	GetBody() *DeleteUdfFunctionResponseBody
}

type DeleteUdfFunctionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUdfFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUdfFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUdfFunctionResponse) GoString() string {
	return s.String()
}

func (s *DeleteUdfFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUdfFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUdfFunctionResponse) GetBody() *DeleteUdfFunctionResponseBody {
	return s.Body
}

func (s *DeleteUdfFunctionResponse) SetHeaders(v map[string]*string) *DeleteUdfFunctionResponse {
	s.Headers = v
	return s
}

func (s *DeleteUdfFunctionResponse) SetStatusCode(v int32) *DeleteUdfFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUdfFunctionResponse) SetBody(v *DeleteUdfFunctionResponseBody) *DeleteUdfFunctionResponse {
	s.Body = v
	return s
}

func (s *DeleteUdfFunctionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
