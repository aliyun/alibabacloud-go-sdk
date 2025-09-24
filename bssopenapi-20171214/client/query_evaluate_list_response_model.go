// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEvaluateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEvaluateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEvaluateListResponse
	GetStatusCode() *int32
	SetBody(v *QueryEvaluateListResponseBody) *QueryEvaluateListResponse
	GetBody() *QueryEvaluateListResponseBody
}

type QueryEvaluateListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEvaluateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEvaluateListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEvaluateListResponse) GoString() string {
	return s.String()
}

func (s *QueryEvaluateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEvaluateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEvaluateListResponse) GetBody() *QueryEvaluateListResponseBody {
	return s.Body
}

func (s *QueryEvaluateListResponse) SetHeaders(v map[string]*string) *QueryEvaluateListResponse {
	s.Headers = v
	return s
}

func (s *QueryEvaluateListResponse) SetStatusCode(v int32) *QueryEvaluateListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEvaluateListResponse) SetBody(v *QueryEvaluateListResponseBody) *QueryEvaluateListResponse {
	s.Body = v
	return s
}

func (s *QueryEvaluateListResponse) Validate() error {
	return dara.Validate(s)
}
