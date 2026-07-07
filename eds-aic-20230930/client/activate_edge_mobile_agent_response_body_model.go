// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateEdgeMobileAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ActivateEdgeMobileAgentResponseBodyData) *ActivateEdgeMobileAgentResponseBody
	GetData() *ActivateEdgeMobileAgentResponseBodyData
	SetRequestId(v string) *ActivateEdgeMobileAgentResponseBody
	GetRequestId() *string
}

type ActivateEdgeMobileAgentResponseBody struct {
	// The response data object.
	Data *ActivateEdgeMobileAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5C5CEF0A-D6E1-58D3-8750-67DB4F82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActivateEdgeMobileAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActivateEdgeMobileAgentResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateEdgeMobileAgentResponseBody) GetData() *ActivateEdgeMobileAgentResponseBodyData {
	return s.Data
}

func (s *ActivateEdgeMobileAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActivateEdgeMobileAgentResponseBody) SetData(v *ActivateEdgeMobileAgentResponseBodyData) *ActivateEdgeMobileAgentResponseBody {
	s.Data = v
	return s
}

func (s *ActivateEdgeMobileAgentResponseBody) SetRequestId(v string) *ActivateEdgeMobileAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateEdgeMobileAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ActivateEdgeMobileAgentResponseBodyData struct {
	// The assigned API key. The plaintext value is returned only upon the first activation.
	//
	// example:
	//
	// cpk-81vfd8t8zdfxdf*****
	AuthToken *string `json:"AuthToken,omitempty" xml:"AuthToken,omitempty"`
	// The device ID.
	//
	// example:
	//
	// sn-0001eevqa6jeapl*****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// Indicates whether the request is an idempotent duplicate request.
	//
	// example:
	//
	// false
	Idempotent *bool `json:"Idempotent,omitempty" xml:"Idempotent,omitempty"`
	// The EdgeMobile instance ID.
	//
	// example:
	//
	// em-uto81vfd8t8z****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ActivateEdgeMobileAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ActivateEdgeMobileAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *ActivateEdgeMobileAgentResponseBodyData) GetAuthToken() *string {
	return s.AuthToken
}

func (s *ActivateEdgeMobileAgentResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ActivateEdgeMobileAgentResponseBodyData) GetIdempotent() *bool {
	return s.Idempotent
}

func (s *ActivateEdgeMobileAgentResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ActivateEdgeMobileAgentResponseBodyData) SetAuthToken(v string) *ActivateEdgeMobileAgentResponseBodyData {
	s.AuthToken = &v
	return s
}

func (s *ActivateEdgeMobileAgentResponseBodyData) SetDeviceId(v string) *ActivateEdgeMobileAgentResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *ActivateEdgeMobileAgentResponseBodyData) SetIdempotent(v bool) *ActivateEdgeMobileAgentResponseBodyData {
	s.Idempotent = &v
	return s
}

func (s *ActivateEdgeMobileAgentResponseBodyData) SetInstanceId(v string) *ActivateEdgeMobileAgentResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ActivateEdgeMobileAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
