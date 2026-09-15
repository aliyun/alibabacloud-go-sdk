// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableConnectorResponseBody
	GetCode() *string
	SetData(v *DisableConnectorResponseBodyData) *DisableConnectorResponseBody
	GetData() *DisableConnectorResponseBodyData
	SetHttpStatusCode(v int32) *DisableConnectorResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DisableConnectorResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableConnectorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableConnectorResponseBody
	GetSuccess() *bool
}

type DisableConnectorResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The Connector details.
	Data *DisableConnectorResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s DisableConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *DisableConnectorResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableConnectorResponseBody) GetData() *DisableConnectorResponseBodyData {
	return s.Data
}

func (s *DisableConnectorResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DisableConnectorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableConnectorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableConnectorResponseBody) SetCode(v string) *DisableConnectorResponseBody {
	s.Code = &v
	return s
}

func (s *DisableConnectorResponseBody) SetData(v *DisableConnectorResponseBodyData) *DisableConnectorResponseBody {
	s.Data = v
	return s
}

func (s *DisableConnectorResponseBody) SetHttpStatusCode(v int32) *DisableConnectorResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisableConnectorResponseBody) SetMessage(v string) *DisableConnectorResponseBody {
	s.Message = &v
	return s
}

func (s *DisableConnectorResponseBody) SetRequestId(v string) *DisableConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableConnectorResponseBody) SetSuccess(v bool) *DisableConnectorResponseBody {
	s.Success = &v
	return s
}

func (s *DisableConnectorResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DisableConnectorResponseBodyData struct {
	// The number of Agents bound to the Connector.
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

func (s DisableConnectorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DisableConnectorResponseBodyData) GoString() string {
	return s.String()
}

func (s *DisableConnectorResponseBodyData) GetBoundAgentCount() *int64 {
	return s.BoundAgentCount
}

func (s *DisableConnectorResponseBodyData) GetEnabledAt() *string {
	return s.EnabledAt
}

func (s *DisableConnectorResponseBodyData) GetMetadata() *string {
	return s.Metadata
}

func (s *DisableConnectorResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DisableConnectorResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DisableConnectorResponseBodyData) SetBoundAgentCount(v int64) *DisableConnectorResponseBodyData {
	s.BoundAgentCount = &v
	return s
}

func (s *DisableConnectorResponseBodyData) SetEnabledAt(v string) *DisableConnectorResponseBodyData {
	s.EnabledAt = &v
	return s
}

func (s *DisableConnectorResponseBodyData) SetMetadata(v string) *DisableConnectorResponseBodyData {
	s.Metadata = &v
	return s
}

func (s *DisableConnectorResponseBodyData) SetName(v string) *DisableConnectorResponseBodyData {
	s.Name = &v
	return s
}

func (s *DisableConnectorResponseBodyData) SetStatus(v string) *DisableConnectorResponseBodyData {
	s.Status = &v
	return s
}

func (s *DisableConnectorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
