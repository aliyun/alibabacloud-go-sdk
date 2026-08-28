// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyMigrationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VerifyMigrationTaskResponseBody
	GetCode() *string
	SetData(v *VerifyMigrationTaskResponseBodyData) *VerifyMigrationTaskResponseBody
	GetData() *VerifyMigrationTaskResponseBodyData
	SetMessage(v string) *VerifyMigrationTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *VerifyMigrationTaskResponseBody
	GetRequestId() *string
}

type VerifyMigrationTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Data *VerifyMigrationTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CB85272A-5F04-58D7-BDE1-8BB5EB390CE1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s VerifyMigrationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyMigrationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyMigrationTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifyMigrationTaskResponseBody) GetData() *VerifyMigrationTaskResponseBodyData {
	return s.Data
}

func (s *VerifyMigrationTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifyMigrationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyMigrationTaskResponseBody) SetCode(v string) *VerifyMigrationTaskResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyMigrationTaskResponseBody) SetData(v *VerifyMigrationTaskResponseBodyData) *VerifyMigrationTaskResponseBody {
	s.Data = v
	return s
}

func (s *VerifyMigrationTaskResponseBody) SetMessage(v string) *VerifyMigrationTaskResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyMigrationTaskResponseBody) SetRequestId(v string) *VerifyMigrationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyMigrationTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyMigrationTaskResponseBodyData struct {
	// example:
	//
	// true
	IsSupported *bool `json:"isSupported,omitempty" xml:"isSupported,omitempty"`
	// example:
	//
	// all routes supported
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Success               *bool                                                       `json:"success,omitempty" xml:"success,omitempty"`
	UnSupportedRouteRules []*VerifyMigrationTaskResponseBodyDataUnSupportedRouteRules `json:"unSupportedRouteRules,omitempty" xml:"unSupportedRouteRules,omitempty" type:"Repeated"`
}

func (s VerifyMigrationTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VerifyMigrationTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifyMigrationTaskResponseBodyData) GetIsSupported() *bool {
	return s.IsSupported
}

func (s *VerifyMigrationTaskResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *VerifyMigrationTaskResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *VerifyMigrationTaskResponseBodyData) GetUnSupportedRouteRules() []*VerifyMigrationTaskResponseBodyDataUnSupportedRouteRules {
	return s.UnSupportedRouteRules
}

func (s *VerifyMigrationTaskResponseBodyData) SetIsSupported(v bool) *VerifyMigrationTaskResponseBodyData {
	s.IsSupported = &v
	return s
}

func (s *VerifyMigrationTaskResponseBodyData) SetMessage(v string) *VerifyMigrationTaskResponseBodyData {
	s.Message = &v
	return s
}

func (s *VerifyMigrationTaskResponseBodyData) SetSuccess(v bool) *VerifyMigrationTaskResponseBodyData {
	s.Success = &v
	return s
}

func (s *VerifyMigrationTaskResponseBodyData) SetUnSupportedRouteRules(v []*VerifyMigrationTaskResponseBodyDataUnSupportedRouteRules) *VerifyMigrationTaskResponseBodyData {
	s.UnSupportedRouteRules = v
	return s
}

func (s *VerifyMigrationTaskResponseBodyData) Validate() error {
	if s.UnSupportedRouteRules != nil {
		for _, item := range s.UnSupportedRouteRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type VerifyMigrationTaskResponseBodyDataUnSupportedRouteRules struct {
	// example:
	//
	// default/my-ingress
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// /api/v1/pets -> pet-service
	Rule                   *string   `json:"rule,omitempty" xml:"rule,omitempty"`
	UnSupportedAnnotations []*string `json:"unSupportedAnnotations,omitempty" xml:"unSupportedAnnotations,omitempty" type:"Repeated"`
}

func (s VerifyMigrationTaskResponseBodyDataUnSupportedRouteRules) String() string {
	return dara.Prettify(s)
}

func (s VerifyMigrationTaskResponseBodyDataUnSupportedRouteRules) GoString() string {
	return s.String()
}

func (s *VerifyMigrationTaskResponseBodyDataUnSupportedRouteRules) GetName() *string {
	return s.Name
}

func (s *VerifyMigrationTaskResponseBodyDataUnSupportedRouteRules) GetRule() *string {
	return s.Rule
}

func (s *VerifyMigrationTaskResponseBodyDataUnSupportedRouteRules) GetUnSupportedAnnotations() []*string {
	return s.UnSupportedAnnotations
}

func (s *VerifyMigrationTaskResponseBodyDataUnSupportedRouteRules) SetName(v string) *VerifyMigrationTaskResponseBodyDataUnSupportedRouteRules {
	s.Name = &v
	return s
}

func (s *VerifyMigrationTaskResponseBodyDataUnSupportedRouteRules) SetRule(v string) *VerifyMigrationTaskResponseBodyDataUnSupportedRouteRules {
	s.Rule = &v
	return s
}

func (s *VerifyMigrationTaskResponseBodyDataUnSupportedRouteRules) SetUnSupportedAnnotations(v []*string) *VerifyMigrationTaskResponseBodyDataUnSupportedRouteRules {
	s.UnSupportedAnnotations = v
	return s
}

func (s *VerifyMigrationTaskResponseBodyDataUnSupportedRouteRules) Validate() error {
	return dara.Validate(s)
}
