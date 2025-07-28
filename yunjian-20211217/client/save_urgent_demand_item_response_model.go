// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveUrgentDemandItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveUrgentDemandItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveUrgentDemandItemResponse
	GetStatusCode() *int32
	SetBody(v *SaveUrgentDemandItemResponseBody) *SaveUrgentDemandItemResponse
	GetBody() *SaveUrgentDemandItemResponseBody
}

type SaveUrgentDemandItemResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveUrgentDemandItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveUrgentDemandItemResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveUrgentDemandItemResponse) GoString() string {
	return s.String()
}

func (s *SaveUrgentDemandItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveUrgentDemandItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveUrgentDemandItemResponse) GetBody() *SaveUrgentDemandItemResponseBody {
	return s.Body
}

func (s *SaveUrgentDemandItemResponse) SetHeaders(v map[string]*string) *SaveUrgentDemandItemResponse {
	s.Headers = v
	return s
}

func (s *SaveUrgentDemandItemResponse) SetStatusCode(v int32) *SaveUrgentDemandItemResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveUrgentDemandItemResponse) SetBody(v *SaveUrgentDemandItemResponseBody) *SaveUrgentDemandItemResponse {
	s.Body = v
	return s
}

func (s *SaveUrgentDemandItemResponse) Validate() error {
	return dara.Validate(s)
}
