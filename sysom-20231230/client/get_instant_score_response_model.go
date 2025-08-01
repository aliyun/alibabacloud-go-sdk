// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstantScoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstantScoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstantScoreResponse
	GetStatusCode() *int32
	SetBody(v *GetInstantScoreResponseBody) *GetInstantScoreResponse
	GetBody() *GetInstantScoreResponseBody
}

type GetInstantScoreResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstantScoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstantScoreResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstantScoreResponse) GoString() string {
	return s.String()
}

func (s *GetInstantScoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstantScoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstantScoreResponse) GetBody() *GetInstantScoreResponseBody {
	return s.Body
}

func (s *GetInstantScoreResponse) SetHeaders(v map[string]*string) *GetInstantScoreResponse {
	s.Headers = v
	return s
}

func (s *GetInstantScoreResponse) SetStatusCode(v int32) *GetInstantScoreResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstantScoreResponse) SetBody(v *GetInstantScoreResponseBody) *GetInstantScoreResponse {
	s.Body = v
	return s
}

func (s *GetInstantScoreResponse) Validate() error {
	return dara.Validate(s)
}
