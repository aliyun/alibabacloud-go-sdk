// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSummariesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOption(v string) *ListSummariesRequest
	GetOption() *string
}

type ListSummariesRequest struct {
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
}

func (s ListSummariesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSummariesRequest) GoString() string {
	return s.String()
}

func (s *ListSummariesRequest) GetOption() *string {
	return s.Option
}

func (s *ListSummariesRequest) SetOption(v string) *ListSummariesRequest {
	s.Option = &v
	return s
}

func (s *ListSummariesRequest) Validate() error {
	return dara.Validate(s)
}
