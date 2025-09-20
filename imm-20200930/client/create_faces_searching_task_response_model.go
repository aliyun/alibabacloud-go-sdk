// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFacesSearchingTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFacesSearchingTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFacesSearchingTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateFacesSearchingTaskResponseBody) *CreateFacesSearchingTaskResponse
	GetBody() *CreateFacesSearchingTaskResponseBody
}

type CreateFacesSearchingTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFacesSearchingTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFacesSearchingTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFacesSearchingTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateFacesSearchingTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFacesSearchingTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFacesSearchingTaskResponse) GetBody() *CreateFacesSearchingTaskResponseBody {
	return s.Body
}

func (s *CreateFacesSearchingTaskResponse) SetHeaders(v map[string]*string) *CreateFacesSearchingTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateFacesSearchingTaskResponse) SetStatusCode(v int32) *CreateFacesSearchingTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFacesSearchingTaskResponse) SetBody(v *CreateFacesSearchingTaskResponseBody) *CreateFacesSearchingTaskResponse {
	s.Body = v
	return s
}

func (s *CreateFacesSearchingTaskResponse) Validate() error {
	return dara.Validate(s)
}
