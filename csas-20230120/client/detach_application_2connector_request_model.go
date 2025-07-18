// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachApplication2ConnectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v []*string) *DetachApplication2ConnectorRequest
	GetApplicationIds() []*string
	SetConnectorId(v string) *DetachApplication2ConnectorRequest
	GetConnectorId() *string
}

type DetachApplication2ConnectorRequest struct {
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

func (s DetachApplication2ConnectorRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachApplication2ConnectorRequest) GoString() string {
	return s.String()
}

func (s *DetachApplication2ConnectorRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *DetachApplication2ConnectorRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *DetachApplication2ConnectorRequest) SetApplicationIds(v []*string) *DetachApplication2ConnectorRequest {
	s.ApplicationIds = v
	return s
}

func (s *DetachApplication2ConnectorRequest) SetConnectorId(v string) *DetachApplication2ConnectorRequest {
	s.ConnectorId = &v
	return s
}

func (s *DetachApplication2ConnectorRequest) Validate() error {
	return dara.Validate(s)
}
