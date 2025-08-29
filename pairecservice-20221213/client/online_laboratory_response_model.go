// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineLaboratoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnlineLaboratoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnlineLaboratoryResponse
	GetStatusCode() *int32
	SetBody(v *OnlineLaboratoryResponseBody) *OnlineLaboratoryResponse
	GetBody() *OnlineLaboratoryResponseBody
}

type OnlineLaboratoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnlineLaboratoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnlineLaboratoryResponse) String() string {
	return dara.Prettify(s)
}

func (s OnlineLaboratoryResponse) GoString() string {
	return s.String()
}

func (s *OnlineLaboratoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnlineLaboratoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnlineLaboratoryResponse) GetBody() *OnlineLaboratoryResponseBody {
	return s.Body
}

func (s *OnlineLaboratoryResponse) SetHeaders(v map[string]*string) *OnlineLaboratoryResponse {
	s.Headers = v
	return s
}

func (s *OnlineLaboratoryResponse) SetStatusCode(v int32) *OnlineLaboratoryResponse {
	s.StatusCode = &v
	return s
}

func (s *OnlineLaboratoryResponse) SetBody(v *OnlineLaboratoryResponseBody) *OnlineLaboratoryResponse {
	s.Body = v
	return s
}

func (s *OnlineLaboratoryResponse) Validate() error {
	return dara.Validate(s)
}
