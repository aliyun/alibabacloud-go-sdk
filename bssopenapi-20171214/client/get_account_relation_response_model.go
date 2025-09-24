// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccountRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccountRelationResponse
	GetStatusCode() *int32
	SetBody(v *GetAccountRelationResponseBody) *GetAccountRelationResponse
	GetBody() *GetAccountRelationResponseBody
}

type GetAccountRelationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccountRelationResponse) GoString() string {
	return s.String()
}

func (s *GetAccountRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccountRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccountRelationResponse) GetBody() *GetAccountRelationResponseBody {
	return s.Body
}

func (s *GetAccountRelationResponse) SetHeaders(v map[string]*string) *GetAccountRelationResponse {
	s.Headers = v
	return s
}

func (s *GetAccountRelationResponse) SetStatusCode(v int32) *GetAccountRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountRelationResponse) SetBody(v *GetAccountRelationResponseBody) *GetAccountRelationResponse {
	s.Body = v
	return s
}

func (s *GetAccountRelationResponse) Validate() error {
	return dara.Validate(s)
}
