// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceInspectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceInspectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceInspectionsResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceInspectionsResponseBody) *GetInstanceInspectionsResponse
	GetBody() *GetInstanceInspectionsResponseBody
}

type GetInstanceInspectionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceInspectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceInspectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionsResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceInspectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceInspectionsResponse) GetBody() *GetInstanceInspectionsResponseBody {
	return s.Body
}

func (s *GetInstanceInspectionsResponse) SetHeaders(v map[string]*string) *GetInstanceInspectionsResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceInspectionsResponse) SetStatusCode(v int32) *GetInstanceInspectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceInspectionsResponse) SetBody(v *GetInstanceInspectionsResponseBody) *GetInstanceInspectionsResponse {
	s.Body = v
	return s
}

func (s *GetInstanceInspectionsResponse) Validate() error {
	return dara.Validate(s)
}
