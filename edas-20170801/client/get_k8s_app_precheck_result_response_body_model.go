// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetK8sAppPrecheckResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetK8sAppPrecheckResultResponseBody
	GetCode() *int32
	SetData(v *GetK8sAppPrecheckResultResponseBodyData) *GetK8sAppPrecheckResultResponseBody
	GetData() *GetK8sAppPrecheckResultResponseBodyData
	SetMessage(v string) *GetK8sAppPrecheckResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetK8sAppPrecheckResultResponseBody
	GetRequestId() *string
}

type GetK8sAppPrecheckResultResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	Data *GetK8sAppPrecheckResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B909AB1F-3763-4963-B1CE-0BDFA192****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetK8sAppPrecheckResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetK8sAppPrecheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetK8sAppPrecheckResultResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetK8sAppPrecheckResultResponseBody) GetData() *GetK8sAppPrecheckResultResponseBodyData {
	return s.Data
}

func (s *GetK8sAppPrecheckResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetK8sAppPrecheckResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetK8sAppPrecheckResultResponseBody) SetCode(v int32) *GetK8sAppPrecheckResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetK8sAppPrecheckResultResponseBody) SetData(v *GetK8sAppPrecheckResultResponseBodyData) *GetK8sAppPrecheckResultResponseBody {
	s.Data = v
	return s
}

func (s *GetK8sAppPrecheckResultResponseBody) SetMessage(v string) *GetK8sAppPrecheckResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetK8sAppPrecheckResultResponseBody) SetRequestId(v string) *GetK8sAppPrecheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetK8sAppPrecheckResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetK8sAppPrecheckResultResponseBodyData struct {
	// The precheck result for the application change.
	JobResults []*GetK8sAppPrecheckResultResponseBodyDataJobResults `json:"JobResults,omitempty" xml:"JobResults,omitempty" type:"Repeated"`
	// The reason why the application failed the precheck. This parameter is left empty when the application passed the precheck.
	//
	// example:
	//
	// The Kubernetes cluster is disconnected from the EDAS control plane.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The precheck state for the application change. Valid values:
	//
	// 	- checking: The application is being prechecked.
	//
	// 	- pass: The application passed the precheck.
	//
	// 	- failed: The application failed the precheck.
	//
	// example:
	//
	// checking
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetK8sAppPrecheckResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetK8sAppPrecheckResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetK8sAppPrecheckResultResponseBodyData) GetJobResults() []*GetK8sAppPrecheckResultResponseBodyDataJobResults {
	return s.JobResults
}

func (s *GetK8sAppPrecheckResultResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *GetK8sAppPrecheckResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetK8sAppPrecheckResultResponseBodyData) SetJobResults(v []*GetK8sAppPrecheckResultResponseBodyDataJobResults) *GetK8sAppPrecheckResultResponseBodyData {
	s.JobResults = v
	return s
}

func (s *GetK8sAppPrecheckResultResponseBodyData) SetReason(v string) *GetK8sAppPrecheckResultResponseBodyData {
	s.Reason = &v
	return s
}

func (s *GetK8sAppPrecheckResultResponseBodyData) SetStatus(v string) *GetK8sAppPrecheckResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetK8sAppPrecheckResultResponseBodyData) Validate() error {
	if s.JobResults != nil {
		for _, item := range s.JobResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetK8sAppPrecheckResultResponseBodyDataJobResults struct {
	// Specifies whether the precheck of the item was interrupted:
	//
	// 	- true: The precheck of the item was interrupted.
	//
	// 	- false: The precheck of the item was not interrupted.
	//
	// example:
	//
	// false
	Interrupted *bool `json:"Interrupted,omitempty" xml:"Interrupted,omitempty"`
	// The name of the precheck item.
	//
	// example:
	//
	// Cluster Health Check
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the precheck item passed the precheck:
	//
	// 	- true: The precheck item passed the precheck.
	//
	// 	- false: The precheck item failed the precheck.
	//
	// example:
	//
	// true
	Pass *bool `json:"Pass,omitempty" xml:"Pass,omitempty"`
	// The reason why the precheck item failed the precheck or the precheck of the item was interrupted. This parameter is left empty when the application passed the precheck.
	//
	// example:
	//
	// The Kubernetes cluster is disconnected from the EDAS control plane.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s GetK8sAppPrecheckResultResponseBodyDataJobResults) String() string {
	return dara.Prettify(s)
}

func (s GetK8sAppPrecheckResultResponseBodyDataJobResults) GoString() string {
	return s.String()
}

func (s *GetK8sAppPrecheckResultResponseBodyDataJobResults) GetInterrupted() *bool {
	return s.Interrupted
}

func (s *GetK8sAppPrecheckResultResponseBodyDataJobResults) GetName() *string {
	return s.Name
}

func (s *GetK8sAppPrecheckResultResponseBodyDataJobResults) GetPass() *bool {
	return s.Pass
}

func (s *GetK8sAppPrecheckResultResponseBodyDataJobResults) GetReason() *string {
	return s.Reason
}

func (s *GetK8sAppPrecheckResultResponseBodyDataJobResults) SetInterrupted(v bool) *GetK8sAppPrecheckResultResponseBodyDataJobResults {
	s.Interrupted = &v
	return s
}

func (s *GetK8sAppPrecheckResultResponseBodyDataJobResults) SetName(v string) *GetK8sAppPrecheckResultResponseBodyDataJobResults {
	s.Name = &v
	return s
}

func (s *GetK8sAppPrecheckResultResponseBodyDataJobResults) SetPass(v bool) *GetK8sAppPrecheckResultResponseBodyDataJobResults {
	s.Pass = &v
	return s
}

func (s *GetK8sAppPrecheckResultResponseBodyDataJobResults) SetReason(v string) *GetK8sAppPrecheckResultResponseBodyDataJobResults {
	s.Reason = &v
	return s
}

func (s *GetK8sAppPrecheckResultResponseBodyDataJobResults) Validate() error {
	return dara.Validate(s)
}
