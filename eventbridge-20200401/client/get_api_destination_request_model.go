// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiDestinationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiDestinationName(v string) *GetApiDestinationRequest
	GetApiDestinationName() *string
}

type GetApiDestinationRequest struct {
	// The name of the API destination.
	//
	// This parameter is required.
	//
	// example:
	//
	// api-destination-name
	ApiDestinationName *string `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
}

func (s GetApiDestinationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApiDestinationRequest) GoString() string {
	return s.String()
}

func (s *GetApiDestinationRequest) GetApiDestinationName() *string {
	return s.ApiDestinationName
}

func (s *GetApiDestinationRequest) SetApiDestinationName(v string) *GetApiDestinationRequest {
	s.ApiDestinationName = &v
	return s
}

func (s *GetApiDestinationRequest) Validate() error {
	return dara.Validate(s)
}
