// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceFeaturesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceFeaturesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceFeaturesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceFeaturesResponseBody) *ModifyInstanceFeaturesResponse
	GetBody() *ModifyInstanceFeaturesResponseBody
}

type ModifyInstanceFeaturesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceFeaturesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceFeaturesResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceFeaturesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceFeaturesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceFeaturesResponse) GetBody() *ModifyInstanceFeaturesResponseBody {
	return s.Body
}

func (s *ModifyInstanceFeaturesResponse) SetHeaders(v map[string]*string) *ModifyInstanceFeaturesResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceFeaturesResponse) SetStatusCode(v int32) *ModifyInstanceFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceFeaturesResponse) SetBody(v *ModifyInstanceFeaturesResponseBody) *ModifyInstanceFeaturesResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceFeaturesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
