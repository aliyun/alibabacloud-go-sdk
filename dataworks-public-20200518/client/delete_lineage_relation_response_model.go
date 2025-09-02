// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLineageRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLineageRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLineageRelationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLineageRelationResponseBody) *DeleteLineageRelationResponse
	GetBody() *DeleteLineageRelationResponseBody
}

type DeleteLineageRelationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLineageRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLineageRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLineageRelationResponse) GoString() string {
	return s.String()
}

func (s *DeleteLineageRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLineageRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLineageRelationResponse) GetBody() *DeleteLineageRelationResponseBody {
	return s.Body
}

func (s *DeleteLineageRelationResponse) SetHeaders(v map[string]*string) *DeleteLineageRelationResponse {
	s.Headers = v
	return s
}

func (s *DeleteLineageRelationResponse) SetStatusCode(v int32) *DeleteLineageRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLineageRelationResponse) SetBody(v *DeleteLineageRelationResponseBody) *DeleteLineageRelationResponse {
	s.Body = v
	return s
}

func (s *DeleteLineageRelationResponse) Validate() error {
	return dara.Validate(s)
}
