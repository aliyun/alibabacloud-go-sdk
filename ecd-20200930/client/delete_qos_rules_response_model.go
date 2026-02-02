// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQosRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQosRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQosRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQosRulesResponseBody) *DeleteQosRulesResponse
	GetBody() *DeleteQosRulesResponseBody
}

type DeleteQosRulesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQosRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQosRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQosRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteQosRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQosRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQosRulesResponse) GetBody() *DeleteQosRulesResponseBody {
	return s.Body
}

func (s *DeleteQosRulesResponse) SetHeaders(v map[string]*string) *DeleteQosRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteQosRulesResponse) SetStatusCode(v int32) *DeleteQosRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQosRulesResponse) SetBody(v *DeleteQosRulesResponseBody) *DeleteQosRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteQosRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
