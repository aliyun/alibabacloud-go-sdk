// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportConversationDetailsResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetExportTaskId(v string) *ExportConversationDetailsResponseBody
  GetExportTaskId() *string 
  SetRequestId(v string) *ExportConversationDetailsResponseBody
  GetRequestId() *string 
}

type ExportConversationDetailsResponseBody struct {
  // example:
  // 
  // 6203fc87271a420c98eab6c2bbc2d856
  ExportTaskId *string `json:"ExportTaskId,omitempty" xml:"ExportTaskId,omitempty"`
  // example:
  // 
  // 75BAAB9B-40B2-5FF5-A59A-7BCF8154C6EE
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportConversationDetailsResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportConversationDetailsResponseBody) GoString() string {
  return s.String()
}

func (s *ExportConversationDetailsResponseBody) GetExportTaskId() *string  {
  return s.ExportTaskId
}

func (s *ExportConversationDetailsResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportConversationDetailsResponseBody) SetExportTaskId(v string) *ExportConversationDetailsResponseBody {
  s.ExportTaskId = &v
  return s
}

func (s *ExportConversationDetailsResponseBody) SetRequestId(v string) *ExportConversationDetailsResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportConversationDetailsResponseBody) Validate() error {
  return dara.Validate(s)
}

