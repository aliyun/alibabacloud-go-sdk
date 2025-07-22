// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateForwardEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExternalIp(v string) *CreateForwardEntryRequest
	GetExternalIp() *string
	SetExternalPort(v string) *CreateForwardEntryRequest
	GetExternalPort() *string
	SetForwardEntryName(v string) *CreateForwardEntryRequest
	GetForwardEntryName() *string
	SetForwardTableId(v string) *CreateForwardEntryRequest
	GetForwardTableId() *string
	SetInternalIp(v string) *CreateForwardEntryRequest
	GetInternalIp() *string
	SetInternalPort(v string) *CreateForwardEntryRequest
	GetInternalPort() *string
	SetIpProtocol(v string) *CreateForwardEntryRequest
	GetIpProtocol() *string
	SetRegionId(v string) *CreateForwardEntryRequest
	GetRegionId() *string
}

type CreateForwardEntryRequest struct {
	// This parameter is required.
	ExternalIp *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	// This parameter is required.
	ExternalPort     *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	ForwardEntryName *string `json:"ForwardEntryName,omitempty" xml:"ForwardEntryName,omitempty"`
	// This parameter is required.
	ForwardTableId *string `json:"ForwardTableId,omitempty" xml:"ForwardTableId,omitempty"`
	// This parameter is required.
	InternalIp *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	// This parameter is required.
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	// This parameter is required.
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateForwardEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateForwardEntryRequest) GetExternalIp() *string {
	return s.ExternalIp
}

func (s *CreateForwardEntryRequest) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *CreateForwardEntryRequest) GetForwardEntryName() *string {
	return s.ForwardEntryName
}

func (s *CreateForwardEntryRequest) GetForwardTableId() *string {
	return s.ForwardTableId
}

func (s *CreateForwardEntryRequest) GetInternalIp() *string {
	return s.InternalIp
}

func (s *CreateForwardEntryRequest) GetInternalPort() *string {
	return s.InternalPort
}

func (s *CreateForwardEntryRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *CreateForwardEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateForwardEntryRequest) SetExternalIp(v string) *CreateForwardEntryRequest {
	s.ExternalIp = &v
	return s
}

func (s *CreateForwardEntryRequest) SetExternalPort(v string) *CreateForwardEntryRequest {
	s.ExternalPort = &v
	return s
}

func (s *CreateForwardEntryRequest) SetForwardEntryName(v string) *CreateForwardEntryRequest {
	s.ForwardEntryName = &v
	return s
}

func (s *CreateForwardEntryRequest) SetForwardTableId(v string) *CreateForwardEntryRequest {
	s.ForwardTableId = &v
	return s
}

func (s *CreateForwardEntryRequest) SetInternalIp(v string) *CreateForwardEntryRequest {
	s.InternalIp = &v
	return s
}

func (s *CreateForwardEntryRequest) SetInternalPort(v string) *CreateForwardEntryRequest {
	s.InternalPort = &v
	return s
}

func (s *CreateForwardEntryRequest) SetIpProtocol(v string) *CreateForwardEntryRequest {
	s.IpProtocol = &v
	return s
}

func (s *CreateForwardEntryRequest) SetRegionId(v string) *CreateForwardEntryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateForwardEntryRequest) Validate() error {
	return dara.Validate(s)
}
