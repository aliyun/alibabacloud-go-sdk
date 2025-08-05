// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiDestinationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiDestinationName(v string) *DeleteApiDestinationRequest
	GetApiDestinationName() *string
}

type DeleteApiDestinationRequest struct {
	// The name of the API destination.
	//
	// This parameter is required.
	//
	// example:
	//
	// ApiDestinationName
	ApiDestinationName *string `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
}

func (s DeleteApiDestinationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiDestinationRequest) GoString() string {
	return s.String()
}

func (s *DeleteApiDestinationRequest) GetApiDestinationName() *string {
	return s.ApiDestinationName
}

func (s *DeleteApiDestinationRequest) SetApiDestinationName(v string) *DeleteApiDestinationRequest {
	s.ApiDestinationName = &v
	return s
}

func (s *DeleteApiDestinationRequest) Validate() error {
	return dara.Validate(s)
}
