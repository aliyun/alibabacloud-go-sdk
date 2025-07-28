// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridCloudGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHybridCloudGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHybridCloudGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateHybridCloudGroupResponseBody) *CreateHybridCloudGroupResponse
	GetBody() *CreateHybridCloudGroupResponseBody
}

type CreateHybridCloudGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHybridCloudGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHybridCloudGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridCloudGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateHybridCloudGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHybridCloudGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHybridCloudGroupResponse) GetBody() *CreateHybridCloudGroupResponseBody {
	return s.Body
}

func (s *CreateHybridCloudGroupResponse) SetHeaders(v map[string]*string) *CreateHybridCloudGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateHybridCloudGroupResponse) SetStatusCode(v int32) *CreateHybridCloudGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHybridCloudGroupResponse) SetBody(v *CreateHybridCloudGroupResponseBody) *CreateHybridCloudGroupResponse {
	s.Body = v
	return s
}

func (s *CreateHybridCloudGroupResponse) Validate() error {
	return dara.Validate(s)
}
