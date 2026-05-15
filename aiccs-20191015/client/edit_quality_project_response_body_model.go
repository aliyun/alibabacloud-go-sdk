// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditQualityProjectResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EditQualityProjectResponseBody
  GetCode() *string 
  SetData(v []*EditQualityProjectResponseBodyData) *EditQualityProjectResponseBody
  GetData() []*EditQualityProjectResponseBodyData 
  SetMessage(v string) *EditQualityProjectResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EditQualityProjectResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EditQualityProjectResponseBody
  GetSuccess() *bool 
}

type EditQualityProjectResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data []*EditQualityProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EditQualityProjectResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditQualityProjectResponseBody) GoString() string {
  return s.String()
}

func (s *EditQualityProjectResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EditQualityProjectResponseBody) GetData() []*EditQualityProjectResponseBodyData  {
  return s.Data
}

func (s *EditQualityProjectResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EditQualityProjectResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditQualityProjectResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EditQualityProjectResponseBody) SetCode(v string) *EditQualityProjectResponseBody {
  s.Code = &v
  return s
}

func (s *EditQualityProjectResponseBody) SetData(v []*EditQualityProjectResponseBodyData) *EditQualityProjectResponseBody {
  s.Data = v
  return s
}

func (s *EditQualityProjectResponseBody) SetMessage(v string) *EditQualityProjectResponseBody {
  s.Message = &v
  return s
}

func (s *EditQualityProjectResponseBody) SetRequestId(v string) *EditQualityProjectResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditQualityProjectResponseBody) SetSuccess(v bool) *EditQualityProjectResponseBody {
  s.Success = &v
  return s
}

func (s *EditQualityProjectResponseBody) Validate() error {
  if s.Data != nil {
    for _, item := range s.Data {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EditQualityProjectResponseBodyData struct {
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
  Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s EditQualityProjectResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EditQualityProjectResponseBodyData) GoString() string {
  return s.String()
}

func (s *EditQualityProjectResponseBodyData) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EditQualityProjectResponseBodyData) GetProjectId() *int64  {
  return s.ProjectId
}

func (s *EditQualityProjectResponseBodyData) GetVersion() *int32  {
  return s.Version
}

func (s *EditQualityProjectResponseBodyData) SetInstanceId(v string) *EditQualityProjectResponseBodyData {
  s.InstanceId = &v
  return s
}

func (s *EditQualityProjectResponseBodyData) SetProjectId(v int64) *EditQualityProjectResponseBodyData {
  s.ProjectId = &v
  return s
}

func (s *EditQualityProjectResponseBodyData) SetVersion(v int32) *EditQualityProjectResponseBodyData {
  s.Version = &v
  return s
}

func (s *EditQualityProjectResponseBodyData) Validate() error {
  return dara.Validate(s)
}

