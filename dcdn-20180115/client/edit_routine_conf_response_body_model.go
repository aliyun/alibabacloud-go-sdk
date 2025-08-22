// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditRoutineConfResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetContent(v map[string]interface{}) *EditRoutineConfResponseBody
  GetContent() map[string]interface{} 
  SetRequestId(v string) *EditRoutineConfResponseBody
  GetRequestId() *string 
}

type EditRoutineConfResponseBody struct {
  // The description of the execution errors and the version number of the latest environment configurations.
  Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // BAECB354-6D42-42C1-87DA-C9992EF1E7C8
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditRoutineConfResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditRoutineConfResponseBody) GoString() string {
  return s.String()
}

func (s *EditRoutineConfResponseBody) GetContent() map[string]interface{}  {
  return s.Content
}

func (s *EditRoutineConfResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditRoutineConfResponseBody) SetContent(v map[string]interface{}) *EditRoutineConfResponseBody {
  s.Content = v
  return s
}

func (s *EditRoutineConfResponseBody) SetRequestId(v string) *EditRoutineConfResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditRoutineConfResponseBody) Validate() error {
  return dara.Validate(s)
}

