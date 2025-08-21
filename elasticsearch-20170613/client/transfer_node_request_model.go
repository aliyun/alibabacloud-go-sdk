// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*TransferNodeRequestBody) *TransferNodeRequest
	GetBody() []*TransferNodeRequestBody
	SetClientToken(v string) *TransferNodeRequest
	GetClientToken() *string
	SetNodeType(v string) *TransferNodeRequest
	GetNodeType() *string
}

type TransferNodeRequest struct {
	Body []*TransferNodeRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WORKER
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
}

func (s TransferNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferNodeRequest) GoString() string {
	return s.String()
}

func (s *TransferNodeRequest) GetBody() []*TransferNodeRequestBody {
	return s.Body
}

func (s *TransferNodeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *TransferNodeRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *TransferNodeRequest) SetBody(v []*TransferNodeRequestBody) *TransferNodeRequest {
	s.Body = v
	return s
}

func (s *TransferNodeRequest) SetClientToken(v string) *TransferNodeRequest {
	s.ClientToken = &v
	return s
}

func (s *TransferNodeRequest) SetNodeType(v string) *TransferNodeRequest {
	s.NodeType = &v
	return s
}

func (s *TransferNodeRequest) Validate() error {
	return dara.Validate(s)
}

type TransferNodeRequestBody struct {
	// example:
	//
	// 192.168.xx.xx
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

func (s TransferNodeRequestBody) String() string {
	return dara.Prettify(s)
}

func (s TransferNodeRequestBody) GoString() string {
	return s.String()
}

func (s *TransferNodeRequestBody) GetHost() *string {
	return s.Host
}

func (s *TransferNodeRequestBody) GetPort() *int32 {
	return s.Port
}

func (s *TransferNodeRequestBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *TransferNodeRequestBody) SetHost(v string) *TransferNodeRequestBody {
	s.Host = &v
	return s
}

func (s *TransferNodeRequestBody) SetPort(v int32) *TransferNodeRequestBody {
	s.Port = &v
	return s
}

func (s *TransferNodeRequestBody) SetZoneId(v string) *TransferNodeRequestBody {
	s.ZoneId = &v
	return s
}

func (s *TransferNodeRequestBody) Validate() error {
	return dara.Validate(s)
}
