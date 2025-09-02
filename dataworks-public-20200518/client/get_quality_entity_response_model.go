// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityEntityResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityEntityResponseBody) *GetQualityEntityResponse
	GetBody() *GetQualityEntityResponseBody
}

type GetQualityEntityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityEntityResponse) GoString() string {
	return s.String()
}

func (s *GetQualityEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityEntityResponse) GetBody() *GetQualityEntityResponseBody {
	return s.Body
}

func (s *GetQualityEntityResponse) SetHeaders(v map[string]*string) *GetQualityEntityResponse {
	s.Headers = v
	return s
}

func (s *GetQualityEntityResponse) SetStatusCode(v int32) *GetQualityEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityEntityResponse) SetBody(v *GetQualityEntityResponseBody) *GetQualityEntityResponse {
	s.Body = v
	return s
}

func (s *GetQualityEntityResponse) Validate() error {
	return dara.Validate(s)
}
