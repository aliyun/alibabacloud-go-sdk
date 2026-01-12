// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeSyncAlgorithmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *InvokeSyncAlgorithmRequest
	GetJobId() *string
	SetName(v string) *InvokeSyncAlgorithmRequest
	GetName() *string
	SetParams(v map[string]interface{}) *InvokeSyncAlgorithmRequest
	GetParams() map[string]interface{}
}

type InvokeSyncAlgorithmRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Params map[string]interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s InvokeSyncAlgorithmRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeSyncAlgorithmRequest) GoString() string {
	return s.String()
}

func (s *InvokeSyncAlgorithmRequest) GetJobId() *string {
	return s.JobId
}

func (s *InvokeSyncAlgorithmRequest) GetName() *string {
	return s.Name
}

func (s *InvokeSyncAlgorithmRequest) GetParams() map[string]interface{} {
	return s.Params
}

func (s *InvokeSyncAlgorithmRequest) SetJobId(v string) *InvokeSyncAlgorithmRequest {
	s.JobId = &v
	return s
}

func (s *InvokeSyncAlgorithmRequest) SetName(v string) *InvokeSyncAlgorithmRequest {
	s.Name = &v
	return s
}

func (s *InvokeSyncAlgorithmRequest) SetParams(v map[string]interface{}) *InvokeSyncAlgorithmRequest {
	s.Params = v
	return s
}

func (s *InvokeSyncAlgorithmRequest) Validate() error {
	return dara.Validate(s)
}
