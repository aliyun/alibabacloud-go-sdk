// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUrgentDemandItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUrgentDemandItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUrgentDemandItemResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUrgentDemandItemResponseBody) *DeleteUrgentDemandItemResponse
	GetBody() *DeleteUrgentDemandItemResponseBody
}

type DeleteUrgentDemandItemResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUrgentDemandItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUrgentDemandItemResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUrgentDemandItemResponse) GoString() string {
	return s.String()
}

func (s *DeleteUrgentDemandItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUrgentDemandItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUrgentDemandItemResponse) GetBody() *DeleteUrgentDemandItemResponseBody {
	return s.Body
}

func (s *DeleteUrgentDemandItemResponse) SetHeaders(v map[string]*string) *DeleteUrgentDemandItemResponse {
	s.Headers = v
	return s
}

func (s *DeleteUrgentDemandItemResponse) SetStatusCode(v int32) *DeleteUrgentDemandItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUrgentDemandItemResponse) SetBody(v *DeleteUrgentDemandItemResponseBody) *DeleteUrgentDemandItemResponse {
	s.Body = v
	return s
}

func (s *DeleteUrgentDemandItemResponse) Validate() error {
	return dara.Validate(s)
}
