// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskInstanceRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskInstanceRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskInstanceRelationResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskInstanceRelationResponseBody) *GetTaskInstanceRelationResponse
	GetBody() *GetTaskInstanceRelationResponseBody
}

type GetTaskInstanceRelationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskInstanceRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskInstanceRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceRelationResponse) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskInstanceRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskInstanceRelationResponse) GetBody() *GetTaskInstanceRelationResponseBody {
	return s.Body
}

func (s *GetTaskInstanceRelationResponse) SetHeaders(v map[string]*string) *GetTaskInstanceRelationResponse {
	s.Headers = v
	return s
}

func (s *GetTaskInstanceRelationResponse) SetStatusCode(v int32) *GetTaskInstanceRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskInstanceRelationResponse) SetBody(v *GetTaskInstanceRelationResponseBody) *GetTaskInstanceRelationResponse {
	s.Body = v
	return s
}

func (s *GetTaskInstanceRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
