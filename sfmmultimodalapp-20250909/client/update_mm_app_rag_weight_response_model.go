// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppRagWeightResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMmAppRagWeightResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMmAppRagWeightResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMmAppRagWeightResponseBody) *UpdateMmAppRagWeightResponse
	GetBody() *UpdateMmAppRagWeightResponseBody
}

type UpdateMmAppRagWeightResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMmAppRagWeightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMmAppRagWeightResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppRagWeightResponse) GoString() string {
	return s.String()
}

func (s *UpdateMmAppRagWeightResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMmAppRagWeightResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMmAppRagWeightResponse) GetBody() *UpdateMmAppRagWeightResponseBody {
	return s.Body
}

func (s *UpdateMmAppRagWeightResponse) SetHeaders(v map[string]*string) *UpdateMmAppRagWeightResponse {
	s.Headers = v
	return s
}

func (s *UpdateMmAppRagWeightResponse) SetStatusCode(v int32) *UpdateMmAppRagWeightResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMmAppRagWeightResponse) SetBody(v *UpdateMmAppRagWeightResponseBody) *UpdateMmAppRagWeightResponse {
	s.Body = v
	return s
}

func (s *UpdateMmAppRagWeightResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
