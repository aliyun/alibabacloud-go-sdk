// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateShrinkNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*ValidateShrinkNodesRequestBody) *ValidateShrinkNodesRequest
	GetBody() []*ValidateShrinkNodesRequestBody
	SetCount(v int32) *ValidateShrinkNodesRequest
	GetCount() *int32
	SetIgnoreStatus(v bool) *ValidateShrinkNodesRequest
	GetIgnoreStatus() *bool
	SetNodeType(v string) *ValidateShrinkNodesRequest
	GetNodeType() *string
}

type ValidateShrinkNodesRequest struct {
	Body []*ValidateShrinkNodesRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// false
	IgnoreStatus *bool `json:"ignoreStatus,omitempty" xml:"ignoreStatus,omitempty"`
	// Returned results:
	//
	// 	- true: can be scaled in
	//
	// 	- false: cannot be scaled in.
	//
	// This parameter is required.
	//
	// example:
	//
	// WORKER
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
}

func (s ValidateShrinkNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateShrinkNodesRequest) GoString() string {
	return s.String()
}

func (s *ValidateShrinkNodesRequest) GetBody() []*ValidateShrinkNodesRequestBody {
	return s.Body
}

func (s *ValidateShrinkNodesRequest) GetCount() *int32 {
	return s.Count
}

func (s *ValidateShrinkNodesRequest) GetIgnoreStatus() *bool {
	return s.IgnoreStatus
}

func (s *ValidateShrinkNodesRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *ValidateShrinkNodesRequest) SetBody(v []*ValidateShrinkNodesRequestBody) *ValidateShrinkNodesRequest {
	s.Body = v
	return s
}

func (s *ValidateShrinkNodesRequest) SetCount(v int32) *ValidateShrinkNodesRequest {
	s.Count = &v
	return s
}

func (s *ValidateShrinkNodesRequest) SetIgnoreStatus(v bool) *ValidateShrinkNodesRequest {
	s.IgnoreStatus = &v
	return s
}

func (s *ValidateShrinkNodesRequest) SetNodeType(v string) *ValidateShrinkNodesRequest {
	s.NodeType = &v
	return s
}

func (s *ValidateShrinkNodesRequest) Validate() error {
	return dara.Validate(s)
}

type ValidateShrinkNodesRequestBody struct {
	// example:
	//
	// 192.168.xx.xx
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// example:
	//
	// es-cn-pl32xxxxxxx-data-f-1
	HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty"`
	// example:
	//
	// WORKER
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
	// example:
	//
	// 9200
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// cn-shanghai-c
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ValidateShrinkNodesRequestBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateShrinkNodesRequestBody) GoString() string {
	return s.String()
}

func (s *ValidateShrinkNodesRequestBody) GetHost() *string {
	return s.Host
}

func (s *ValidateShrinkNodesRequestBody) GetHostName() *string {
	return s.HostName
}

func (s *ValidateShrinkNodesRequestBody) GetNodeType() *string {
	return s.NodeType
}

func (s *ValidateShrinkNodesRequestBody) GetPort() *int32 {
	return s.Port
}

func (s *ValidateShrinkNodesRequestBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *ValidateShrinkNodesRequestBody) SetHost(v string) *ValidateShrinkNodesRequestBody {
	s.Host = &v
	return s
}

func (s *ValidateShrinkNodesRequestBody) SetHostName(v string) *ValidateShrinkNodesRequestBody {
	s.HostName = &v
	return s
}

func (s *ValidateShrinkNodesRequestBody) SetNodeType(v string) *ValidateShrinkNodesRequestBody {
	s.NodeType = &v
	return s
}

func (s *ValidateShrinkNodesRequestBody) SetPort(v int32) *ValidateShrinkNodesRequestBody {
	s.Port = &v
	return s
}

func (s *ValidateShrinkNodesRequestBody) SetZoneId(v string) *ValidateShrinkNodesRequestBody {
	s.ZoneId = &v
	return s
}

func (s *ValidateShrinkNodesRequestBody) Validate() error {
	return dara.Validate(s)
}
