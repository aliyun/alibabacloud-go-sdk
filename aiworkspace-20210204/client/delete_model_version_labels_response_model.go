// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelVersionLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteModelVersionLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteModelVersionLabelsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteModelVersionLabelsResponseBody) *DeleteModelVersionLabelsResponse
	GetBody() *DeleteModelVersionLabelsResponseBody
}

type DeleteModelVersionLabelsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelVersionLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelVersionLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelVersionLabelsResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelVersionLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteModelVersionLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteModelVersionLabelsResponse) GetBody() *DeleteModelVersionLabelsResponseBody {
	return s.Body
}

func (s *DeleteModelVersionLabelsResponse) SetHeaders(v map[string]*string) *DeleteModelVersionLabelsResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelVersionLabelsResponse) SetStatusCode(v int32) *DeleteModelVersionLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelVersionLabelsResponse) SetBody(v *DeleteModelVersionLabelsResponseBody) *DeleteModelVersionLabelsResponse {
	s.Body = v
	return s
}

func (s *DeleteModelVersionLabelsResponse) Validate() error {
	return dara.Validate(s)
}
