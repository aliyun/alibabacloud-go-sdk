// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLindormV2InstanceParameterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLindormV2InstanceParameterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLindormV2InstanceParameterResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLindormV2InstanceParameterResponseBody) *UpdateLindormV2InstanceParameterResponse
	GetBody() *UpdateLindormV2InstanceParameterResponseBody
}

type UpdateLindormV2InstanceParameterResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLindormV2InstanceParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLindormV2InstanceParameterResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLindormV2InstanceParameterResponse) GoString() string {
	return s.String()
}

func (s *UpdateLindormV2InstanceParameterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLindormV2InstanceParameterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLindormV2InstanceParameterResponse) GetBody() *UpdateLindormV2InstanceParameterResponseBody {
	return s.Body
}

func (s *UpdateLindormV2InstanceParameterResponse) SetHeaders(v map[string]*string) *UpdateLindormV2InstanceParameterResponse {
	s.Headers = v
	return s
}

func (s *UpdateLindormV2InstanceParameterResponse) SetStatusCode(v int32) *UpdateLindormV2InstanceParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterResponse) SetBody(v *UpdateLindormV2InstanceParameterResponseBody) *UpdateLindormV2InstanceParameterResponse {
	s.Body = v
	return s
}

func (s *UpdateLindormV2InstanceParameterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
