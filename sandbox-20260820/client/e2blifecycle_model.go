// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iE2BLifecycle interface {
  dara.Model
  String() string
  GoString() string
  SetAutoResume(v bool) *E2BLifecycle
  GetAutoResume() *bool 
  SetOnTimeout(v string) *E2BLifecycle
  GetOnTimeout() *string 
}

type E2BLifecycle struct {
  AutoResume *bool `json:"autoResume,omitempty" xml:"autoResume,omitempty"`
  OnTimeout *string `json:"onTimeout,omitempty" xml:"onTimeout,omitempty"`
}

func (s E2BLifecycle) String() string {
  return dara.Prettify(s)
}

func (s E2BLifecycle) GoString() string {
  return s.String()
}

func (s *E2BLifecycle) GetAutoResume() *bool  {
  return s.AutoResume
}

func (s *E2BLifecycle) GetOnTimeout() *string  {
  return s.OnTimeout
}

func (s *E2BLifecycle) SetAutoResume(v bool) *E2BLifecycle {
  s.AutoResume = &v
  return s
}

func (s *E2BLifecycle) SetOnTimeout(v string) *E2BLifecycle {
  s.OnTimeout = &v
  return s
}

func (s *E2BLifecycle) Validate() error {
  return dara.Validate(s)
}

