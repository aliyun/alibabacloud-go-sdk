// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQosPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFederationId(v string) *DeleteQosPolicyRequest
	GetFederationId() *string
	SetFileSystemId(v string) *DeleteQosPolicyRequest
	GetFileSystemId() *string
	SetInputRegionId(v string) *DeleteQosPolicyRequest
	GetInputRegionId() *string
	SetQosPolicyId(v string) *DeleteQosPolicyRequest
	GetQosPolicyId() *string
}

type DeleteQosPolicyRequest struct {
	FederationId *string `json:"FederationId,omitempty" xml:"FederationId,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// This parameter is required.
	QosPolicyId *string `json:"QosPolicyId,omitempty" xml:"QosPolicyId,omitempty"`
}

func (s DeleteQosPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQosPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteQosPolicyRequest) GetFederationId() *string {
	return s.FederationId
}

func (s *DeleteQosPolicyRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteQosPolicyRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *DeleteQosPolicyRequest) GetQosPolicyId() *string {
	return s.QosPolicyId
}

func (s *DeleteQosPolicyRequest) SetFederationId(v string) *DeleteQosPolicyRequest {
	s.FederationId = &v
	return s
}

func (s *DeleteQosPolicyRequest) SetFileSystemId(v string) *DeleteQosPolicyRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteQosPolicyRequest) SetInputRegionId(v string) *DeleteQosPolicyRequest {
	s.InputRegionId = &v
	return s
}

func (s *DeleteQosPolicyRequest) SetQosPolicyId(v string) *DeleteQosPolicyRequest {
	s.QosPolicyId = &v
	return s
}

func (s *DeleteQosPolicyRequest) Validate() error {
	return dara.Validate(s)
}
