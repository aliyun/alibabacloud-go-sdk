// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceConnectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSericeConnectionType(v string) *ListServiceConnectionsRequest
	GetSericeConnectionType() *string
}

type ListServiceConnectionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// codeup
	SericeConnectionType *string `json:"sericeConnectionType,omitempty" xml:"sericeConnectionType,omitempty"`
}

func (s ListServiceConnectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceConnectionsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceConnectionsRequest) GetSericeConnectionType() *string {
	return s.SericeConnectionType
}

func (s *ListServiceConnectionsRequest) SetSericeConnectionType(v string) *ListServiceConnectionsRequest {
	s.SericeConnectionType = &v
	return s
}

func (s *ListServiceConnectionsRequest) Validate() error {
	return dara.Validate(s)
}
