// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportSuspEventsResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetFileName(v string) *ExportSuspEventsResponseBody
  GetFileName() *string 
  SetId(v int32) *ExportSuspEventsResponseBody
  GetId() *int32 
  SetRequestId(v string) *ExportSuspEventsResponseBody
  GetRequestId() *string 
}

type ExportSuspEventsResponseBody struct {
  // The name of the exported file.
  // 
  // example:
  // 
  // suspicious_event_20221209
  FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
  // The ID of the export record of the anomalous event.
  // 
  // example:
  // 
  // 1
  Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // EF145C20-6A19-529A-8BDD-0671DXXXXXX
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportSuspEventsResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportSuspEventsResponseBody) GoString() string {
  return s.String()
}

func (s *ExportSuspEventsResponseBody) GetFileName() *string  {
  return s.FileName
}

func (s *ExportSuspEventsResponseBody) GetId() *int32  {
  return s.Id
}

func (s *ExportSuspEventsResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportSuspEventsResponseBody) SetFileName(v string) *ExportSuspEventsResponseBody {
  s.FileName = &v
  return s
}

func (s *ExportSuspEventsResponseBody) SetId(v int32) *ExportSuspEventsResponseBody {
  s.Id = &v
  return s
}

func (s *ExportSuspEventsResponseBody) SetRequestId(v string) *ExportSuspEventsResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportSuspEventsResponseBody) Validate() error {
  return dara.Validate(s)
}

