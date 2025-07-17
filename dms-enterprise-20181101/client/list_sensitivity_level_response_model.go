// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSensitivityLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSensitivityLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSensitivityLevelResponse
	GetStatusCode() *int32
	SetBody(v *ListSensitivityLevelResponseBody) *ListSensitivityLevelResponse
	GetBody() *ListSensitivityLevelResponseBody
}

type ListSensitivityLevelResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSensitivityLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSensitivityLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSensitivityLevelResponse) GoString() string {
	return s.String()
}

func (s *ListSensitivityLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSensitivityLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSensitivityLevelResponse) GetBody() *ListSensitivityLevelResponseBody {
	return s.Body
}

func (s *ListSensitivityLevelResponse) SetHeaders(v map[string]*string) *ListSensitivityLevelResponse {
	s.Headers = v
	return s
}

func (s *ListSensitivityLevelResponse) SetStatusCode(v int32) *ListSensitivityLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSensitivityLevelResponse) SetBody(v *ListSensitivityLevelResponseBody) *ListSensitivityLevelResponse {
	s.Body = v
	return s
}

func (s *ListSensitivityLevelResponse) Validate() error {
	return dara.Validate(s)
}
