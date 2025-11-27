// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExtractDocumentTextResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetDocumentText(v string) *ExtractDocumentTextResponseBody
  GetDocumentText() *string 
  SetRequestId(v string) *ExtractDocumentTextResponseBody
  GetRequestId() *string 
}

type ExtractDocumentTextResponseBody struct {
  // The text content of the document.
  // 
  // example:
  // 
  // 测试内容。
  DocumentText *string `json:"DocumentText,omitempty" xml:"DocumentText,omitempty"`
  // Request ID.
  // 
  // example:
  // 
  // 94D6F994-E298-037E-8E8B-0090F27*****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExtractDocumentTextResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExtractDocumentTextResponseBody) GoString() string {
  return s.String()
}

func (s *ExtractDocumentTextResponseBody) GetDocumentText() *string  {
  return s.DocumentText
}

func (s *ExtractDocumentTextResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExtractDocumentTextResponseBody) SetDocumentText(v string) *ExtractDocumentTextResponseBody {
  s.DocumentText = &v
  return s
}

func (s *ExtractDocumentTextResponseBody) SetRequestId(v string) *ExtractDocumentTextResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExtractDocumentTextResponseBody) Validate() error {
  return dara.Validate(s)
}

