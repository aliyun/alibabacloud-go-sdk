// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsHiveJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApsHiveJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApsHiveJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateApsHiveJobResponseBody) *CreateApsHiveJobResponse
	GetBody() *CreateApsHiveJobResponseBody
}

type CreateApsHiveJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApsHiveJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApsHiveJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApsHiveJobResponse) GoString() string {
	return s.String()
}

func (s *CreateApsHiveJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApsHiveJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApsHiveJobResponse) GetBody() *CreateApsHiveJobResponseBody {
	return s.Body
}

func (s *CreateApsHiveJobResponse) SetHeaders(v map[string]*string) *CreateApsHiveJobResponse {
	s.Headers = v
	return s
}

func (s *CreateApsHiveJobResponse) SetStatusCode(v int32) *CreateApsHiveJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApsHiveJobResponse) SetBody(v *CreateApsHiveJobResponseBody) *CreateApsHiveJobResponse {
	s.Body = v
	return s
}

func (s *CreateApsHiveJobResponse) Validate() error {
	return dara.Validate(s)
}
