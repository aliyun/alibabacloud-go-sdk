// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDISyncTaskConfigForUpdatingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateDISyncTaskConfigForUpdatingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateDISyncTaskConfigForUpdatingResponse
	GetStatusCode() *int32
	SetBody(v *GenerateDISyncTaskConfigForUpdatingResponseBody) *GenerateDISyncTaskConfigForUpdatingResponse
	GetBody() *GenerateDISyncTaskConfigForUpdatingResponseBody
}

type GenerateDISyncTaskConfigForUpdatingResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateDISyncTaskConfigForUpdatingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateDISyncTaskConfigForUpdatingResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateDISyncTaskConfigForUpdatingResponse) GoString() string {
	return s.String()
}

func (s *GenerateDISyncTaskConfigForUpdatingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateDISyncTaskConfigForUpdatingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateDISyncTaskConfigForUpdatingResponse) GetBody() *GenerateDISyncTaskConfigForUpdatingResponseBody {
	return s.Body
}

func (s *GenerateDISyncTaskConfigForUpdatingResponse) SetHeaders(v map[string]*string) *GenerateDISyncTaskConfigForUpdatingResponse {
	s.Headers = v
	return s
}

func (s *GenerateDISyncTaskConfigForUpdatingResponse) SetStatusCode(v int32) *GenerateDISyncTaskConfigForUpdatingResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDISyncTaskConfigForUpdatingResponse) SetBody(v *GenerateDISyncTaskConfigForUpdatingResponseBody) *GenerateDISyncTaskConfigForUpdatingResponse {
	s.Body = v
	return s
}

func (s *GenerateDISyncTaskConfigForUpdatingResponse) Validate() error {
	return dara.Validate(s)
}
