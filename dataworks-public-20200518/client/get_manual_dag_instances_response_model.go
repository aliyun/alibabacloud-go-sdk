// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManualDagInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetManualDagInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetManualDagInstancesResponse
	GetStatusCode() *int32
	SetBody(v *GetManualDagInstancesResponseBody) *GetManualDagInstancesResponse
	GetBody() *GetManualDagInstancesResponseBody
}

type GetManualDagInstancesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetManualDagInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetManualDagInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetManualDagInstancesResponse) GoString() string {
	return s.String()
}

func (s *GetManualDagInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetManualDagInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetManualDagInstancesResponse) GetBody() *GetManualDagInstancesResponseBody {
	return s.Body
}

func (s *GetManualDagInstancesResponse) SetHeaders(v map[string]*string) *GetManualDagInstancesResponse {
	s.Headers = v
	return s
}

func (s *GetManualDagInstancesResponse) SetStatusCode(v int32) *GetManualDagInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetManualDagInstancesResponse) SetBody(v *GetManualDagInstancesResponseBody) *GetManualDagInstancesResponse {
	s.Body = v
	return s
}

func (s *GetManualDagInstancesResponse) Validate() error {
	return dara.Validate(s)
}
