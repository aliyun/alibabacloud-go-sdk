// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySiteFeaturesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySiteFeaturesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySiteFeaturesResponse
	GetStatusCode() *int32
	SetBody(v *ModifySiteFeaturesResponseBody) *ModifySiteFeaturesResponse
	GetBody() *ModifySiteFeaturesResponseBody
}

type ModifySiteFeaturesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySiteFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySiteFeaturesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySiteFeaturesResponse) GoString() string {
	return s.String()
}

func (s *ModifySiteFeaturesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySiteFeaturesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySiteFeaturesResponse) GetBody() *ModifySiteFeaturesResponseBody {
	return s.Body
}

func (s *ModifySiteFeaturesResponse) SetHeaders(v map[string]*string) *ModifySiteFeaturesResponse {
	s.Headers = v
	return s
}

func (s *ModifySiteFeaturesResponse) SetStatusCode(v int32) *ModifySiteFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySiteFeaturesResponse) SetBody(v *ModifySiteFeaturesResponseBody) *ModifySiteFeaturesResponse {
	s.Body = v
	return s
}

func (s *ModifySiteFeaturesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
