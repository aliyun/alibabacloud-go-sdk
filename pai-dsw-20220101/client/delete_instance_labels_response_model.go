// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInstanceLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInstanceLabelsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInstanceLabelsResponseBody) *DeleteInstanceLabelsResponse
	GetBody() *DeleteInstanceLabelsResponseBody
}

type DeleteInstanceLabelsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceLabelsResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInstanceLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInstanceLabelsResponse) GetBody() *DeleteInstanceLabelsResponseBody {
	return s.Body
}

func (s *DeleteInstanceLabelsResponse) SetHeaders(v map[string]*string) *DeleteInstanceLabelsResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceLabelsResponse) SetStatusCode(v int32) *DeleteInstanceLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceLabelsResponse) SetBody(v *DeleteInstanceLabelsResponseBody) *DeleteInstanceLabelsResponse {
	s.Body = v
	return s
}

func (s *DeleteInstanceLabelsResponse) Validate() error {
	return dara.Validate(s)
}
