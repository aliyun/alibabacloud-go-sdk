// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenAckServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *OpenAckServiceRequest
	GetType() *string
}

type OpenAckServiceRequest struct {
	// The type of service that you want to activate. Valid values:
	//
	// 	- `propayasgo`: ACK clusters (including ACK managed clusters and ACK dedicated clusters), ACK Serverless clusters, and registered clusters.
	//
	// 	- `edgepayasgo`: ACK Edge clusters.
	//
	// example:
	//
	// propayasgo
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s OpenAckServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenAckServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenAckServiceRequest) GetType() *string {
	return s.Type
}

func (s *OpenAckServiceRequest) SetType(v string) *OpenAckServiceRequest {
	s.Type = &v
	return s
}

func (s *OpenAckServiceRequest) Validate() error {
	return dara.Validate(s)
}
