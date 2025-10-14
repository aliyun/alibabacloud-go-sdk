// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnrichShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAdults(v int32) *EnrichShrinkRequest
  GetAdults() *int32 
  SetCabinClass(v string) *EnrichShrinkRequest
  GetCabinClass() *string 
  SetChildren(v int32) *EnrichShrinkRequest
  GetChildren() *int32 
  SetInfants(v int32) *EnrichShrinkRequest
  GetInfants() *int32 
  SetJourneyParamListShrink(v string) *EnrichShrinkRequest
  GetJourneyParamListShrink() *string 
  SetSolutionId(v string) *EnrichShrinkRequest
  GetSolutionId() *string 
}

type EnrichShrinkRequest struct {
  // adult passenger amount 1-9
  // 
  // example:
  // 
  // 1
  Adults *int32 `json:"adults,omitempty" xml:"adults,omitempty"`
  // cabin class
  // 
  // 1. **ALL_CABIN*	- : all cabin class
  // 
  // 2. **Y*	- : economy class
  // 
  // 3. **FC*	- : first class and business class
  // 
  // 4. **S*	- : premium economy class
  // 
  // 5. **YS*	- : economy class and premium economy class
  // 
  // 6. **YSC*	- : economy class, premium economy class and business class
  // 
  // example:
  // 
  // ALL_CABIN
  CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
  // child passenger amount 0-9
  // 
  // example:
  // 
  // 1
  Children *int32 `json:"children,omitempty" xml:"children,omitempty"`
  // infant passenger amount 0-9
  // 
  // example:
  // 
  // 1
  Infants *int32 `json:"infants,omitempty" xml:"infants,omitempty"`
  // journey list
  JourneyParamListShrink *string `json:"journey_param_list,omitempty" xml:"journey_param_list,omitempty"`
  // solution_id returned by Search
  // 
  // example:
  // 
  // eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
  SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s EnrichShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EnrichShrinkRequest) GoString() string {
  return s.String()
}

func (s *EnrichShrinkRequest) GetAdults() *int32  {
  return s.Adults
}

func (s *EnrichShrinkRequest) GetCabinClass() *string  {
  return s.CabinClass
}

func (s *EnrichShrinkRequest) GetChildren() *int32  {
  return s.Children
}

func (s *EnrichShrinkRequest) GetInfants() *int32  {
  return s.Infants
}

func (s *EnrichShrinkRequest) GetJourneyParamListShrink() *string  {
  return s.JourneyParamListShrink
}

func (s *EnrichShrinkRequest) GetSolutionId() *string  {
  return s.SolutionId
}

func (s *EnrichShrinkRequest) SetAdults(v int32) *EnrichShrinkRequest {
  s.Adults = &v
  return s
}

func (s *EnrichShrinkRequest) SetCabinClass(v string) *EnrichShrinkRequest {
  s.CabinClass = &v
  return s
}

func (s *EnrichShrinkRequest) SetChildren(v int32) *EnrichShrinkRequest {
  s.Children = &v
  return s
}

func (s *EnrichShrinkRequest) SetInfants(v int32) *EnrichShrinkRequest {
  s.Infants = &v
  return s
}

func (s *EnrichShrinkRequest) SetJourneyParamListShrink(v string) *EnrichShrinkRequest {
  s.JourneyParamListShrink = &v
  return s
}

func (s *EnrichShrinkRequest) SetSolutionId(v string) *EnrichShrinkRequest {
  s.SolutionId = &v
  return s
}

func (s *EnrichShrinkRequest) Validate() error {
  return dara.Validate(s)
}

