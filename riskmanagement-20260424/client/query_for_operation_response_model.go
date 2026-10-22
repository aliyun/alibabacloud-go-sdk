// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryForOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryForOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryForOperationResponse
	GetStatusCode() *int32
	SetBody(v *QueryForOperationResponseBody) *QueryForOperationResponse
	GetBody() *QueryForOperationResponseBody
}

type QueryForOperationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryForOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryForOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryForOperationResponse) GoString() string {
	return s.String()
}

func (s *QueryForOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryForOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryForOperationResponse) GetBody() *QueryForOperationResponseBody {
	return s.Body
}

func (s *QueryForOperationResponse) SetHeaders(v map[string]*string) *QueryForOperationResponse {
	s.Headers = v
	return s
}

func (s *QueryForOperationResponse) SetStatusCode(v int32) *QueryForOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryForOperationResponse) SetBody(v *QueryForOperationResponseBody) *QueryForOperationResponse {
	s.Body = v
	return s
}

func (s *QueryForOperationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
