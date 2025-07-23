// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUmodelDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUmodelDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUmodelDataResponse
	GetStatusCode() *int32
	SetBody(v *GetUmodelDataResponseBody) *GetUmodelDataResponse
	GetBody() *GetUmodelDataResponseBody
}

type GetUmodelDataResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUmodelDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUmodelDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUmodelDataResponse) GoString() string {
	return s.String()
}

func (s *GetUmodelDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUmodelDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUmodelDataResponse) GetBody() *GetUmodelDataResponseBody {
	return s.Body
}

func (s *GetUmodelDataResponse) SetHeaders(v map[string]*string) *GetUmodelDataResponse {
	s.Headers = v
	return s
}

func (s *GetUmodelDataResponse) SetStatusCode(v int32) *GetUmodelDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUmodelDataResponse) SetBody(v *GetUmodelDataResponseBody) *GetUmodelDataResponse {
	s.Body = v
	return s
}

func (s *GetUmodelDataResponse) Validate() error {
	return dara.Validate(s)
}
