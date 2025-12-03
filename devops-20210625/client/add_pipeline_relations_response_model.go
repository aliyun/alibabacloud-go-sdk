// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPipelineRelationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPipelineRelationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPipelineRelationsResponse
	GetStatusCode() *int32
	SetBody(v *AddPipelineRelationsResponseBody) *AddPipelineRelationsResponse
	GetBody() *AddPipelineRelationsResponseBody
}

type AddPipelineRelationsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPipelineRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPipelineRelationsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPipelineRelationsResponse) GoString() string {
	return s.String()
}

func (s *AddPipelineRelationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPipelineRelationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPipelineRelationsResponse) GetBody() *AddPipelineRelationsResponseBody {
	return s.Body
}

func (s *AddPipelineRelationsResponse) SetHeaders(v map[string]*string) *AddPipelineRelationsResponse {
	s.Headers = v
	return s
}

func (s *AddPipelineRelationsResponse) SetStatusCode(v int32) *AddPipelineRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPipelineRelationsResponse) SetBody(v *AddPipelineRelationsResponseBody) *AddPipelineRelationsResponse {
	s.Body = v
	return s
}

func (s *AddPipelineRelationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
