// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpandGroupCapacityHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *ExpandGroupCapacityHeaders
  GetCommonHeaders() map[string]*string 
  SetAccountContext(v *ExpandGroupCapacityHeadersAccountContext) *ExpandGroupCapacityHeaders
  GetAccountContext() *ExpandGroupCapacityHeadersAccountContext 
}

type ExpandGroupCapacityHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  AccountContext *ExpandGroupCapacityHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ExpandGroupCapacityHeaders) String() string {
  return dara.Prettify(s)
}

func (s ExpandGroupCapacityHeaders) GoString() string {
  return s.String()
}

func (s *ExpandGroupCapacityHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *ExpandGroupCapacityHeaders) GetAccountContext() *ExpandGroupCapacityHeadersAccountContext  {
  return s.AccountContext
}

func (s *ExpandGroupCapacityHeaders) SetCommonHeaders(v map[string]*string) *ExpandGroupCapacityHeaders {
  s.CommonHeaders = v
  return s
}

func (s *ExpandGroupCapacityHeaders) SetAccountContext(v *ExpandGroupCapacityHeadersAccountContext) *ExpandGroupCapacityHeaders {
  s.AccountContext = v
  return s
}

func (s *ExpandGroupCapacityHeaders) Validate() error {
  if s.AccountContext != nil {
    if err := s.AccountContext.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExpandGroupCapacityHeadersAccountContext struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 012345
  AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ExpandGroupCapacityHeadersAccountContext) String() string {
  return dara.Prettify(s)
}

func (s ExpandGroupCapacityHeadersAccountContext) GoString() string {
  return s.String()
}

func (s *ExpandGroupCapacityHeadersAccountContext) GetAccountId() *string  {
  return s.AccountId
}

func (s *ExpandGroupCapacityHeadersAccountContext) SetAccountId(v string) *ExpandGroupCapacityHeadersAccountContext {
  s.AccountId = &v
  return s
}

func (s *ExpandGroupCapacityHeadersAccountContext) Validate() error {
  return dara.Validate(s)
}

