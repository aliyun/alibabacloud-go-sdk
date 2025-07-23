// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertUmodelDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpsertUmodelDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpsertUmodelDataResponse
	GetStatusCode() *int32
	SetBody(v *UpsertUmodelDataResponseBody) *UpsertUmodelDataResponse
	GetBody() *UpsertUmodelDataResponseBody
}

type UpsertUmodelDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpsertUmodelDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpsertUmodelDataResponse) String() string {
	return dara.Prettify(s)
}

func (s UpsertUmodelDataResponse) GoString() string {
	return s.String()
}

func (s *UpsertUmodelDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpsertUmodelDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpsertUmodelDataResponse) GetBody() *UpsertUmodelDataResponseBody {
	return s.Body
}

func (s *UpsertUmodelDataResponse) SetHeaders(v map[string]*string) *UpsertUmodelDataResponse {
	s.Headers = v
	return s
}

func (s *UpsertUmodelDataResponse) SetStatusCode(v int32) *UpsertUmodelDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpsertUmodelDataResponse) SetBody(v *UpsertUmodelDataResponseBody) *UpsertUmodelDataResponse {
	s.Body = v
	return s
}

func (s *UpsertUmodelDataResponse) Validate() error {
	return dara.Validate(s)
}
