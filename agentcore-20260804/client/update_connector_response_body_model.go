// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateConnectorResponseBody
	GetCode() *string
	SetData(v *UpdateConnectorResponseBodyData) *UpdateConnectorResponseBody
	GetData() *UpdateConnectorResponseBodyData
	SetHttpStatusCode(v int32) *UpdateConnectorResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateConnectorResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateConnectorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateConnectorResponseBody
	GetSuccess() *bool
}

type UpdateConnectorResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The Connector details.
	Data *UpdateConnectorResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConnectorResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateConnectorResponseBody) GetData() *UpdateConnectorResponseBodyData {
	return s.Data
}

func (s *UpdateConnectorResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateConnectorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConnectorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateConnectorResponseBody) SetCode(v string) *UpdateConnectorResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConnectorResponseBody) SetData(v *UpdateConnectorResponseBodyData) *UpdateConnectorResponseBody {
	s.Data = v
	return s
}

func (s *UpdateConnectorResponseBody) SetHttpStatusCode(v int32) *UpdateConnectorResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateConnectorResponseBody) SetMessage(v string) *UpdateConnectorResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConnectorResponseBody) SetRequestId(v string) *UpdateConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConnectorResponseBody) SetSuccess(v bool) *UpdateConnectorResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateConnectorResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateConnectorResponseBodyData struct {
	// The number of Agents attached to the Connector.
	//
	// example:
	//
	// 3
	BoundAgentCount *int64 `json:"boundAgentCount,omitempty" xml:"boundAgentCount,omitempty"`
	// The time when the Connector was enabled.
	//
	// example:
	//
	// 2026-09-01T08:00:00Z
	EnabledAt *string `json:"enabledAt,omitempty" xml:"enabledAt,omitempty"`
	// A JSON string. For qodercli: {"site":"global|cn","organizationId":"...","apiKey":"...","serviceAccountKeys":[{"id":"ckey-xxx","name":"default","serviceAccountKey":"..."}]}. This field is empty when the Connector is not enabled.
	//
	// example:
	//
	// {"site":"global","organizationId":"org-xxxx"}
	Metadata *string `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The Connector name.
	//
	// example:
	//
	// qodercli
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The Connector status.
	//
	// example:
	//
	// ENABLED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateConnectorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectorResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateConnectorResponseBodyData) GetBoundAgentCount() *int64 {
	return s.BoundAgentCount
}

func (s *UpdateConnectorResponseBodyData) GetEnabledAt() *string {
	return s.EnabledAt
}

func (s *UpdateConnectorResponseBodyData) GetMetadata() *string {
	return s.Metadata
}

func (s *UpdateConnectorResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateConnectorResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateConnectorResponseBodyData) SetBoundAgentCount(v int64) *UpdateConnectorResponseBodyData {
	s.BoundAgentCount = &v
	return s
}

func (s *UpdateConnectorResponseBodyData) SetEnabledAt(v string) *UpdateConnectorResponseBodyData {
	s.EnabledAt = &v
	return s
}

func (s *UpdateConnectorResponseBodyData) SetMetadata(v string) *UpdateConnectorResponseBodyData {
	s.Metadata = &v
	return s
}

func (s *UpdateConnectorResponseBodyData) SetName(v string) *UpdateConnectorResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateConnectorResponseBodyData) SetStatus(v string) *UpdateConnectorResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateConnectorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
