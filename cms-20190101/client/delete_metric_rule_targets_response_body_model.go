// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetricRuleTargetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMetricRuleTargetsResponseBody
	GetCode() *string
	SetFailIds(v *DeleteMetricRuleTargetsResponseBodyFailIds) *DeleteMetricRuleTargetsResponseBody
	GetFailIds() *DeleteMetricRuleTargetsResponseBodyFailIds
	SetMessage(v string) *DeleteMetricRuleTargetsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMetricRuleTargetsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMetricRuleTargetsResponseBody
	GetSuccess() *bool
}

type DeleteMetricRuleTargetsResponseBody struct {
	// The HTTP status code.
	//
	// **
	//
	// **Description*	- The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The IDs of the resources that failed to be deleted.
	FailIds *DeleteMetricRuleTargetsResponseBodyFailIds `json:"FailIds,omitempty" xml:"FailIds,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 786E92D2-AC66-4250-B76F-F1E2FCDDBA1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMetricRuleTargetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetricRuleTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTargetsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMetricRuleTargetsResponseBody) GetFailIds() *DeleteMetricRuleTargetsResponseBodyFailIds {
	return s.FailIds
}

func (s *DeleteMetricRuleTargetsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMetricRuleTargetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMetricRuleTargetsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMetricRuleTargetsResponseBody) SetCode(v string) *DeleteMetricRuleTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMetricRuleTargetsResponseBody) SetFailIds(v *DeleteMetricRuleTargetsResponseBodyFailIds) *DeleteMetricRuleTargetsResponseBody {
	s.FailIds = v
	return s
}

func (s *DeleteMetricRuleTargetsResponseBody) SetMessage(v string) *DeleteMetricRuleTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMetricRuleTargetsResponseBody) SetRequestId(v string) *DeleteMetricRuleTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMetricRuleTargetsResponseBody) SetSuccess(v bool) *DeleteMetricRuleTargetsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMetricRuleTargetsResponseBody) Validate() error {
	if s.FailIds != nil {
		if err := s.FailIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMetricRuleTargetsResponseBodyFailIds struct {
	// The IDs of the resources that failed to be deleted.
	TargetIds *DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds `json:"TargetIds,omitempty" xml:"TargetIds,omitempty" type:"Struct"`
}

func (s DeleteMetricRuleTargetsResponseBodyFailIds) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetricRuleTargetsResponseBodyFailIds) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTargetsResponseBodyFailIds) GetTargetIds() *DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds {
	return s.TargetIds
}

func (s *DeleteMetricRuleTargetsResponseBodyFailIds) SetTargetIds(v *DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds) *DeleteMetricRuleTargetsResponseBodyFailIds {
	s.TargetIds = v
	return s
}

func (s *DeleteMetricRuleTargetsResponseBodyFailIds) Validate() error {
	if s.TargetIds != nil {
		if err := s.TargetIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds struct {
	TargetId []*string `json:"TargetId,omitempty" xml:"TargetId,omitempty" type:"Repeated"`
}

func (s DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds) GetTargetId() []*string {
	return s.TargetId
}

func (s *DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds) SetTargetId(v []*string) *DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds {
	s.TargetId = v
	return s
}

func (s *DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds) Validate() error {
	return dara.Validate(s)
}
