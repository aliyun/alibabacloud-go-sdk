// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubCrowdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSubCrowdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSubCrowdResponse
	GetStatusCode() *int32
	SetBody(v *GetSubCrowdResponseBody) *GetSubCrowdResponse
	GetBody() *GetSubCrowdResponseBody
}

type GetSubCrowdResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubCrowdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubCrowdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSubCrowdResponse) GoString() string {
	return s.String()
}

func (s *GetSubCrowdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSubCrowdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSubCrowdResponse) GetBody() *GetSubCrowdResponseBody {
	return s.Body
}

func (s *GetSubCrowdResponse) SetHeaders(v map[string]*string) *GetSubCrowdResponse {
	s.Headers = v
	return s
}

func (s *GetSubCrowdResponse) SetStatusCode(v int32) *GetSubCrowdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubCrowdResponse) SetBody(v *GetSubCrowdResponseBody) *GetSubCrowdResponse {
	s.Body = v
	return s
}

func (s *GetSubCrowdResponse) Validate() error {
	return dara.Validate(s)
}
