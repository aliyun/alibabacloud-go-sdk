// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineWithAssetsCodeVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRoutineWithAssetsCodeVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRoutineWithAssetsCodeVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateRoutineWithAssetsCodeVersionResponseBody) *CreateRoutineWithAssetsCodeVersionResponse
	GetBody() *CreateRoutineWithAssetsCodeVersionResponseBody
}

type CreateRoutineWithAssetsCodeVersionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRoutineWithAssetsCodeVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRoutineWithAssetsCodeVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineWithAssetsCodeVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateRoutineWithAssetsCodeVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRoutineWithAssetsCodeVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRoutineWithAssetsCodeVersionResponse) GetBody() *CreateRoutineWithAssetsCodeVersionResponseBody {
	return s.Body
}

func (s *CreateRoutineWithAssetsCodeVersionResponse) SetHeaders(v map[string]*string) *CreateRoutineWithAssetsCodeVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionResponse) SetStatusCode(v int32) *CreateRoutineWithAssetsCodeVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionResponse) SetBody(v *CreateRoutineWithAssetsCodeVersionResponseBody) *CreateRoutineWithAssetsCodeVersionResponse {
	s.Body = v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionResponse) Validate() error {
	return dara.Validate(s)
}
