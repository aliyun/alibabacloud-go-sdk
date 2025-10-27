// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConcernNecessityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyConcernNecessityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyConcernNecessityResponse
	GetStatusCode() *int32
	SetBody(v *ModifyConcernNecessityResponseBody) *ModifyConcernNecessityResponse
	GetBody() *ModifyConcernNecessityResponseBody
}

type ModifyConcernNecessityResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyConcernNecessityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyConcernNecessityResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyConcernNecessityResponse) GoString() string {
	return s.String()
}

func (s *ModifyConcernNecessityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyConcernNecessityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyConcernNecessityResponse) GetBody() *ModifyConcernNecessityResponseBody {
	return s.Body
}

func (s *ModifyConcernNecessityResponse) SetHeaders(v map[string]*string) *ModifyConcernNecessityResponse {
	s.Headers = v
	return s
}

func (s *ModifyConcernNecessityResponse) SetStatusCode(v int32) *ModifyConcernNecessityResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyConcernNecessityResponse) SetBody(v *ModifyConcernNecessityResponseBody) *ModifyConcernNecessityResponse {
	s.Body = v
	return s
}

func (s *ModifyConcernNecessityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
