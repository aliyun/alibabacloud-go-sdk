// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeployApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeployApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeployApisResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeployApisResponseBody) *BatchDeployApisResponse
	GetBody() *BatchDeployApisResponseBody
}

type BatchDeployApisResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeployApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeployApisResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeployApisResponse) GoString() string {
	return s.String()
}

func (s *BatchDeployApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeployApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeployApisResponse) GetBody() *BatchDeployApisResponseBody {
	return s.Body
}

func (s *BatchDeployApisResponse) SetHeaders(v map[string]*string) *BatchDeployApisResponse {
	s.Headers = v
	return s
}

func (s *BatchDeployApisResponse) SetStatusCode(v int32) *BatchDeployApisResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeployApisResponse) SetBody(v *BatchDeployApisResponseBody) *BatchDeployApisResponse {
	s.Body = v
	return s
}

func (s *BatchDeployApisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
