// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelCardsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProvider(v string) *ListModelCardsRequest
	GetProvider() *string
}

type ListModelCardsRequest struct {
	// The identifier of the model provider used to filter model cards by provider.
	//
	// example:
	//
	// provider-xxxxxxxx
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s ListModelCardsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelCardsRequest) GoString() string {
	return s.String()
}

func (s *ListModelCardsRequest) GetProvider() *string {
	return s.Provider
}

func (s *ListModelCardsRequest) SetProvider(v string) *ListModelCardsRequest {
	s.Provider = &v
	return s
}

func (s *ListModelCardsRequest) Validate() error {
	return dara.Validate(s)
}
