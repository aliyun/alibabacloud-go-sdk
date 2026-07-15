// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEstimatedPriceQueryV2ResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *EstimatedPriceQueryV2ResponseBody
  GetCode() *int32 
  SetModule(v []*EstimatedPriceQueryV2ResponseBodyModule) *EstimatedPriceQueryV2ResponseBody
  GetModule() []*EstimatedPriceQueryV2ResponseBodyModule 
  SetRequestId(v string) *EstimatedPriceQueryV2ResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EstimatedPriceQueryV2ResponseBody
  GetSuccess() *bool 
  SetTraceId(v string) *EstimatedPriceQueryV2ResponseBody
  GetTraceId() *string 
}

type EstimatedPriceQueryV2ResponseBody struct {
  Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
  Module []*EstimatedPriceQueryV2ResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
  TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s EstimatedPriceQueryV2ResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EstimatedPriceQueryV2ResponseBody) GoString() string {
  return s.String()
}

func (s *EstimatedPriceQueryV2ResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *EstimatedPriceQueryV2ResponseBody) GetModule() []*EstimatedPriceQueryV2ResponseBodyModule  {
  return s.Module
}

func (s *EstimatedPriceQueryV2ResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EstimatedPriceQueryV2ResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EstimatedPriceQueryV2ResponseBody) GetTraceId() *string  {
  return s.TraceId
}

func (s *EstimatedPriceQueryV2ResponseBody) SetCode(v int32) *EstimatedPriceQueryV2ResponseBody {
  s.Code = &v
  return s
}

func (s *EstimatedPriceQueryV2ResponseBody) SetModule(v []*EstimatedPriceQueryV2ResponseBodyModule) *EstimatedPriceQueryV2ResponseBody {
  s.Module = v
  return s
}

func (s *EstimatedPriceQueryV2ResponseBody) SetRequestId(v string) *EstimatedPriceQueryV2ResponseBody {
  s.RequestId = &v
  return s
}

func (s *EstimatedPriceQueryV2ResponseBody) SetSuccess(v bool) *EstimatedPriceQueryV2ResponseBody {
  s.Success = &v
  return s
}

func (s *EstimatedPriceQueryV2ResponseBody) SetTraceId(v string) *EstimatedPriceQueryV2ResponseBody {
  s.TraceId = &v
  return s
}

func (s *EstimatedPriceQueryV2ResponseBody) Validate() error {
  if s.Module != nil {
    for _, item := range s.Module {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EstimatedPriceQueryV2ResponseBodyModule struct {
  BizType *string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
  PriceMap map[string]*ModulePriceMapValue `json:"price_map,omitempty" xml:"price_map,omitempty"`
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s EstimatedPriceQueryV2ResponseBodyModule) String() string {
  return dara.Prettify(s)
}

func (s EstimatedPriceQueryV2ResponseBodyModule) GoString() string {
  return s.String()
}

func (s *EstimatedPriceQueryV2ResponseBodyModule) GetBizType() *string  {
  return s.BizType
}

func (s *EstimatedPriceQueryV2ResponseBodyModule) GetPriceMap() map[string]*ModulePriceMapValue  {
  return s.PriceMap
}

func (s *EstimatedPriceQueryV2ResponseBodyModule) GetType() *string  {
  return s.Type
}

func (s *EstimatedPriceQueryV2ResponseBodyModule) SetBizType(v string) *EstimatedPriceQueryV2ResponseBodyModule {
  s.BizType = &v
  return s
}

func (s *EstimatedPriceQueryV2ResponseBodyModule) SetPriceMap(v map[string]*ModulePriceMapValue) *EstimatedPriceQueryV2ResponseBodyModule {
  s.PriceMap = v
  return s
}

func (s *EstimatedPriceQueryV2ResponseBodyModule) SetType(v string) *EstimatedPriceQueryV2ResponseBodyModule {
  s.Type = &v
  return s
}

func (s *EstimatedPriceQueryV2ResponseBodyModule) Validate() error {
  return dara.Validate(s)
}

