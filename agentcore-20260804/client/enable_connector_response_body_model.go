// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableConnectorResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableConnectorResponseBody
  GetCode() *string 
  SetData(v *EnableConnectorResponseBodyData) *EnableConnectorResponseBody
  GetData() *EnableConnectorResponseBodyData 
  SetHttpStatusCode(v int32) *EnableConnectorResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *EnableConnectorResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableConnectorResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableConnectorResponseBody
  GetSuccess() *bool 
}

type EnableConnectorResponseBody struct {
  // The business status code.
  // 
  // example:
  // 
  // SUCCESS
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  // The Connector details.
  Data *EnableConnectorResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s EnableConnectorResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableConnectorResponseBody) GoString() string {
  return s.String()
}

func (s *EnableConnectorResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableConnectorResponseBody) GetData() *EnableConnectorResponseBodyData  {
  return s.Data
}

func (s *EnableConnectorResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EnableConnectorResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableConnectorResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableConnectorResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableConnectorResponseBody) SetCode(v string) *EnableConnectorResponseBody {
  s.Code = &v
  return s
}

func (s *EnableConnectorResponseBody) SetData(v *EnableConnectorResponseBodyData) *EnableConnectorResponseBody {
  s.Data = v
  return s
}

func (s *EnableConnectorResponseBody) SetHttpStatusCode(v int32) *EnableConnectorResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EnableConnectorResponseBody) SetMessage(v string) *EnableConnectorResponseBody {
  s.Message = &v
  return s
}

func (s *EnableConnectorResponseBody) SetRequestId(v string) *EnableConnectorResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableConnectorResponseBody) SetSuccess(v bool) *EnableConnectorResponseBody {
  s.Success = &v
  return s
}

func (s *EnableConnectorResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnableConnectorResponseBodyData struct {
  // The number of bound agents.
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
  // The Connector configuration JSON string. This is sensitive information.
  // 
  // example:
  // 
  // {"site":"global","organizationId":"org-xxxx"}
  Metadata *string `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // The Connector name. The current value is qodercli.
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

func (s EnableConnectorResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnableConnectorResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnableConnectorResponseBodyData) GetBoundAgentCount() *int64  {
  return s.BoundAgentCount
}

func (s *EnableConnectorResponseBodyData) GetEnabledAt() *string  {
  return s.EnabledAt
}

func (s *EnableConnectorResponseBodyData) GetMetadata() *string  {
  return s.Metadata
}

func (s *EnableConnectorResponseBodyData) GetName() *string  {
  return s.Name
}

func (s *EnableConnectorResponseBodyData) GetStatus() *string  {
  return s.Status
}

func (s *EnableConnectorResponseBodyData) SetBoundAgentCount(v int64) *EnableConnectorResponseBodyData {
  s.BoundAgentCount = &v
  return s
}

func (s *EnableConnectorResponseBodyData) SetEnabledAt(v string) *EnableConnectorResponseBodyData {
  s.EnabledAt = &v
  return s
}

func (s *EnableConnectorResponseBodyData) SetMetadata(v string) *EnableConnectorResponseBodyData {
  s.Metadata = &v
  return s
}

func (s *EnableConnectorResponseBodyData) SetName(v string) *EnableConnectorResponseBodyData {
  s.Name = &v
  return s
}

func (s *EnableConnectorResponseBodyData) SetStatus(v string) *EnableConnectorResponseBodyData {
  s.Status = &v
  return s
}

func (s *EnableConnectorResponseBodyData) Validate() error {
  return dara.Validate(s)
}

