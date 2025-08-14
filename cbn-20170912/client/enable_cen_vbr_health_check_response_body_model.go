// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCenVbrHealthCheckResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableCenVbrHealthCheckResponseBody
  GetRequestId() *string 
}

type EnableCenVbrHealthCheckResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 1F59F19C-EFD2-40B1-94D5-65B40CA8E34A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableCenVbrHealthCheckResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableCenVbrHealthCheckResponseBody) GoString() string {
  return s.String()
}

func (s *EnableCenVbrHealthCheckResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableCenVbrHealthCheckResponseBody) SetRequestId(v string) *EnableCenVbrHealthCheckResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableCenVbrHealthCheckResponseBody) Validate() error {
  return dara.Validate(s)
}

