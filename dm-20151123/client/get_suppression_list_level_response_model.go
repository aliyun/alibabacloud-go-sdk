// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuppressionListLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSuppressionListLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSuppressionListLevelResponse
	GetStatusCode() *int32
	SetBody(v *GetSuppressionListLevelResponseBody) *GetSuppressionListLevelResponse
	GetBody() *GetSuppressionListLevelResponseBody
}

type GetSuppressionListLevelResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSuppressionListLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSuppressionListLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSuppressionListLevelResponse) GoString() string {
	return s.String()
}

func (s *GetSuppressionListLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSuppressionListLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSuppressionListLevelResponse) GetBody() *GetSuppressionListLevelResponseBody {
	return s.Body
}

func (s *GetSuppressionListLevelResponse) SetHeaders(v map[string]*string) *GetSuppressionListLevelResponse {
	s.Headers = v
	return s
}

func (s *GetSuppressionListLevelResponse) SetStatusCode(v int32) *GetSuppressionListLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSuppressionListLevelResponse) SetBody(v *GetSuppressionListLevelResponseBody) *GetSuppressionListLevelResponse {
	s.Body = v
	return s
}

func (s *GetSuppressionListLevelResponse) Validate() error {
	return dara.Validate(s)
}
