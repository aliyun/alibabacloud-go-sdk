// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecuritySuggestionNumberRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetSecuritySuggestionNumberRequest struct {
}

func (s GetSecuritySuggestionNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecuritySuggestionNumberRequest) GoString() string {
	return s.String()
}

func (s *GetSecuritySuggestionNumberRequest) Validate() error {
	return dara.Validate(s)
}
