// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDISyncTaskConfigForCreatingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateDISyncTaskConfigForCreatingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateDISyncTaskConfigForCreatingResponse
	GetStatusCode() *int32
	SetBody(v *GenerateDISyncTaskConfigForCreatingResponseBody) *GenerateDISyncTaskConfigForCreatingResponse
	GetBody() *GenerateDISyncTaskConfigForCreatingResponseBody
}

type GenerateDISyncTaskConfigForCreatingResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateDISyncTaskConfigForCreatingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateDISyncTaskConfigForCreatingResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateDISyncTaskConfigForCreatingResponse) GoString() string {
	return s.String()
}

func (s *GenerateDISyncTaskConfigForCreatingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateDISyncTaskConfigForCreatingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateDISyncTaskConfigForCreatingResponse) GetBody() *GenerateDISyncTaskConfigForCreatingResponseBody {
	return s.Body
}

func (s *GenerateDISyncTaskConfigForCreatingResponse) SetHeaders(v map[string]*string) *GenerateDISyncTaskConfigForCreatingResponse {
	s.Headers = v
	return s
}

func (s *GenerateDISyncTaskConfigForCreatingResponse) SetStatusCode(v int32) *GenerateDISyncTaskConfigForCreatingResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDISyncTaskConfigForCreatingResponse) SetBody(v *GenerateDISyncTaskConfigForCreatingResponseBody) *GenerateDISyncTaskConfigForCreatingResponse {
	s.Body = v
	return s
}

func (s *GenerateDISyncTaskConfigForCreatingResponse) Validate() error {
	return dara.Validate(s)
}
