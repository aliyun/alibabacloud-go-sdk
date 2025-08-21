// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnsSaleConditionControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEnsSaleConditionControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEnsSaleConditionControlResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEnsSaleConditionControlResponseBody) *DeleteEnsSaleConditionControlResponse
	GetBody() *DeleteEnsSaleConditionControlResponseBody
}

type DeleteEnsSaleConditionControlResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEnsSaleConditionControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEnsSaleConditionControlResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnsSaleConditionControlResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnsSaleConditionControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEnsSaleConditionControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEnsSaleConditionControlResponse) GetBody() *DeleteEnsSaleConditionControlResponseBody {
	return s.Body
}

func (s *DeleteEnsSaleConditionControlResponse) SetHeaders(v map[string]*string) *DeleteEnsSaleConditionControlResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnsSaleConditionControlResponse) SetStatusCode(v int32) *DeleteEnsSaleConditionControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEnsSaleConditionControlResponse) SetBody(v *DeleteEnsSaleConditionControlResponseBody) *DeleteEnsSaleConditionControlResponse {
	s.Body = v
	return s
}

func (s *DeleteEnsSaleConditionControlResponse) Validate() error {
	return dara.Validate(s)
}
