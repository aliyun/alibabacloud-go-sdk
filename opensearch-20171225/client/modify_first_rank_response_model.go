// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFirstRankResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyFirstRankResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyFirstRankResponse
	GetStatusCode() *int32
	SetBody(v *ModifyFirstRankResponseBody) *ModifyFirstRankResponse
	GetBody() *ModifyFirstRankResponseBody
}

type ModifyFirstRankResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFirstRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFirstRankResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyFirstRankResponse) GoString() string {
	return s.String()
}

func (s *ModifyFirstRankResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyFirstRankResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyFirstRankResponse) GetBody() *ModifyFirstRankResponseBody {
	return s.Body
}

func (s *ModifyFirstRankResponse) SetHeaders(v map[string]*string) *ModifyFirstRankResponse {
	s.Headers = v
	return s
}

func (s *ModifyFirstRankResponse) SetStatusCode(v int32) *ModifyFirstRankResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFirstRankResponse) SetBody(v *ModifyFirstRankResponseBody) *ModifyFirstRankResponse {
	s.Body = v
	return s
}

func (s *ModifyFirstRankResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
