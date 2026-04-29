// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportRecallManagementTableResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRecallManagementJobId(v string) *ExportRecallManagementTableResponseBody
  GetRecallManagementJobId() *string 
  SetRequestId(v string) *ExportRecallManagementTableResponseBody
  GetRequestId() *string 
}

type ExportRecallManagementTableResponseBody struct {
  // example:
  // 
  // 1
  RecallManagementJobId *string `json:"RecallManagementJobId,omitempty" xml:"RecallManagementJobId,omitempty"`
  // example:
  // 
  // 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportRecallManagementTableResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportRecallManagementTableResponseBody) GoString() string {
  return s.String()
}

func (s *ExportRecallManagementTableResponseBody) GetRecallManagementJobId() *string  {
  return s.RecallManagementJobId
}

func (s *ExportRecallManagementTableResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportRecallManagementTableResponseBody) SetRecallManagementJobId(v string) *ExportRecallManagementTableResponseBody {
  s.RecallManagementJobId = &v
  return s
}

func (s *ExportRecallManagementTableResponseBody) SetRequestId(v string) *ExportRecallManagementTableResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportRecallManagementTableResponseBody) Validate() error {
  return dara.Validate(s)
}

