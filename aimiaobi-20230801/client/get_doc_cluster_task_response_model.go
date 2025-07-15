// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocClusterTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocClusterTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocClusterTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetDocClusterTaskResponseBody) *GetDocClusterTaskResponse
	GetBody() *GetDocClusterTaskResponseBody
}

type GetDocClusterTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocClusterTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocClusterTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocClusterTaskResponse) GoString() string {
	return s.String()
}

func (s *GetDocClusterTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocClusterTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocClusterTaskResponse) GetBody() *GetDocClusterTaskResponseBody {
	return s.Body
}

func (s *GetDocClusterTaskResponse) SetHeaders(v map[string]*string) *GetDocClusterTaskResponse {
	s.Headers = v
	return s
}

func (s *GetDocClusterTaskResponse) SetStatusCode(v int32) *GetDocClusterTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocClusterTaskResponse) SetBody(v *GetDocClusterTaskResponseBody) *GetDocClusterTaskResponse {
	s.Body = v
	return s
}

func (s *GetDocClusterTaskResponse) Validate() error {
	return dara.Validate(s)
}
