// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableServicesResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetFailedServices(v []*EnableServicesResponseBodyFailedServices) *EnableServicesResponseBody
  GetFailedServices() []*EnableServicesResponseBodyFailedServices 
  SetRequestId(v string) *EnableServicesResponseBody
  GetRequestId() *string 
}

type EnableServicesResponseBody struct {
  FailedServices []*EnableServicesResponseBodyFailedServices `json:"FailedServices,omitempty" xml:"FailedServices,omitempty" type:"Repeated"`
  // example:
  // 
  // E1BD3327-6BEE-53AD-8788-D892EB575962
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableServicesResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableServicesResponseBody) GoString() string {
  return s.String()
}

func (s *EnableServicesResponseBody) GetFailedServices() []*EnableServicesResponseBodyFailedServices  {
  return s.FailedServices
}

func (s *EnableServicesResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableServicesResponseBody) SetFailedServices(v []*EnableServicesResponseBodyFailedServices) *EnableServicesResponseBody {
  s.FailedServices = v
  return s
}

func (s *EnableServicesResponseBody) SetRequestId(v string) *EnableServicesResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableServicesResponseBody) Validate() error {
  if s.FailedServices != nil {
    for _, item := range s.FailedServices {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EnableServicesResponseBodyFailedServices struct {
  // example:
  // 
  // 400
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // Failed
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // ACVS
  ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s EnableServicesResponseBodyFailedServices) String() string {
  return dara.Prettify(s)
}

func (s EnableServicesResponseBodyFailedServices) GoString() string {
  return s.String()
}

func (s *EnableServicesResponseBodyFailedServices) GetCode() *string  {
  return s.Code
}

func (s *EnableServicesResponseBodyFailedServices) GetMessage() *string  {
  return s.Message
}

func (s *EnableServicesResponseBodyFailedServices) GetServiceName() *string  {
  return s.ServiceName
}

func (s *EnableServicesResponseBodyFailedServices) SetCode(v string) *EnableServicesResponseBodyFailedServices {
  s.Code = &v
  return s
}

func (s *EnableServicesResponseBodyFailedServices) SetMessage(v string) *EnableServicesResponseBodyFailedServices {
  s.Message = &v
  return s
}

func (s *EnableServicesResponseBodyFailedServices) SetServiceName(v string) *EnableServicesResponseBodyFailedServices {
  s.ServiceName = &v
  return s
}

func (s *EnableServicesResponseBodyFailedServices) Validate() error {
  return dara.Validate(s)
}

