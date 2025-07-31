// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApiImpactsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataServiceApiImpactsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataServiceApiImpactsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataServiceApiImpactsResponseBody) *ListDataServiceApiImpactsResponse
	GetBody() *ListDataServiceApiImpactsResponseBody
}

type ListDataServiceApiImpactsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataServiceApiImpactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataServiceApiImpactsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiImpactsResponse) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiImpactsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataServiceApiImpactsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataServiceApiImpactsResponse) GetBody() *ListDataServiceApiImpactsResponseBody {
	return s.Body
}

func (s *ListDataServiceApiImpactsResponse) SetHeaders(v map[string]*string) *ListDataServiceApiImpactsResponse {
	s.Headers = v
	return s
}

func (s *ListDataServiceApiImpactsResponse) SetStatusCode(v int32) *ListDataServiceApiImpactsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataServiceApiImpactsResponse) SetBody(v *ListDataServiceApiImpactsResponseBody) *ListDataServiceApiImpactsResponse {
	s.Body = v
	return s
}

func (s *ListDataServiceApiImpactsResponse) Validate() error {
	return dara.Validate(s)
}
