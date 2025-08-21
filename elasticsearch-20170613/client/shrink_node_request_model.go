// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShrinkNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*ShrinkNodeRequestBody) *ShrinkNodeRequest
	GetBody() []*ShrinkNodeRequestBody
	SetClientToken(v string) *ShrinkNodeRequest
	GetClientToken() *string
	SetCount(v int32) *ShrinkNodeRequest
	GetCount() *int32
	SetIgnoreStatus(v bool) *ShrinkNodeRequest
	GetIgnoreStatus() *bool
	SetNodeType(v string) *ShrinkNodeRequest
	GetNodeType() *string
}

type ShrinkNodeRequest struct {
	Body []*ShrinkNodeRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// 2
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// example:
	//
	// false
	IgnoreStatus *bool `json:"ignoreStatus,omitempty" xml:"ignoreStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WORKER
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
}

func (s ShrinkNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s ShrinkNodeRequest) GoString() string {
	return s.String()
}

func (s *ShrinkNodeRequest) GetBody() []*ShrinkNodeRequestBody {
	return s.Body
}

func (s *ShrinkNodeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ShrinkNodeRequest) GetCount() *int32 {
	return s.Count
}

func (s *ShrinkNodeRequest) GetIgnoreStatus() *bool {
	return s.IgnoreStatus
}

func (s *ShrinkNodeRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *ShrinkNodeRequest) SetBody(v []*ShrinkNodeRequestBody) *ShrinkNodeRequest {
	s.Body = v
	return s
}

func (s *ShrinkNodeRequest) SetClientToken(v string) *ShrinkNodeRequest {
	s.ClientToken = &v
	return s
}

func (s *ShrinkNodeRequest) SetCount(v int32) *ShrinkNodeRequest {
	s.Count = &v
	return s
}

func (s *ShrinkNodeRequest) SetIgnoreStatus(v bool) *ShrinkNodeRequest {
	s.IgnoreStatus = &v
	return s
}

func (s *ShrinkNodeRequest) SetNodeType(v string) *ShrinkNodeRequest {
	s.NodeType = &v
	return s
}

func (s *ShrinkNodeRequest) Validate() error {
	return dara.Validate(s)
}

type ShrinkNodeRequestBody struct {
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

func (s ShrinkNodeRequestBody) String() string {
	return dara.Prettify(s)
}

func (s ShrinkNodeRequestBody) GoString() string {
	return s.String()
}

func (s *ShrinkNodeRequestBody) GetHost() *string {
	return s.Host
}

func (s *ShrinkNodeRequestBody) GetHostName() *string {
	return s.HostName
}

func (s *ShrinkNodeRequestBody) GetNodeType() *string {
	return s.NodeType
}

func (s *ShrinkNodeRequestBody) GetPort() *int32 {
	return s.Port
}

func (s *ShrinkNodeRequestBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *ShrinkNodeRequestBody) SetHost(v string) *ShrinkNodeRequestBody {
	s.Host = &v
	return s
}

func (s *ShrinkNodeRequestBody) SetHostName(v string) *ShrinkNodeRequestBody {
	s.HostName = &v
	return s
}

func (s *ShrinkNodeRequestBody) SetNodeType(v string) *ShrinkNodeRequestBody {
	s.NodeType = &v
	return s
}

func (s *ShrinkNodeRequestBody) SetPort(v int32) *ShrinkNodeRequestBody {
	s.Port = &v
	return s
}

func (s *ShrinkNodeRequestBody) SetZoneId(v string) *ShrinkNodeRequestBody {
	s.ZoneId = &v
	return s
}

func (s *ShrinkNodeRequestBody) Validate() error {
	return dara.Validate(s)
}
