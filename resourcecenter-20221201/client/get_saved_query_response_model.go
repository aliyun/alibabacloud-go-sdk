// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSavedQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSavedQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSavedQueryResponse
	GetStatusCode() *int32
	SetBody(v *GetSavedQueryResponseBody) *GetSavedQueryResponse
	GetBody() *GetSavedQueryResponseBody
}

type GetSavedQueryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSavedQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSavedQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSavedQueryResponse) GoString() string {
	return s.String()
}

func (s *GetSavedQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSavedQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSavedQueryResponse) GetBody() *GetSavedQueryResponseBody {
	return s.Body
}

func (s *GetSavedQueryResponse) SetHeaders(v map[string]*string) *GetSavedQueryResponse {
	s.Headers = v
	return s
}

func (s *GetSavedQueryResponse) SetStatusCode(v int32) *GetSavedQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSavedQueryResponse) SetBody(v *GetSavedQueryResponseBody) *GetSavedQueryResponse {
	s.Body = v
	return s
}

func (s *GetSavedQueryResponse) Validate() error {
	return dara.Validate(s)
}
