// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndDialogueResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EndDialogueResponseBody
  GetRequestId() *string 
}

type EndDialogueResponseBody struct {
  // example:
  // 
  // e48e45dd-e47a-4744-a063-f08cbebb1c5a
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EndDialogueResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EndDialogueResponseBody) GoString() string {
  return s.String()
}

func (s *EndDialogueResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EndDialogueResponseBody) SetRequestId(v string) *EndDialogueResponseBody {
  s.RequestId = &v
  return s
}

func (s *EndDialogueResponseBody) Validate() error {
  return dara.Validate(s)
}

