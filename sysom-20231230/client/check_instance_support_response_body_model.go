// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceSupportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckInstanceSupportResponseBody
	GetCode() *string
	SetData(v []*CheckInstanceSupportResponseBodyData) *CheckInstanceSupportResponseBody
	GetData() []*CheckInstanceSupportResponseBodyData
	SetMessage(v string) *CheckInstanceSupportResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckInstanceSupportResponseBody
	GetRequestId() *string
}

type CheckInstanceSupportResponseBody struct {
	// example:
	//
	// Success
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data []*CheckInstanceSupportResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CheckInstanceSupportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceSupportResponseBody) GoString() string {
	return s.String()
}

func (s *CheckInstanceSupportResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckInstanceSupportResponseBody) GetData() []*CheckInstanceSupportResponseBodyData {
	return s.Data
}

func (s *CheckInstanceSupportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckInstanceSupportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckInstanceSupportResponseBody) SetCode(v string) *CheckInstanceSupportResponseBody {
	s.Code = &v
	return s
}

func (s *CheckInstanceSupportResponseBody) SetData(v []*CheckInstanceSupportResponseBodyData) *CheckInstanceSupportResponseBody {
	s.Data = v
	return s
}

func (s *CheckInstanceSupportResponseBody) SetMessage(v string) *CheckInstanceSupportResponseBody {
	s.Message = &v
	return s
}

func (s *CheckInstanceSupportResponseBody) SetRequestId(v string) *CheckInstanceSupportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckInstanceSupportResponseBody) Validate() error {
	return dara.Validate(s)
}

type CheckInstanceSupportResponseBodyData struct {
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// instance not found in ecs
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// example:
	//
	// true
	Support *bool `json:"support,omitempty" xml:"support,omitempty"`
}

func (s CheckInstanceSupportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceSupportResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckInstanceSupportResponseBodyData) GetInstance() *string {
	return s.Instance
}

func (s *CheckInstanceSupportResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *CheckInstanceSupportResponseBodyData) GetSupport() *bool {
	return s.Support
}

func (s *CheckInstanceSupportResponseBodyData) SetInstance(v string) *CheckInstanceSupportResponseBodyData {
	s.Instance = &v
	return s
}

func (s *CheckInstanceSupportResponseBodyData) SetReason(v string) *CheckInstanceSupportResponseBodyData {
	s.Reason = &v
	return s
}

func (s *CheckInstanceSupportResponseBodyData) SetSupport(v bool) *CheckInstanceSupportResponseBodyData {
	s.Support = &v
	return s
}

func (s *CheckInstanceSupportResponseBodyData) Validate() error {
	return dara.Validate(s)
}
