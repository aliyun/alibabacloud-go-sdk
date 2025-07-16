// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFormListInAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFormListInAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFormListInAppResponse
	GetStatusCode() *int32
	SetBody(v *GetFormListInAppResponseBody) *GetFormListInAppResponse
	GetBody() *GetFormListInAppResponseBody
}

type GetFormListInAppResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFormListInAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFormListInAppResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFormListInAppResponse) GoString() string {
	return s.String()
}

func (s *GetFormListInAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFormListInAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFormListInAppResponse) GetBody() *GetFormListInAppResponseBody {
	return s.Body
}

func (s *GetFormListInAppResponse) SetHeaders(v map[string]*string) *GetFormListInAppResponse {
	s.Headers = v
	return s
}

func (s *GetFormListInAppResponse) SetStatusCode(v int32) *GetFormListInAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFormListInAppResponse) SetBody(v *GetFormListInAppResponseBody) *GetFormListInAppResponse {
	s.Body = v
	return s
}

func (s *GetFormListInAppResponse) Validate() error {
	return dara.Validate(s)
}
