// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAgentConfigSupportedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckAgentConfigSupportedResponseBody
	GetCode() *string
	SetData(v []*CheckAgentConfigSupportedResponseBodyData) *CheckAgentConfigSupportedResponseBody
	GetData() []*CheckAgentConfigSupportedResponseBodyData
	SetMessage(v string) *CheckAgentConfigSupportedResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckAgentConfigSupportedResponseBody
	GetRequestId() *string
	SetTotal(v int64) *CheckAgentConfigSupportedResponseBody
	GetTotal() *int64
}

type CheckAgentConfigSupportedResponseBody struct {
	// The status code.
	//
	// - If `code` is `Success`, the authorization is successful.
	//
	// - Other status codes indicate that the authorization failed. If the authorization failed, check the `message` field for the detailed fault information.
	//
	// example:
	//
	// Success  Fail
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned result.
	Data []*CheckAgentConfigSupportedResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error message.
	//
	// - If code is Success, this field is empty.
	//
	// - Otherwise, this field contains the error information of the request.
	//
	// example:
	//
	// server error
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of returned results.
	//
	// example:
	//
	// 42
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s CheckAgentConfigSupportedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckAgentConfigSupportedResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAgentConfigSupportedResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckAgentConfigSupportedResponseBody) GetData() []*CheckAgentConfigSupportedResponseBodyData {
	return s.Data
}

func (s *CheckAgentConfigSupportedResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckAgentConfigSupportedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckAgentConfigSupportedResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *CheckAgentConfigSupportedResponseBody) SetCode(v string) *CheckAgentConfigSupportedResponseBody {
	s.Code = &v
	return s
}

func (s *CheckAgentConfigSupportedResponseBody) SetData(v []*CheckAgentConfigSupportedResponseBodyData) *CheckAgentConfigSupportedResponseBody {
	s.Data = v
	return s
}

func (s *CheckAgentConfigSupportedResponseBody) SetMessage(v string) *CheckAgentConfigSupportedResponseBody {
	s.Message = &v
	return s
}

func (s *CheckAgentConfigSupportedResponseBody) SetRequestId(v string) *CheckAgentConfigSupportedResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAgentConfigSupportedResponseBody) SetTotal(v int64) *CheckAgentConfigSupportedResponseBody {
	s.Total = &v
	return s
}

func (s *CheckAgentConfigSupportedResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckAgentConfigSupportedResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// a36a284e-fc8a-4c8a-abad-03d441e6ebcf
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The ECS instance ID.
	//
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// The reason why the instance is not supported for management. This field is not empty when success is false.
	//
	// example:
	//
	// instance not found in ecs
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s CheckAgentConfigSupportedResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CheckAgentConfigSupportedResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckAgentConfigSupportedResponseBodyData) GetCluster() *string {
	return s.Cluster
}

func (s *CheckAgentConfigSupportedResponseBodyData) GetInstance() *string {
	return s.Instance
}

func (s *CheckAgentConfigSupportedResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *CheckAgentConfigSupportedResponseBodyData) SetCluster(v string) *CheckAgentConfigSupportedResponseBodyData {
	s.Cluster = &v
	return s
}

func (s *CheckAgentConfigSupportedResponseBodyData) SetInstance(v string) *CheckAgentConfigSupportedResponseBodyData {
	s.Instance = &v
	return s
}

func (s *CheckAgentConfigSupportedResponseBodyData) SetReason(v string) *CheckAgentConfigSupportedResponseBodyData {
	s.Reason = &v
	return s
}

func (s *CheckAgentConfigSupportedResponseBodyData) Validate() error {
	return dara.Validate(s)
}
