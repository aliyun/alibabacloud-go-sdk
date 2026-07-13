// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProtectionPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateProtectionPolicyResponseBodyData) *CreateProtectionPolicyResponseBody
	GetData() *CreateProtectionPolicyResponseBodyData
	SetRequestId(v string) *CreateProtectionPolicyResponseBody
	GetRequestId() *string
}

type CreateProtectionPolicyResponseBody struct {
	// The data returned.
	Data *CreateProtectionPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique ID of the request.
	//
	// example:
	//
	// 34081B20-C4C0-514F-93F6-8EEC3D1A587E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProtectionPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectionPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProtectionPolicyResponseBody) GetData() *CreateProtectionPolicyResponseBodyData {
	return s.Data
}

func (s *CreateProtectionPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProtectionPolicyResponseBody) SetData(v *CreateProtectionPolicyResponseBodyData) *CreateProtectionPolicyResponseBody {
	s.Data = v
	return s
}

func (s *CreateProtectionPolicyResponseBody) SetRequestId(v string) *CreateProtectionPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProtectionPolicyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProtectionPolicyResponseBodyData struct {
	// The ID of the protection policy.
	//
	// example:
	//
	// p-123***7890
	ProtectionPolicyId *string `json:"ProtectionPolicyId,omitempty" xml:"ProtectionPolicyId,omitempty"`
}

func (s CreateProtectionPolicyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectionPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProtectionPolicyResponseBodyData) GetProtectionPolicyId() *string {
	return s.ProtectionPolicyId
}

func (s *CreateProtectionPolicyResponseBodyData) SetProtectionPolicyId(v string) *CreateProtectionPolicyResponseBodyData {
	s.ProtectionPolicyId = &v
	return s
}

func (s *CreateProtectionPolicyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
