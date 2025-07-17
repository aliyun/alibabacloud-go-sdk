// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *ConfigAppResponseBody
	GetData() *string
	SetRequestId(v string) *ConfigAppResponseBody
	GetRequestId() *string
}

type ConfigAppResponseBody struct {
	// The result of turning on or off the main switch of the agent or the main switch status of the agent. Indicates whether the request was successful. Valid values:
	//
	// 	- success
	//
	// 	- failed
	//
	// The main switch status of the agent. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// abc@def success\\nghi@jkl success\\n
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16AF921B-8187-489F-9913-43C808B4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigAppResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigAppResponseBody) GetData() *string {
	return s.Data
}

func (s *ConfigAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigAppResponseBody) SetData(v string) *ConfigAppResponseBody {
	s.Data = &v
	return s
}

func (s *ConfigAppResponseBody) SetRequestId(v string) *ConfigAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigAppResponseBody) Validate() error {
	return dara.Validate(s)
}
