// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQualityRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQualityRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQualityRulesResponseBody) *DeleteQualityRulesResponse
	GetBody() *DeleteQualityRulesResponseBody
}

type DeleteQualityRulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQualityRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQualityRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteQualityRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQualityRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQualityRulesResponse) GetBody() *DeleteQualityRulesResponseBody {
	return s.Body
}

func (s *DeleteQualityRulesResponse) SetHeaders(v map[string]*string) *DeleteQualityRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteQualityRulesResponse) SetStatusCode(v int32) *DeleteQualityRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQualityRulesResponse) SetBody(v *DeleteQualityRulesResponseBody) *DeleteQualityRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteQualityRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
