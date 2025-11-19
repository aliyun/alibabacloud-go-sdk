// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityIPGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSecurityIPGroupResponseBody
	GetCode() *string
	SetData(v *DeleteSecurityIPGroupResponseBodyData) *DeleteSecurityIPGroupResponseBody
	GetData() *DeleteSecurityIPGroupResponseBodyData
	SetMessage(v string) *DeleteSecurityIPGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSecurityIPGroupResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteSecurityIPGroupResponseBody
	GetSuccess() *string
}

type DeleteSecurityIPGroupResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// ListResult<InstanceSSL>
	Data *DeleteSecurityIPGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSecurityIPGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityIPGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityIPGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSecurityIPGroupResponseBody) GetData() *DeleteSecurityIPGroupResponseBodyData {
	return s.Data
}

func (s *DeleteSecurityIPGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSecurityIPGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecurityIPGroupResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteSecurityIPGroupResponseBody) SetCode(v string) *DeleteSecurityIPGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSecurityIPGroupResponseBody) SetData(v *DeleteSecurityIPGroupResponseBodyData) *DeleteSecurityIPGroupResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSecurityIPGroupResponseBody) SetMessage(v string) *DeleteSecurityIPGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSecurityIPGroupResponseBody) SetRequestId(v string) *DeleteSecurityIPGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecurityIPGroupResponseBody) SetSuccess(v string) *DeleteSecurityIPGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSecurityIPGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSecurityIPGroupResponseBodyData struct {
	GlobalSecurityIPGroup []*DeleteSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup `json:"GlobalSecurityIPGroup,omitempty" xml:"GlobalSecurityIPGroup,omitempty" type:"Repeated"`
}

func (s DeleteSecurityIPGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityIPGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteSecurityIPGroupResponseBodyData) GetGlobalSecurityIPGroup() []*DeleteSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup {
	return s.GlobalSecurityIPGroup
}

func (s *DeleteSecurityIPGroupResponseBodyData) SetGlobalSecurityIPGroup(v []*DeleteSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) *DeleteSecurityIPGroupResponseBodyData {
	s.GlobalSecurityIPGroup = v
	return s
}

func (s *DeleteSecurityIPGroupResponseBodyData) Validate() error {
	if s.GlobalSecurityIPGroup != nil {
		for _, item := range s.GlobalSecurityIPGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup struct {
	// example:
	//
	// g-1no2rzybnqcv0m****
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
}

func (s DeleteSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) GoString() string {
	return s.String()
}

func (s *DeleteSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *DeleteSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) SetGlobalSecurityGroupId(v string) *DeleteSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *DeleteSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) Validate() error {
	return dara.Validate(s)
}
