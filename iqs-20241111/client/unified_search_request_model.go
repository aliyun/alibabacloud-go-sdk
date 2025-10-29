// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnifiedSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UnifiedSearchInput) *UnifiedSearchRequest
	GetBody() *UnifiedSearchInput
}

type UnifiedSearchRequest struct {
	Body *UnifiedSearchInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnifiedSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s UnifiedSearchRequest) GoString() string {
	return s.String()
}

func (s *UnifiedSearchRequest) GetBody() *UnifiedSearchInput {
	return s.Body
}

func (s *UnifiedSearchRequest) SetBody(v *UnifiedSearchInput) *UnifiedSearchRequest {
	s.Body = v
	return s
}

func (s *UnifiedSearchRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
