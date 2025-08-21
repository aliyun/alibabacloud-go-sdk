// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateTransferableNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*ValidateTransferableNodesRequestBody) *ValidateTransferableNodesRequest
	GetBody() []*ValidateTransferableNodesRequestBody
	SetNodeType(v string) *ValidateTransferableNodesRequest
	GetNodeType() *string
}

type ValidateTransferableNodesRequest struct {
	Body []*ValidateTransferableNodesRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// WORKER
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
}

func (s ValidateTransferableNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateTransferableNodesRequest) GoString() string {
	return s.String()
}

func (s *ValidateTransferableNodesRequest) GetBody() []*ValidateTransferableNodesRequestBody {
	return s.Body
}

func (s *ValidateTransferableNodesRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *ValidateTransferableNodesRequest) SetBody(v []*ValidateTransferableNodesRequestBody) *ValidateTransferableNodesRequest {
	s.Body = v
	return s
}

func (s *ValidateTransferableNodesRequest) SetNodeType(v string) *ValidateTransferableNodesRequest {
	s.NodeType = &v
	return s
}

func (s *ValidateTransferableNodesRequest) Validate() error {
	return dara.Validate(s)
}

type ValidateTransferableNodesRequestBody struct {
	// example:
	//
	// 172.16.xx.xx
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// example:
	//
	// 9200
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// cn-shanghai-c
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ValidateTransferableNodesRequestBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateTransferableNodesRequestBody) GoString() string {
	return s.String()
}

func (s *ValidateTransferableNodesRequestBody) GetHost() *string {
	return s.Host
}

func (s *ValidateTransferableNodesRequestBody) GetPort() *int32 {
	return s.Port
}

func (s *ValidateTransferableNodesRequestBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *ValidateTransferableNodesRequestBody) SetHost(v string) *ValidateTransferableNodesRequestBody {
	s.Host = &v
	return s
}

func (s *ValidateTransferableNodesRequestBody) SetPort(v int32) *ValidateTransferableNodesRequestBody {
	s.Port = &v
	return s
}

func (s *ValidateTransferableNodesRequestBody) SetZoneId(v string) *ValidateTransferableNodesRequestBody {
	s.ZoneId = &v
	return s
}

func (s *ValidateTransferableNodesRequestBody) Validate() error {
	return dara.Validate(s)
}
