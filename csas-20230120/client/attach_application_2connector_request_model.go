// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachApplication2ConnectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v []*string) *AttachApplication2ConnectorRequest
	GetApplicationIds() []*string
	SetConnectorId(v string) *AttachApplication2ConnectorRequest
	GetConnectorId() *string
}

type AttachApplication2ConnectorRequest struct {
	// This parameter is required.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// ConnectorID。
	//
	// This parameter is required.
	//
	// example:
	//
	// connector-94db94e06b98****
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
}

func (s AttachApplication2ConnectorRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachApplication2ConnectorRequest) GoString() string {
	return s.String()
}

func (s *AttachApplication2ConnectorRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *AttachApplication2ConnectorRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *AttachApplication2ConnectorRequest) SetApplicationIds(v []*string) *AttachApplication2ConnectorRequest {
	s.ApplicationIds = v
	return s
}

func (s *AttachApplication2ConnectorRequest) SetConnectorId(v string) *AttachApplication2ConnectorRequest {
	s.ConnectorId = &v
	return s
}

func (s *AttachApplication2ConnectorRequest) Validate() error {
	return dara.Validate(s)
}
