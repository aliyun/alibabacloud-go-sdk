// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawMCPServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribePolarClawMCPServersResponseBody
	GetApplicationId() *string
	SetCode(v int32) *DescribePolarClawMCPServersResponseBody
	GetCode() *int32
	SetMessage(v string) *DescribePolarClawMCPServersResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePolarClawMCPServersResponseBody
	GetRequestId() *string
	SetServers(v map[string]interface{}) *DescribePolarClawMCPServersResponseBody
	GetServers() map[string]interface{}
}

type DescribePolarClawMCPServersResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2281C6C9-CBAB-1AFD-8400-670750CF6025_2212
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {
	//
	//     "dev-mcp-server": "{\\"args\\":[\\"-y\\",\\"@polarclaw/mcp-dev\\"],\\"command\\":\\"node\\"}"
	//
	// }
	Servers map[string]interface{} `json:"Servers,omitempty" xml:"Servers,omitempty"`
}

func (s DescribePolarClawMCPServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawMCPServersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarClawMCPServersResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawMCPServersResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribePolarClawMCPServersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePolarClawMCPServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarClawMCPServersResponseBody) GetServers() map[string]interface{} {
	return s.Servers
}

func (s *DescribePolarClawMCPServersResponseBody) SetApplicationId(v string) *DescribePolarClawMCPServersResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawMCPServersResponseBody) SetCode(v int32) *DescribePolarClawMCPServersResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePolarClawMCPServersResponseBody) SetMessage(v string) *DescribePolarClawMCPServersResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePolarClawMCPServersResponseBody) SetRequestId(v string) *DescribePolarClawMCPServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarClawMCPServersResponseBody) SetServers(v map[string]interface{}) *DescribePolarClawMCPServersResponseBody {
	s.Servers = v
	return s
}

func (s *DescribePolarClawMCPServersResponseBody) Validate() error {
	return dara.Validate(s)
}
