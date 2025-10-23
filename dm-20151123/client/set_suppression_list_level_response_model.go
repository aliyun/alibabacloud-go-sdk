// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSuppressionListLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetSuppressionListLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetSuppressionListLevelResponse
	GetStatusCode() *int32
	SetBody(v *SetSuppressionListLevelResponseBody) *SetSuppressionListLevelResponse
	GetBody() *SetSuppressionListLevelResponseBody
}

type SetSuppressionListLevelResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSuppressionListLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSuppressionListLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s SetSuppressionListLevelResponse) GoString() string {
	return s.String()
}

func (s *SetSuppressionListLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetSuppressionListLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetSuppressionListLevelResponse) GetBody() *SetSuppressionListLevelResponseBody {
	return s.Body
}

func (s *SetSuppressionListLevelResponse) SetHeaders(v map[string]*string) *SetSuppressionListLevelResponse {
	s.Headers = v
	return s
}

func (s *SetSuppressionListLevelResponse) SetStatusCode(v int32) *SetSuppressionListLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSuppressionListLevelResponse) SetBody(v *SetSuppressionListLevelResponseBody) *SetSuppressionListLevelResponse {
	s.Body = v
	return s
}

func (s *SetSuppressionListLevelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
