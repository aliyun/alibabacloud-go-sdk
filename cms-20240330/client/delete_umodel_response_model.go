// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUmodelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUmodelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUmodelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUmodelResponseBody) *DeleteUmodelResponse
	GetBody() *DeleteUmodelResponseBody
}

type DeleteUmodelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUmodelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUmodelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUmodelResponse) GoString() string {
	return s.String()
}

func (s *DeleteUmodelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUmodelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUmodelResponse) GetBody() *DeleteUmodelResponseBody {
	return s.Body
}

func (s *DeleteUmodelResponse) SetHeaders(v map[string]*string) *DeleteUmodelResponse {
	s.Headers = v
	return s
}

func (s *DeleteUmodelResponse) SetStatusCode(v int32) *DeleteUmodelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUmodelResponse) SetBody(v *DeleteUmodelResponseBody) *DeleteUmodelResponse {
	s.Body = v
	return s
}

func (s *DeleteUmodelResponse) Validate() error {
	return dara.Validate(s)
}
