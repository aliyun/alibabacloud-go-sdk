// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceAuthsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceAuthType(v string) *ListServiceAuthsRequest
	GetServiceAuthType() *string
}

type ListServiceAuthsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// RAM
	ServiceAuthType *string `json:"serviceAuthType,omitempty" xml:"serviceAuthType,omitempty"`
}

func (s ListServiceAuthsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceAuthsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceAuthsRequest) GetServiceAuthType() *string {
	return s.ServiceAuthType
}

func (s *ListServiceAuthsRequest) SetServiceAuthType(v string) *ListServiceAuthsRequest {
	s.ServiceAuthType = &v
	return s
}

func (s *ListServiceAuthsRequest) Validate() error {
	return dara.Validate(s)
}
