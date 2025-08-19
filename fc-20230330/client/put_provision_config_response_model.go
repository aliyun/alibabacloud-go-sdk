// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutProvisionConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutProvisionConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutProvisionConfigResponse
	GetStatusCode() *int32
	SetBody(v *ProvisionConfig) *PutProvisionConfigResponse
	GetBody() *ProvisionConfig
}

type PutProvisionConfigResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProvisionConfig   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutProvisionConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s PutProvisionConfigResponse) GoString() string {
	return s.String()
}

func (s *PutProvisionConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutProvisionConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutProvisionConfigResponse) GetBody() *ProvisionConfig {
	return s.Body
}

func (s *PutProvisionConfigResponse) SetHeaders(v map[string]*string) *PutProvisionConfigResponse {
	s.Headers = v
	return s
}

func (s *PutProvisionConfigResponse) SetStatusCode(v int32) *PutProvisionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PutProvisionConfigResponse) SetBody(v *ProvisionConfig) *PutProvisionConfigResponse {
	s.Body = v
	return s
}

func (s *PutProvisionConfigResponse) Validate() error {
	return dara.Validate(s)
}
