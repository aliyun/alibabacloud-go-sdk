// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEduProblemSolveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEduProblemSolveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEduProblemSolveResponse
	GetStatusCode() *int32
	SetId(v string) *CreateEduProblemSolveResponse
	GetId() *string
	SetEvent(v string) *CreateEduProblemSolveResponse
	GetEvent() *string
	SetBody(v *CreateEduProblemSolveResponseBody) *CreateEduProblemSolveResponse
	GetBody() *CreateEduProblemSolveResponseBody
}

type CreateEduProblemSolveResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                            `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                            `json:"event,omitempty" xml:"event,omitempty"`
	Body       *CreateEduProblemSolveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEduProblemSolveResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEduProblemSolveResponse) GoString() string {
	return s.String()
}

func (s *CreateEduProblemSolveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEduProblemSolveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEduProblemSolveResponse) GetId() *string {
	return s.Id
}

func (s *CreateEduProblemSolveResponse) GetEvent() *string {
	return s.Event
}

func (s *CreateEduProblemSolveResponse) GetBody() *CreateEduProblemSolveResponseBody {
	return s.Body
}

func (s *CreateEduProblemSolveResponse) SetHeaders(v map[string]*string) *CreateEduProblemSolveResponse {
	s.Headers = v
	return s
}

func (s *CreateEduProblemSolveResponse) SetStatusCode(v int32) *CreateEduProblemSolveResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEduProblemSolveResponse) SetId(v string) *CreateEduProblemSolveResponse {
	s.Id = &v
	return s
}

func (s *CreateEduProblemSolveResponse) SetEvent(v string) *CreateEduProblemSolveResponse {
	s.Event = &v
	return s
}

func (s *CreateEduProblemSolveResponse) SetBody(v *CreateEduProblemSolveResponseBody) *CreateEduProblemSolveResponse {
	s.Body = v
	return s
}

func (s *CreateEduProblemSolveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
