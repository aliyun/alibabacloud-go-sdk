// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExtractFingerPrintResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExtractFingerPrintResponseBodyData) *ExtractFingerPrintResponseBody
  GetData() *ExtractFingerPrintResponseBodyData 
  SetRequestId(v string) *ExtractFingerPrintResponseBody
  GetRequestId() *string 
}

type ExtractFingerPrintResponseBody struct {
  Data *ExtractFingerPrintResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // example:
  // 
  // D21978CC-C1E7-4A7A-A1B2-D5896BDC7ADF
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExtractFingerPrintResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExtractFingerPrintResponseBody) GoString() string {
  return s.String()
}

func (s *ExtractFingerPrintResponseBody) GetData() *ExtractFingerPrintResponseBodyData  {
  return s.Data
}

func (s *ExtractFingerPrintResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExtractFingerPrintResponseBody) SetData(v *ExtractFingerPrintResponseBodyData) *ExtractFingerPrintResponseBody {
  s.Data = v
  return s
}

func (s *ExtractFingerPrintResponseBody) SetRequestId(v string) *ExtractFingerPrintResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExtractFingerPrintResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExtractFingerPrintResponseBodyData struct {
  // example:
  // 
  // /9j/4AAQSkZJRgABAQAAAQABAAD****
  FingerPrint *string `json:"FingerPrint,omitempty" xml:"FingerPrint,omitempty"`
}

func (s ExtractFingerPrintResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExtractFingerPrintResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExtractFingerPrintResponseBodyData) GetFingerPrint() *string  {
  return s.FingerPrint
}

func (s *ExtractFingerPrintResponseBodyData) SetFingerPrint(v string) *ExtractFingerPrintResponseBodyData {
  s.FingerPrint = &v
  return s
}

func (s *ExtractFingerPrintResponseBodyData) Validate() error {
  return dara.Validate(s)
}

