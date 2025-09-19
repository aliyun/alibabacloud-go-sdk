// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContainerScanTaskByAppNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateContainerScanTaskByAppNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateContainerScanTaskByAppNameResponse
	GetStatusCode() *int32
	SetBody(v *CreateContainerScanTaskByAppNameResponseBody) *CreateContainerScanTaskByAppNameResponse
	GetBody() *CreateContainerScanTaskByAppNameResponseBody
}

type CreateContainerScanTaskByAppNameResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContainerScanTaskByAppNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContainerScanTaskByAppNameResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateContainerScanTaskByAppNameResponse) GoString() string {
	return s.String()
}

func (s *CreateContainerScanTaskByAppNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateContainerScanTaskByAppNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateContainerScanTaskByAppNameResponse) GetBody() *CreateContainerScanTaskByAppNameResponseBody {
	return s.Body
}

func (s *CreateContainerScanTaskByAppNameResponse) SetHeaders(v map[string]*string) *CreateContainerScanTaskByAppNameResponse {
	s.Headers = v
	return s
}

func (s *CreateContainerScanTaskByAppNameResponse) SetStatusCode(v int32) *CreateContainerScanTaskByAppNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContainerScanTaskByAppNameResponse) SetBody(v *CreateContainerScanTaskByAppNameResponseBody) *CreateContainerScanTaskByAppNameResponse {
	s.Body = v
	return s
}

func (s *CreateContainerScanTaskByAppNameResponse) Validate() error {
	return dara.Validate(s)
}
