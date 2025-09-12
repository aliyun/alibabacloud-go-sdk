// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCollectionGraphRAGResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetJobId(v string) *EnableCollectionGraphRAGResponseBody
  GetJobId() *string 
  SetMessage(v string) *EnableCollectionGraphRAGResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableCollectionGraphRAGResponseBody
  GetRequestId() *string 
  SetStatus(v string) *EnableCollectionGraphRAGResponseBody
  GetStatus() *string 
}

type EnableCollectionGraphRAGResponseBody struct {
  // example:
  // 
  // 231460f8-75dc-405e-a669-0c5204887e91
  JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // ABB39CC3-4488-4857-905D-2E4A051D0521
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // success
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EnableCollectionGraphRAGResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableCollectionGraphRAGResponseBody) GoString() string {
  return s.String()
}

func (s *EnableCollectionGraphRAGResponseBody) GetJobId() *string  {
  return s.JobId
}

func (s *EnableCollectionGraphRAGResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableCollectionGraphRAGResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableCollectionGraphRAGResponseBody) GetStatus() *string  {
  return s.Status
}

func (s *EnableCollectionGraphRAGResponseBody) SetJobId(v string) *EnableCollectionGraphRAGResponseBody {
  s.JobId = &v
  return s
}

func (s *EnableCollectionGraphRAGResponseBody) SetMessage(v string) *EnableCollectionGraphRAGResponseBody {
  s.Message = &v
  return s
}

func (s *EnableCollectionGraphRAGResponseBody) SetRequestId(v string) *EnableCollectionGraphRAGResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableCollectionGraphRAGResponseBody) SetStatus(v string) *EnableCollectionGraphRAGResponseBody {
  s.Status = &v
  return s
}

func (s *EnableCollectionGraphRAGResponseBody) Validate() error {
  return dara.Validate(s)
}

