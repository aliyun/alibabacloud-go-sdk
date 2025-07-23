// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUmodelDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUmodelDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUmodelDataResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUmodelDataResponseBody) *DeleteUmodelDataResponse
	GetBody() *DeleteUmodelDataResponseBody
}

type DeleteUmodelDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUmodelDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUmodelDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUmodelDataResponse) GoString() string {
	return s.String()
}

func (s *DeleteUmodelDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUmodelDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUmodelDataResponse) GetBody() *DeleteUmodelDataResponseBody {
	return s.Body
}

func (s *DeleteUmodelDataResponse) SetHeaders(v map[string]*string) *DeleteUmodelDataResponse {
	s.Headers = v
	return s
}

func (s *DeleteUmodelDataResponse) SetStatusCode(v int32) *DeleteUmodelDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUmodelDataResponse) SetBody(v *DeleteUmodelDataResponseBody) *DeleteUmodelDataResponse {
	s.Body = v
	return s
}

func (s *DeleteUmodelDataResponse) Validate() error {
	return dara.Validate(s)
}
