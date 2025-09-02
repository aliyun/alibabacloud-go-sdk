// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBaselineConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBaselineConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBaselineConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetBaselineConfigResponseBody) *GetBaselineConfigResponse
	GetBody() *GetBaselineConfigResponseBody
}

type GetBaselineConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBaselineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBaselineConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineConfigResponse) GoString() string {
	return s.String()
}

func (s *GetBaselineConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBaselineConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBaselineConfigResponse) GetBody() *GetBaselineConfigResponseBody {
	return s.Body
}

func (s *GetBaselineConfigResponse) SetHeaders(v map[string]*string) *GetBaselineConfigResponse {
	s.Headers = v
	return s
}

func (s *GetBaselineConfigResponse) SetStatusCode(v int32) *GetBaselineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBaselineConfigResponse) SetBody(v *GetBaselineConfigResponseBody) *GetBaselineConfigResponse {
	s.Body = v
	return s
}

func (s *GetBaselineConfigResponse) Validate() error {
	return dara.Validate(s)
}
