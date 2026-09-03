// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateApiKeyStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *ModelRouterUpdateApiKeyStatusRequest
	GetStatus() *string
}

type ModelRouterUpdateApiKeyStatusRequest struct {
	// The status of the API key. Valid values:
	//
	// - active: The API key is valid.
	//
	// - disabled: The API key is invalid.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ModelRouterUpdateApiKeyStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateApiKeyStatusRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateApiKeyStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *ModelRouterUpdateApiKeyStatusRequest) SetStatus(v string) *ModelRouterUpdateApiKeyStatusRequest {
	s.Status = &v
	return s
}

func (s *ModelRouterUpdateApiKeyStatusRequest) Validate() error {
	return dara.Validate(s)
}
