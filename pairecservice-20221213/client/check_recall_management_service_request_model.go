// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckRecallManagementServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextParams(v map[string]interface{}) *CheckRecallManagementServiceRequest
	GetContextParams() map[string]interface{}
	SetExposureList(v []*string) *CheckRecallManagementServiceRequest
	GetExposureList() []*string
	SetInstanceId(v string) *CheckRecallManagementServiceRequest
	GetInstanceId() *string
	SetRecallManagementServiceVersionId(v string) *CheckRecallManagementServiceRequest
	GetRecallManagementServiceVersionId() *string
	SetRecalls(v map[string]interface{}) *CheckRecallManagementServiceRequest
	GetRecalls() map[string]interface{}
	SetRetainFields(v []*string) *CheckRecallManagementServiceRequest
	GetRetainFields() []*string
	SetUid(v string) *CheckRecallManagementServiceRequest
	GetUid() *string
}

type CheckRecallManagementServiceRequest struct {
	// The context parameters.
	ContextParams map[string]interface{} `json:"ContextParams,omitempty" xml:"ContextParams,omitempty"`
	// The exposure list.
	ExposureList []*string `json:"ExposureList,omitempty" xml:"ExposureList,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// learn-pairec-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The recall management service version ID.
	//
	// example:
	//
	// 2
	RecallManagementServiceVersionId *string `json:"RecallManagementServiceVersionId,omitempty" xml:"RecallManagementServiceVersionId,omitempty"`
	// The recall data details.
	Recalls map[string]interface{} `json:"Recalls,omitempty" xml:"Recalls,omitempty"`
	// The retained fields.
	RetainFields []*string `json:"RetainFields,omitempty" xml:"RetainFields,omitempty" type:"Repeated"`
	// The user ID.
	//
	// example:
	//
	// 1001
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s CheckRecallManagementServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckRecallManagementServiceRequest) GoString() string {
	return s.String()
}

func (s *CheckRecallManagementServiceRequest) GetContextParams() map[string]interface{} {
	return s.ContextParams
}

func (s *CheckRecallManagementServiceRequest) GetExposureList() []*string {
	return s.ExposureList
}

func (s *CheckRecallManagementServiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CheckRecallManagementServiceRequest) GetRecallManagementServiceVersionId() *string {
	return s.RecallManagementServiceVersionId
}

func (s *CheckRecallManagementServiceRequest) GetRecalls() map[string]interface{} {
	return s.Recalls
}

func (s *CheckRecallManagementServiceRequest) GetRetainFields() []*string {
	return s.RetainFields
}

func (s *CheckRecallManagementServiceRequest) GetUid() *string {
	return s.Uid
}

func (s *CheckRecallManagementServiceRequest) SetContextParams(v map[string]interface{}) *CheckRecallManagementServiceRequest {
	s.ContextParams = v
	return s
}

func (s *CheckRecallManagementServiceRequest) SetExposureList(v []*string) *CheckRecallManagementServiceRequest {
	s.ExposureList = v
	return s
}

func (s *CheckRecallManagementServiceRequest) SetInstanceId(v string) *CheckRecallManagementServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *CheckRecallManagementServiceRequest) SetRecallManagementServiceVersionId(v string) *CheckRecallManagementServiceRequest {
	s.RecallManagementServiceVersionId = &v
	return s
}

func (s *CheckRecallManagementServiceRequest) SetRecalls(v map[string]interface{}) *CheckRecallManagementServiceRequest {
	s.Recalls = v
	return s
}

func (s *CheckRecallManagementServiceRequest) SetRetainFields(v []*string) *CheckRecallManagementServiceRequest {
	s.RetainFields = v
	return s
}

func (s *CheckRecallManagementServiceRequest) SetUid(v string) *CheckRecallManagementServiceRequest {
	s.Uid = &v
	return s
}

func (s *CheckRecallManagementServiceRequest) Validate() error {
	return dara.Validate(s)
}
