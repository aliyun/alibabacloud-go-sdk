// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2InstanceForTerraformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLindormV2InstanceForTerraformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLindormV2InstanceForTerraformResponse
	GetStatusCode() *int32
	SetBody(v *GetLindormV2InstanceForTerraformResponseBody) *GetLindormV2InstanceForTerraformResponse
	GetBody() *GetLindormV2InstanceForTerraformResponseBody
}

type GetLindormV2InstanceForTerraformResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLindormV2InstanceForTerraformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLindormV2InstanceForTerraformResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceForTerraformResponse) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceForTerraformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLindormV2InstanceForTerraformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLindormV2InstanceForTerraformResponse) GetBody() *GetLindormV2InstanceForTerraformResponseBody {
	return s.Body
}

func (s *GetLindormV2InstanceForTerraformResponse) SetHeaders(v map[string]*string) *GetLindormV2InstanceForTerraformResponse {
	s.Headers = v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponse) SetStatusCode(v int32) *GetLindormV2InstanceForTerraformResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponse) SetBody(v *GetLindormV2InstanceForTerraformResponseBody) *GetLindormV2InstanceForTerraformResponse {
	s.Body = v
	return s
}

func (s *GetLindormV2InstanceForTerraformResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
