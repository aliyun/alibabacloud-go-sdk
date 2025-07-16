// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpandGroupCapacityShrinkHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *ExpandGroupCapacityShrinkHeaders
  GetCommonHeaders() map[string]*string 
  SetAccountContextShrink(v string) *ExpandGroupCapacityShrinkHeaders
  GetAccountContextShrink() *string 
}

type ExpandGroupCapacityShrinkHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  AccountContextShrink *string `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ExpandGroupCapacityShrinkHeaders) String() string {
  return dara.Prettify(s)
}

func (s ExpandGroupCapacityShrinkHeaders) GoString() string {
  return s.String()
}

func (s *ExpandGroupCapacityShrinkHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *ExpandGroupCapacityShrinkHeaders) GetAccountContextShrink() *string  {
  return s.AccountContextShrink
}

func (s *ExpandGroupCapacityShrinkHeaders) SetCommonHeaders(v map[string]*string) *ExpandGroupCapacityShrinkHeaders {
  s.CommonHeaders = v
  return s
}

func (s *ExpandGroupCapacityShrinkHeaders) SetAccountContextShrink(v string) *ExpandGroupCapacityShrinkHeaders {
  s.AccountContextShrink = &v
  return s
}

func (s *ExpandGroupCapacityShrinkHeaders) Validate() error {
  return dara.Validate(s)
}

