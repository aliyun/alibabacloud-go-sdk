// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowJSONAssestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFlowJSONAssestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFlowJSONAssestResponse
	GetStatusCode() *int32
	SetBody(v *GetFlowJSONAssestResponseBody) *GetFlowJSONAssestResponse
	GetBody() *GetFlowJSONAssestResponseBody
}

type GetFlowJSONAssestResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFlowJSONAssestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFlowJSONAssestResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFlowJSONAssestResponse) GoString() string {
	return s.String()
}

func (s *GetFlowJSONAssestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFlowJSONAssestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFlowJSONAssestResponse) GetBody() *GetFlowJSONAssestResponseBody {
	return s.Body
}

func (s *GetFlowJSONAssestResponse) SetHeaders(v map[string]*string) *GetFlowJSONAssestResponse {
	s.Headers = v
	return s
}

func (s *GetFlowJSONAssestResponse) SetStatusCode(v int32) *GetFlowJSONAssestResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFlowJSONAssestResponse) SetBody(v *GetFlowJSONAssestResponseBody) *GetFlowJSONAssestResponse {
	s.Body = v
	return s
}

func (s *GetFlowJSONAssestResponse) Validate() error {
	return dara.Validate(s)
}
