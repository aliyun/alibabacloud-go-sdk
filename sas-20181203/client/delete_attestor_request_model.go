// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAttestorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteAttestorRequest
	GetName() *string
	SetResourceOwnerId(v int64) *DeleteAttestorRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DeleteAttestorRequest
	GetSourceIp() *string
}

type DeleteAttestorRequest struct {
	// The name of the witness.
	//
	// example:
	//
	// attestor
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 221.214.XXX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteAttestorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAttestorRequest) GoString() string {
	return s.String()
}

func (s *DeleteAttestorRequest) GetName() *string {
	return s.Name
}

func (s *DeleteAttestorRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteAttestorRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DeleteAttestorRequest) SetName(v string) *DeleteAttestorRequest {
	s.Name = &v
	return s
}

func (s *DeleteAttestorRequest) SetResourceOwnerId(v int64) *DeleteAttestorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAttestorRequest) SetSourceIp(v string) *DeleteAttestorRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteAttestorRequest) Validate() error {
	return dara.Validate(s)
}
